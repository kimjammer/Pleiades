package main

import (
	"encoding/json"
	"errors"
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
func applyCommandToProject(projectId string, command Command) {
	connection := joinSpace(projectId)

	connection.command_tx <- command
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
func updateProject(projectId string, updater func(*Project) error) {
	applyCommandToProject(projectId, FunctionCommand{function: updater})
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
	slice := reflect.ValueOf(value).Elem()

	if slice.Kind() != reflect.Slice && slice.Kind() != reflect.Array {
		return reflect.ValueOf(nil), errors.New("Cannot index into " + slice.Kind().String() + " because it is neither an array nor a slice")
	}

	len := slice.Len()

	for i := range len {
		value := slice.Index(i)

		if value.Kind() != reflect.Struct {
			return reflect.ValueOf(nil), errors.New("Cannot get field " + self.key + " of " + value.Kind().String() + " because it is not a structure")

		}

		field := value.FieldByName(self.key)

		if !field.IsValid() {
			return reflect.ValueOf(nil), errors.New("The field " + self.key + " does not exist on the type " + value.Kind().String())
		}

		if v, ok := field.Interface().(string); ok {
			if v == self.value {
				return field.Addr(), nil
			}
		} else {
			return reflect.ValueOf(nil), errors.New("The field " + self.key + "  is not a string type" + value.Kind().String())
		}
	}

	return reflect.ValueOf(nil), errors.New("Unable to find an element with field " + self.key + " matching " + self.value)
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

func decodeSelector(selector string) ([]SelectorLevel, error) {
	if selector == "" {
		return []SelectorLevel{}, nil
	}

	var removed string
	var parsedSelector SelectorLevel

	var maybeArraySelector = arraySelectorRegex.FindString(selector)

	if maybeArraySelector != "" {
		removed = arraySelectorRegex.ReplaceAllString(selector, "")

		values := strings.Split(strings.Trim(maybeArraySelector, "[]"), "=")

		parsedSelector = ArrayElementSelector{key: values[0], value: values[1]}
	} else {
		var maybeFieldSelector = fieldSelectorRegex.FindString(selector)

		if maybeFieldSelector != "" {
			removed = fieldSelectorRegex.ReplaceAllString(selector, "")

			parsedSelector = FieldSelector{key: strings.Trim(maybeFieldSelector, ".")}
		} else if wrongArraySelectorRegex.MatchString(selector) {
			return nil, errors.New("Indexing by numerical index is unsupported because it is very likely to lead to data races if a value is inserted into the array after picking the index. Instead select by the contents of one of the fields, like `[Id=whatever]`. If you really do need this feature, contact Henry")
		} else {
			return nil, errors.New("Could not parse the selector string: " + selector)
		}
	}

	var selectors, err = decodeSelector(removed)
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
	selector, err := decodeSelector(self.Selector)
	if err != nil {
		return err
	}

	selected := reflect.ValueOf(state)
	for _, selectorLevel := range selector {
		selected, err = selectorLevel.doSelect(selected)
		if err != nil {
			return err
		}
	}

	spot := selected.Interface()
	err = json.Unmarshal(self.NewValue, spot)
	if err != nil {
		return err
	}

	return nil
}
