<script lang="ts">
    import { Button } from "$lib/components/ui/button"
    import * as Dialog from "$lib/components/ui/dialog"
    import * as Tabs from "$lib/components/ui/tabs/index.js"
    import type { ProjectState } from "$lib/project_state.svelte"
    import QrAlert from "./QrAlert.svelte"
    import Teammates from "./Teammates.svelte"

    let { project }: { project: ProjectState } = $props()

    let leaveDialogOpen = $state(false)
    let deleteDialogOpen = $state(false)
</script>

<Tabs.Content value="settings">
    <QrAlert />
    <Dialog.Root bind:open={leaveDialogOpen}>
        <Dialog.Trigger>
            <Button variant="destructive">Leave Project</Button>
        </Dialog.Trigger>
        <Dialog.Content>
            <Dialog.Header>Are you sure you want to leave the project?</Dialog.Header>
            <div>
                <Button
                    onclick={() => {
                        project.leave()
                    }}>Confirm</Button
                >
                <Button
                    onclick={() => {
                        leaveDialogOpen = false
                    }}>Cancel</Button
                >
            </div>
        </Dialog.Content>
    </Dialog.Root>

    <Dialog.Root bind:open={deleteDialogOpen}>
        <Dialog.Trigger>
            <Button variant="destructive">Delete Project</Button>
        </Dialog.Trigger>
        <Dialog.Content>
            <Dialog.Header>Are you sure you want to delete the project?</Dialog.Header>
            <div>
                <Button
                    onclick={() => {
                        project.delete()
                    }}>Confirm</Button
                >
                <Button
                    onclick={() => {
                        deleteDialogOpen = false
                    }}>Cancel</Button
                >
            </div>
        </Dialog.Content>
    </Dialog.Root>
    <Teammates {project} />
</Tabs.Content>
