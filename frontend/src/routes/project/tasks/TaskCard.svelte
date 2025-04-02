<script lang="ts">
    import type { ProjectState, Task } from "$lib/project_state.svelte"
    import * as ContextMenu from "$lib/components/ui/context-menu"

    let { project, task }: { project: ProjectState; task: Task } = $props()

    let dueDate = $derived(new Date(task.dueDate))
    let assignees = $derived(
        task.assignees
            .map(id => project.users.find(user => user.id === id))
            .map(user => `${user?.firstName} ${user?.lastName}`),
    )
</script>

<ContextMenu.Root>
    <ContextMenu.Trigger>
        <div
            class="task bg-muted"
            draggable="true"
            ondragstart={e => {
                if (e.dataTransfer === null) {
                    throw "Bruh"
                }

                e.dataTransfer.setData("text/plain", task.id)
                e.dataTransfer.dropEffect = "move"
            }}
        >
            <h1 class="text-[1.2em]">{task.title}</h1>
            <p>{task.description}</p>
            {#if task.dueDate !== 0}
                <p class="date">{dueDate.toLocaleDateString()}</p>
            {/if}

            {#each assignees as assignee}
                <p>{assignee}</p>
            {/each}
        </div>
    </ContextMenu.Trigger>
    <ContextMenu.Content>
        <ContextMenu.Item
            onclick={() => {
                // Todo
            }}>Assign</ContextMenu.Item
        >
        <ContextMenu.Item
            onclick={() => {
                project.deleteInProject(`Tasks[Id=${task.id}]`)
            }}>Delete</ContextMenu.Item
        >
    </ContextMenu.Content>
</ContextMenu.Root>

<style>
    .task {
        margin: 0.5em;
        border-radius: 0.5em;
        position: relative; /* Allow the date to position itself absolutely inside this element */
    }

    .date {
        position: absolute;
        top: 0px;
        right: 0.4em;
    }
</style>
