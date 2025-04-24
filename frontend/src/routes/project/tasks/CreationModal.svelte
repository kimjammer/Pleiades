<script lang="ts">
    import UserAvatar from "$lib/components/UserAvatar.svelte"
    import { buttonVariants } from "$lib/components/ui/button"
    import * as Dialog from "$lib/components/ui/dialog"
    import * as Form from "$lib/components/ui/form"
    import { Input } from "$lib/components/ui/input"
    import * as ToggleGroup from "$lib/components/ui/toggle-group"
    import type { ProjectState, Task } from "$lib/project_state.svelte"
    import { recordEvent } from "$lib/restApi"
    import { taskformSchema, type TaskFormSchema, type UserId } from "$lib/schema"
    import { debounce } from "$lib/utils"
    import * as chrono from "chrono-node"
    import type { FormEventHandler } from "svelte/elements"
    import { superForm, type Infer, type SuperValidated } from "sveltekit-superforms"
    import { zodClient } from "sveltekit-superforms/adapters"

    let {
        project,
        data,
    }: { project: ProjectState; data: { taskform: SuperValidated<Infer<TaskFormSchema>> } } =
        $props()
    let createDialogOpen = $state(false)
    const form = superForm(data.taskform, {
        validators: zodClient(taskformSchema),
    })

    const { form: formData } = form

    //For assigning users
    let toggledUsers = new Map<UserId, boolean>()

    function toggleUser(userID: UserId) {
        // Flip user's toggle state
        toggledUsers.set(userID, !toggledUsers.get(userID))
    }

    const onTitleChange: FormEventHandler<HTMLInputElement> = ev => {
        let newVal = (ev.target as HTMLInputElement).value
        let usedNlp = false

        // nlp date
        const parseResult = chrono.parse(newVal, new Date())[0]
        if (parseResult) {
            let parsedDate = parseResult.date()
            // Date input can't handle Date objects ironically, so I'm using this workaround
            // to ensure when converting to ISO string (UTC) it remains the correct date: https://stackoverflow.com/a/29774197
            const offset = parsedDate.getTimezoneOffset()
            parsedDate = new Date(parsedDate.getTime() - offset * 60 * 1000)
            const dueDate = parsedDate.toISOString().split("T")[0]

            usedNlp = true
            newVal = newVal.replace(parseResult.text, "").trim()
            form.form.update(form => ({
                ...form,
                dueDate,
                title: newVal,
            }))
        }

        // nlp time estimate
        let totalHourEstimate = 0
        const minutesMatch = newVal.match(/(\d+)(?:m| *min(?:utes?)?)(?= |$|\d)/)
        const hoursMatch = newVal.match(/(\d+)(?:h| *(?:hours?|hrs?))(?= |$|\d)/)

        if (hoursMatch) {
            usedNlp = true
            const { [0]: fullMatch, [1]: hourVal } = hoursMatch
            newVal = newVal.replace(fullMatch, "").trim()
            totalHourEstimate += Number.parseInt(hourVal)
        }
        if (minutesMatch) {
            usedNlp = true
            const { [0]: fullMatch, [1]: minuteVal } = minutesMatch
            newVal = newVal.replace(fullMatch, "").trim()
            totalHourEstimate += Number.parseInt(minuteVal) / 60
        }

        form.form.update(form => ({
            ...form,
            timeEstimate: form.timeEstimate + totalHourEstimate,
            title: newVal,
        }))
        if (usedNlp) recordEvent("task_nlp")
    }

    async function createTask(e: Event) {
        // Otherwise the form would overwrite the URL params and you would quit out of the project on reload
        e.preventDefault()

        const validationResult = await form.validateForm({ update: true })
        if (!validationResult.valid) return
        console.log(validationResult.data)

        const offset = new Date().getTimezoneOffset()

        project.appendInProject("Tasks", {
            ...validationResult.data,
            id: crypto.randomUUID(),
            projectId: project.id,
            dueDate: validationResult.data.dueDate
                ? new Date(validationResult.data.dueDate).getTime() + offset * 60 * 1000
                : 0,
            timeEstimate: validationResult.data.timeEstimate * 60 * 60 * 1000,
            kanbanColumn: "",
            completed: false,
            sessions: [],
            assignees: Array.from(toggledUsers.entries())
                .filter(([_, value]) => value)
                .map(([userID]) => userID),
        } satisfies Task)
        console.log("assigned users:")
        console.log(
            Array.from(toggledUsers.entries())
                .filter(([_, value]) => value)
                .map(([userID]) => userID),
        )
        recordEvent("tasks")
        createDialogOpen = false
        form.reset()
    }
</script>

<Dialog.Root bind:open={createDialogOpen}>
    <Dialog.Trigger class={buttonVariants({ variant: "outline" })}>
        Create a new task
    </Dialog.Trigger>
    <Dialog.Content>
        <Dialog.Header>
            <Dialog.Title>Create new task</Dialog.Title>
            <Dialog.Description>
                Only title required, but all strongly recommended
            </Dialog.Description>
            <form onsubmit={createTask}>
                <Form.Field
                    {form}
                    name="title"
                >
                    <Form.Control>
                        {#snippet children({ props })}
                            <Form.Label>Title</Form.Label>
                            <Input
                                {...props}
                                bind:value={$formData.title}
                                oninput={debounce(onTitleChange, 1000)}
                            />
                        {/snippet}
                    </Form.Control>
                    <Form.Description
                        >Tip: use natural language here to set time estimate and due date (eg: 30
                        min tuesday)</Form.Description
                    >
                    <Form.FieldErrors />
                </Form.Field>

                <Form.Field
                    {form}
                    name="description"
                >
                    <Form.Control>
                        {#snippet children({ props })}
                            <Form.Label>Description</Form.Label>
                            <Input
                                {...props}
                                bind:value={$formData.description}
                            />
                        {/snippet}
                    </Form.Control>
                    <Form.Description>Implementation details, progress, or notes</Form.Description>
                    <Form.FieldErrors />
                </Form.Field>

                <Form.Field
                    {form}
                    name="dueDate"
                >
                    <Form.Control>
                        {#snippet children({ props })}
                            <Form.Label>Due date</Form.Label>
                            <Input
                                {...props}
                                type="date"
                                bind:value={$formData.dueDate}
                            />
                        {/snippet}
                    </Form.Control>
                    <Form.FieldErrors />
                </Form.Field>

                <Form.Field
                    {form}
                    name="timeEstimate"
                >
                    <Form.Control>
                        {#snippet children({ props })}
                            <Form.Label>Time Estimate</Form.Label>
                            <Input
                                {...props}
                                type="number"
                                min="0"
                                step="0.01"
                                bind:value={$formData.timeEstimate}
                            />
                        {/snippet}
                    </Form.Control>
                    <Form.FieldErrors />
                </Form.Field>

                <ToggleGroup.Root
                    type="multiple"
                    variant="outline"
                >
                    {#each project.users as user}
                        <ToggleGroup.Item
                            value={user.id}
                            class="p-3"
                            onclick={() => toggleUser(user.id)}
                        >
                            <UserAvatar
                                {project}
                                userID={user.id}
                            />
                        </ToggleGroup.Item>
                    {/each}
                </ToggleGroup.Root>
                <Dialog.Footer>
                    <Form.Button>Create!</Form.Button>
                </Dialog.Footer>
            </form>
        </Dialog.Header>
    </Dialog.Content>
</Dialog.Root>
