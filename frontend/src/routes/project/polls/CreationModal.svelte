<script lang="ts">
    import { buttonVariants } from "$lib/components/ui/button"
    import { Input } from "$lib/components/ui/input"
    import { Option, Poll, type ProjectState } from "$lib/project_state.svelte"
    import { toast } from "svelte-sonner"
    import { pollformSchema, type PollFormSchema } from "$lib/schema"
    import { type Infer, superForm, type SuperValidated } from "sveltekit-superforms"
    import { zodClient } from "sveltekit-superforms/adapters"
    import * as Dialog from "$lib/components/ui/dialog"
    import * as Form from "$lib/components/ui/form"

    let {
        project,
        data,
    }: { project: ProjectState; data: { pollform: SuperValidated<Infer<PollFormSchema>> } } =
        $props()
    let createDialogOpen = $state(false)
    let newPoll: Poll

    const form = superForm(data.pollform, {
        validators: zodClient(pollformSchema),
    })

    let error = false
    let options = $state("")
    const { form: formData } = form

    // function validateInput() {
    //     //does error checking
    //     error = false
    //     if ($formData.title.trim() === "") {
    //         error = true
    //         toast.error("Title must be at least one character")
    //     }
    //     if ($formData.due === "") {
    //         error = true
    //         toast.error("The poll must have an end date")
    //     }
    //     if (!error) {
    //         createPoll()
    //     }
    // }

    async function createPoll() {
        console.log("creating poll")
        console.log($formData)
        const validationResult = await form.validateForm({ update: true })
        if (!validationResult.valid) return
        console.log(validationResult.data)

        project.appendInProject<Poll>("Polls", {
            id: crypto.randomUUID(),
            title: validationResult.data.title,
            description: validationResult.data.description ?? "",
            dueDate: validationResult.data.dueDate,
            options: validationResult.data.options.split(",").map(title => {
                return {
                    id: crypto.randomUUID(),
                    title: title,
                    likedUsers: [],
                    neutralUsers: [],
                    dislikedUsers: [],
                }
            }),
        })

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
            <Dialog.Description>Enter poll question and options</Dialog.Description>
            <form onsubmit={createPoll}>
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
                    <Form.Description></Form.Description>
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

                <Dialog.Footer>
                    <Form.Button>Create</Form.Button>
                </Dialog.Footer>
            </form>
        </Dialog.Header>
    </Dialog.Content>
</Dialog.Root>
