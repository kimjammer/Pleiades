<script lang="ts">
    import { Button, buttonVariants } from "$lib/components/ui/button"
    import * as Dialog from "$lib/components/ui/dialog"
    import { Label } from "$lib/components/ui/label"
    import { Input } from "$lib/components/ui/input"
    import type { ProjectState } from "$lib/project_state.svelte"
    import * as Form from "$lib/components/ui/form"
    import { formSchema, type FormSchema } from "./schema"
    import { type SuperValidated, type Infer, superForm } from "sveltekit-superforms"
    import { zodClient } from "sveltekit-superforms/adapters"

    let {
        project,
        data,
    }: { project: ProjectState; data: { form: SuperValidated<Infer<FormSchema>> } } = $props()
    let createDialogOpen = $state(false)

    const form = superForm(data.form, {
        validators: zodClient(formSchema),
    })

    const { form: formData, enhance } = form

    function createTask() {}
</script>

<Dialog.Root bind:open={createDialogOpen}>
    <Dialog.Trigger class={buttonVariants({ variant: "outline" })}>
        Create a new project
    </Dialog.Trigger>
    <Dialog.Content>
        <Dialog.Header>
            <Dialog.Title>Create new Project</Dialog.Title>
            <Dialog.Description>
                Choose a name and description for your project and click create!
            </Dialog.Description>
            <form
                method="POST"
                use:enhance
            >
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
                    <Form.Description>This is your public display name.</Form.Description>
                    <Form.FieldErrors />
                </Form.Field>

                <Dialog.Footer>
                    <Form.Button>Create!</Form.Button>
                </Dialog.Footer>
            </form>
        </Dialog.Header>
    </Dialog.Content>
</Dialog.Root>
