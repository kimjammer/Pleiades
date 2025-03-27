<script lang="ts">
    import { buttonVariants } from "$lib/components/ui/button"
    import * as Dialog from "$lib/components/ui/dialog"
    import * as Form from "$lib/components/ui/form"
    import { Input } from "$lib/components/ui/input"
    import type { ProjectState } from "$lib/project_state.svelte"
    import { formSchema, type FormSchema } from "$lib/schema"
    import { type Infer, superForm, type SuperValidated } from "sveltekit-superforms"
    import { zodClient } from "sveltekit-superforms/adapters"

    let { project, data }: { project: ProjectState; data: { form: SuperValidated<Infer<FormSchema>> } } = $props()
    let createDialogOpen = $state(false)

    const form = superForm(data.form, {
        validators: zodClient(formSchema),
    })

    const { form: formData } = form

    async function createPoll() {
        const validationResult = await form.validateForm({ update: true })
        if (!validationResult.valid) return
        console.log(validationResult.data)
        project.appendInProject("Polls", validationResult.data)
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
            <form onsubmit={createPoll}>
                <Form.Field {form} name="question">
                    <Form.Control>
                        {#snippet children({ props })}
                        <Form.Label>Question</Form.Label>
                        <Input {...props} bind:value={$formData.question} />
                        {/snippet}
                    </Form.Control>
                    <Form.FieldErrors />
                </Form.Field>

                <Form.Field {form} name="options">
                    <Form.Control>
                        {#snippet children({ props })}
                        <Form.Label>Options (comma separated)</Form.Label>
                        <Input {...props} bind:value={$formData.options} />
                        {/snippet}
                    </Form.Control>
                    <Form.Description>Enter poll choices separated by commas</Form.Description>
                    <Form.FieldErrors />
                </Form.Field>

                <Form.Field {form} name="deadline">
                    <Form.Control>
                        {#snippet children({ props })}
                        <Form.Label>Voting Deadline</Form.Label>
                        <Input {...props} type="date" bind:value={$formData.deadline} />
                        {/snippet}
                    </Form.Control>
                    <Form.FieldErrors />
                </Form.Field>

                <Dialog.Footer>
                    <Form.Button>Create Poll</Form.Button>
                </Dialog.Footer>
            </form>
        </Dialog.Header>
    </Dialog.Content>
</Dialog.Root>
