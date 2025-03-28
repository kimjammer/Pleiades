package main

import (
	"encoding/json"
	"errors"
	"log"
	"reflect"
	"regexp"
	"slices"
	"strings"
)

type Command interface {
	apply(project *Project) error
}

// Apply a one-off command to a project server side
//
// This will automatically update all connected users of the change.
func applyCommandToProject(projectId string, command Command) error {
	connection := joinSpace(projectId)

	connection.command_tx <- command

	return <-connection.errors
}

type FunctionCommand struct {
	function func(*Project) error
}

func (self FunctionCommand) apply(project *Project) error {
	return self.function(project)
}

// Update the project via a function
//
// This will automatically update all connected users of the change.
func updateProject(projectId string, updater func(*Project) error) error {
	return applyCommandToProject(projectId, FunctionCommand{function: updater})
}

type CommandMessage struct {
	Name string
	Args json.RawMessage
}

func decodeCommand(command CommandMessage, userId string) (Command, error) {
	if command.Name == "demoButtonState" {
		var newState string
		err := json.Unmarshal(command.Args, &newState)
		if err != nil {
			return nil, err
		}

		return DemoButtonCommand{newState}, nil
	}

	if command.Name == "leave" {
		return UserLeave{userId: userId}, nil
	}

	if command.Name == "delete" {
		return Delete{userId: userId}, nil
	}

	if command.Name == "update" {
		var parsedCommand UpdateInProject
		err := json.Unmarshal(command.Args, &parsedCommand)
		if err != nil {
			return nil, err
		}

		return parsedCommand, nil
	}

	if command.Name == "append" {
		var parsedCommand AppendInProject
		err := json.Unmarshal(command.Args, &parsedCommand)
		if err != nil {
			return nil, err
		}

		return parsedCommand, nil
	}

	if command.Name == "remove" {
		var parsedCommand DeleteInProject
		err := json.Unmarshal(command.Args, &parsedCommand)
		if err != nil {
			return nil, err
		}

		return parsedCommand, nil
	}

	return nil, errors.New("Unknown command: " + command.Name)
}

type DemoButtonCommand struct {
	newState string
}

func (self DemoButtonCommand) apply(state *Project) error {
	if self.newState == "" {
		state.DemoButtonState = ""
	} else if state.DemoButtonState == "" && slices.Contains([]string{"a", "b"}, self.newState) {
		state.DemoButtonState = self.newState
	}

	return nil
}

type UserLeave struct {
	userId string
}

func (self UserLeave) apply(state *Project) error {
	for i, user := range state.Users {
		if user.User == self.userId {
			state.Users[i].LeftProject = true

			return nil
		}
	}

	return errors.New("Could not find the user in the project")
}

type Delete struct {
	userId string
}

func (self Delete) apply(state *Project) error {
	for _, user := range state.Users {
		if user.User != self.userId && !user.LeftProject {
			return errors.New("Cannot delete the project when there are other users still in the project")
		}
	}

	return UserLeave(self).apply(state)
}

type SelectorLevel interface {
	doSelect(value reflect.Value) (reflect.Value, error)
}

type ArrayElementSelector struct {
	key   string
	value string
}

func (self ArrayElementSelector) doSelect(value reflect.Value) (reflect.Value, error) {
	slice := value.Elem()

	if slice.Kind() != reflect.Slice && slice.Kind() != reflect.Array {
		return reflect.ValueOf(nil), errors.New("Cannot index into " + slice.Kind().String() + " because it is neither an array nor a slice")
	}

	len := slice.Len()

	for i := range len {
		value := slice.Index(i)

		ok, err := self.elementMatches(value)
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		if ok {
			return value.Addr(), nil
		}
	}

	return reflect.ValueOf(nil), nil
}

func (self ArrayElementSelector) elementMatches(value reflect.Value) (bool, error) {
	var keyValue reflect.Value

	if self.key == "$IT" {
		keyValue = value
	} else {
		if value.Kind() != reflect.Struct {
			return false, errors.New("Cannot get field " + self.key + " of " + value.Kind().String() + " because it is not a structure")
		}

		field := value.FieldByName(self.key)

		if !field.IsValid() {
			return false, errors.New("The field " + self.key + " does not exist on the type " + value.Kind().String())
		}

		keyValue = field
	}

	if v, ok := keyValue.Interface().(string); ok {
		if v == self.value {
			return true, nil
		}
	} else {
		return false, errors.New("The field " + self.key + "  is not a string type" + value.Kind().String())
	}

	return false, nil
}

type FieldSelector struct {
	key string
}

func (self FieldSelector) doSelect(value reflect.Value) (reflect.Value, error) {
	structure := value.Elem()

	if structure.Kind() != reflect.Struct {
		return reflect.ValueOf(nil), errors.New("Cannot get field " + self.key + " of " + structure.Kind().String() + " because it is not a structure")
	}

	field := structure.FieldByName(self.key)

	if !field.IsValid() {
		return reflect.ValueOf(nil), errors.New("The field " + self.key + " does not exist on " + structure.Kind().String())
	}

	return field.Addr(), nil
}

// Matches the format `[Id=whatever]` or `[Id=whatever].`
var arraySelectorRegex, _ = regexp.Compile(`^\[[^=\[\] ]+=[^\[\]]+\]\.?`)

// Matches the format `[12]` (for a useful error message)
var wrongArraySelectorRegex, _ = regexp.Compile(`^\[[0-9]*\]`)

// Matches the format `whatever` or `whatever.`
var fieldSelectorRegex, _ = regexp.Compile(`^[a-zA-Z]+\.?`)

func decodeSelector(selector string, first bool) ([]SelectorLevel, error) {
	if selector == "" {
		return []SelectorLevel{}, nil
	}

	var removed string
	var parsedSelector SelectorLevel

	var maybeArraySelector = arraySelectorRegex.FindString(selector)

	if maybeArraySelector != "" {
		removed = arraySelectorRegex.ReplaceAllString(selector, "")

		trimmed := strings.Trim(maybeArraySelector, "[].")
		values := strings.Split(trimmed, "=")

		parsedSelector = ArrayElementSelector{key: values[0], value: values[1]}
	} else {
		var maybeFieldSelector = fieldSelectorRegex.FindString(selector)

		if maybeFieldSelector != "" {
			removed = fieldSelectorRegex.ReplaceAllString(selector, "")

			key := strings.Trim(maybeFieldSelector, ".")

			if key == "Users" && first {
				return nil, errors.New("Cannot index into the `Users` field, this data should be maintained by the server or by specialized websockets endpoints (like `leave` or `delete`)")
			}

			parsedSelector = FieldSelector{key: key}
		} else if wrongArraySelectorRegex.MatchString(selector) {
			return nil, errors.New("Indexing by numerical index is unsupported because it is very likely to lead to data races if a value is inserted/deleted into/from the array after picking the index. Instead select by the contents of one of the fields, like `[Id=whatever]`. If you really do need this feature, contact Henry")
		} else {
			return nil, errors.New("Could not parse the selector string: " + selector)
		}
	}

	var selectors, err = decodeSelector(removed, false)
	if err != nil {
		return nil, err
	}

	return slices.Insert(selectors, 0, parsedSelector), nil
}

type UpdateInProject struct {
	Selector string
	NewValue json.RawMessage
}

func (self UpdateInProject) apply(state *Project) error {
	selector, err := decodeSelector(self.Selector, true)
	if err != nil {
		return err
	}

	selected := reflect.ValueOf(state)
	for _, selectorLevel := range selector {
		selected, err = selectorLevel.doSelect(selected)
		if err != nil {
			return err
		}
		if !selected.IsValid() {
			log.Println("Unable to find the value to update. Is it a member of an element that was deleted?")
			return nil
		}
	}

	spot := selected.Interface()
	err = json.Unmarshal(self.NewValue, spot)
	if err != nil {
		return err
	}

	return nil
}

type AppendInProject struct {
	Selector string
	NewValue json.RawMessage
}

func (self AppendInProject) apply(state *Project) error {
	selector, err := decodeSelector(self.Selector, true)
	if err != nil {
		return err
	}

	selected := reflect.ValueOf(state)
	for _, selectorLevel := range selector {
		selected, err = selectorLevel.doSelect(selected)
		if err != nil {
			return err
		}
		if !selected.IsValid() {
			log.Println("Unable to find the slice to append to. Is it a member of an element that was deleted?")
			return nil
		}
	}

	slice := reflect.ValueOf(selected.Interface()).Elem()

	if slice.Kind() != reflect.Slice {
		return errors.New(slice.Kind().String() + " is not a slice. Did you mean to use updateInProject?")
	}

	value := reflect.New(slice.Type().Elem())

	spot := value.Interface()
	err = json.Unmarshal(self.NewValue, spot)
	if err != nil {
		return err
	}

	slice.Set(reflect.Append(slice, value.Elem()))

	return nil
}

type DeleteInProject struct {
	Selector string
}

func (self DeleteInProject) apply(state *Project) error {
	selector, err := decodeSelector(self.Selector, true)
	if err != nil {
		return err
	}

	var finalSelector SelectorLevel
	selected := reflect.ValueOf(state)
	for i, selectorLevel := range selector {
		if i == len(selector)-1 {
			finalSelector = selectorLevel
			break
		}

		selected, err = selectorLevel.doSelect(selected)
		if err != nil {
			return err
		}
		if !selected.IsValid() {
			log.Println("Unable to find the slice to delete from. Is it a member of an element that was deleted?")
			return nil
		}
	}

	slice := selected.Elem()

	if slice.Kind() != reflect.Slice {
		return errors.New(slice.Kind().String() + " is not a slice. Did you mean to use updateInProject?")
	}

	var element ArrayElementSelector
	if maybeElement, ok := finalSelector.(ArrayElementSelector); ok {
		element = maybeElement
	} else {
		return errors.New("The final selector is not an index selector.")
	}

	len := slice.Len()

	for i := range len {
		value := slice.Index(i)

		ok, err := element.elementMatches(value)
		if err != nil {
			return err
		}
		if ok {
			slice1 := slice.Slice(0, i)
			slice2 := slice.Slice(i+1, len)
			finalSlice := reflect.AppendSlice(slice1, slice2)

			slice.Set(finalSlice)

			return nil
		}
	}

	log.Println("Unable to find the element to delete. Was it already deleted?")

	return nil
}
