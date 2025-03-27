<script lang="ts">
    import { buttonVariants } from "$lib/components/ui/button"
    import * as Dialog from "$lib/components/ui/dialog"
    import * as Form from "$lib/components/ui/form"
    import { Input } from "$lib/components/ui/input"
    import type { ProjectState } from "$lib/project_state.svelte"
    import { pollformSchema, type PollFormSchema } from "$lib/schema"
    import { type Infer, superForm, type SuperValidated } from "sveltekit-superforms"
    import { zodClient } from "sveltekit-superforms/adapters"

    let { project }: { project: ProjectState } = $props()
    let createDialogOpen = $state(true)


    async function createPoll() {

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
        </Dialog.Header>
    </Dialog.Content>
</Dialog.Root>
