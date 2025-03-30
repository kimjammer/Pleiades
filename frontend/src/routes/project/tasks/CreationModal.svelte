<script lang="ts">
    import { buttonVariants } from "$lib/components/ui/button"
    import * as Dialog from "$lib/components/ui/dialog"
    import * as Form from "$lib/components/ui/form"
    import { Input } from "$lib/components/ui/input"
    import type { ProjectState, Task } from "$lib/project_state.svelte"
    import { taskformSchema, type TaskFormSchema } from "$lib/schema"
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

    async function createTask() {
        const validationResult = await form.validateForm({ update: true })
        if (!validationResult.valid) return
        const taskData = validationResult.data
        console.log(validationResult.data)
        project.appendInProject("Tasks", {
            ...validationResult.data,
            id: crypto.randomUUID(),
            dueDate: validationResult.data.due ? new Date(validationResult.data.due).getTime() : 0,
            kanbanColumn: "",
            completed: false,
            sessions: [],
        } satisfies Task)
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
                    name="due"
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
                    name="estimate"
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

                <!-- TODO: assignees -->

                <Dialog.Footer>
                    <Form.Button>Create!</Form.Button>
                </Dialog.Footer>
            </form>
        </Dialog.Header>
    </Dialog.Content>
</Dialog.Root>
