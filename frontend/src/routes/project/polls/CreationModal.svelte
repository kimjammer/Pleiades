<script lang="ts">
    import { buttonVariants } from "$lib/components/ui/button"
    import * as Dialog from "$lib/components/ui/dialog"
    import * as Form from "$lib/components/ui/form"
    import { Input } from "$lib/components/ui/input"
    import { type ProjectState, Poll } from "$lib/project_state.svelte"
    import { toast } from "svelte-sonner"
    import { pollformSchema, type PollFormSchema } from "$lib/schema"
    import { type Infer, superForm, type SuperValidated} from "sveltekit-superforms";
    import {zodClient} from "sveltekit-superforms/adapters";

    let {
        project,
        data,
    }: { project: ProjectState; data: { form: SuperValidated<Infer<PollFormSchema>> } } = $props()
    let createDialogOpen = $state(false)

    const form = superForm(pollformSchema, {
        validators: zodClient(pollformSchema),
        dataType: "json",
    })

    let error = false

    const { form: formData } = form


    function validateInput() { //does error checking
        error = false
        if ($formData.title.trim() === "") {
            error = true
            toast.error("Title must be at least one character")
        }
        if ($formData.options.split(",").length < 2) {
            error = true
            toast.error("The poll must have at least two options")
        }
        if ($formData.due === "") {
            error = true
            toast.error("The poll must have an end date")
        }
        if (!error) {
            createPoll()
        }

    }

    async function createPoll() {
        const validationResult = await form.validateForm({ update: true })
        if (!validationResult.valid) return
        console.log(validationResult.data)

        let optionsArray: Option[] = $formData.options.split(",").map((option, index) => ({
            id: String(index), // Assign unique ID (you may generate UUID if needed)
            title: option.trim(), // Remove any extra spaces
            likedUsers: [],
            neutralUsers: [],
            dislikedUsers: []
        }))

        let newPoll: Poll = {
            id: crypto.randomUUID(), // Generate a unique ID
            title: $formData.title,
            description: $formData.description,
            options: optionsArray,
            dueDate: due
        }

        project.appendInProject("Polls", newPoll)
        createDialogOpen = false
        form.reset()
    }


</script>

<Dialog.Root bind:open={createDialogOpen}>
    <Dialog.Trigger class={buttonVariants({ variant: "outline" })}>
        Create a new Poll
    </Dialog.Trigger>
    <Dialog.Content>
        <Dialog.Header>
            <Dialog.Title>Create new poll</Dialog.Title>
            <Dialog.Description>
                Enter poll question and options
            </Dialog.Description>
            <form onsubmit={validateInput}>
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
                        name="options"
                >
                    <Form.Control>
                        {#snippet children({ props })}
                        <Form.Label>Options (comma-separated)</Form.Label>
                        <Input
                                {...props}
                                bind:value={$formData.options}
                        />
                        {/snippet}
                    </Form.Control>
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
                                bind:value={$formData.due}
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
