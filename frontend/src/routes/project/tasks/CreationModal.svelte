<script lang="ts">
    import { Button, buttonVariants } from "$lib/components/ui/button"
    import * as Dialog from "$lib/components/ui/dialog"
    import { Label } from "$lib/components/ui/label"
    import { Input } from "$lib/components/ui/input"
    import type { ProjectState } from "$lib/project_state.svelte"

    let { project }: { project: ProjectState } = $props()
    let createDialogOpen = $state(false)

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
            <div class="flex flex-col gap-5 py-5">
                <div>
                    <Label
                        for="title"
                        class="text-right">Title</Label
                    >
                    <Input
                        id="title"
                        placeholder="Project Name"
                    />
                </div>
                <div>
                    <Label
                        for="description"
                        class="text-right">Description</Label
                    >
                    <Input
                        id="description"
                        placeholder="Project Description"
                    />
                </div>
            </div>
            <Dialog.Footer>
                <Button
                    onclick={createTask}
                    onkeypress={createTask}
                    type="submit"
                >
                    Create!
                </Button>
            </Dialog.Footer>
        </Dialog.Header>
    </Dialog.Content>
</Dialog.Root>
