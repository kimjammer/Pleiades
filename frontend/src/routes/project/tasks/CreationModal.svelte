<script lang="ts">
    import UserAvatar from "$lib/components/UserAvatar.svelte"
    import { buttonVariants } from "$lib/components/ui/button"
    import * as Dialog from "$lib/components/ui/dialog"
    import * as Form from "$lib/components/ui/form"
    import { Input } from "$lib/components/ui/input"
    import * as ToggleGroup from "$lib/components/ui/toggle-group"
    import type { ProjectState, Task } from "$lib/project_state.svelte"
    import { taskformSchema, type TaskFormSchema, type UserId } from "$lib/schema"
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

    async function createTask(e: Event) {
        // Otherwise the form would overwrite the URL params and you would quit out of the project on reload
        e.preventDefault()

        const validationResult = await form.validateForm({ update: true })
        if (!validationResult.valid) return
        console.log(validationResult.data)
        project.appendInProject("Tasks", {
            ...validationResult.data,
            id: crypto.randomUUID(),
            dueDate: validationResult.data.dueDate
                ? new Date(validationResult.data.dueDate).getTime()
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
