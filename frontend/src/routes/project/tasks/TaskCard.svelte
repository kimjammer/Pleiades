<script lang="ts">
    import { ProjectState, Task, mouse } from "$lib/project_state.svelte"
    import * as ContextMenu from "$lib/components/ui/context-menu"

    let { project, task }: { project: ProjectState; task: Task } = $props()

    let dueDate = $derived(new Date(task.dueDate))
    let assignees = $derived(
        task.assignees
            .map(id => project.users.find(user => user.id === id))
            .map(user => `${user?.firstName} ${user?.lastName}`),
    )

    let dragging = $state(false)
    let startX: number = $state(0)
    let startY: number = $state(0)

    let card: HTMLDivElement
</script>

<ContextMenu.Root>
    <ContextMenu.Trigger>
        <div
            bind:this={card}
            class="task bg-muted"
            role="listitem"
            onmousedown={e => {
                startX = mouse.x
                startY = mouse.y
                dragging = true
            }}
            onmouseup={e => {
                setTimeout(() => {
                    dragging = false
                }, 100)

                card.hidden = true
                let elemBelow = document.elementFromPoint(e.clientX, e.clientY)
                card.hidden = false

                let elt = elemBelow?.closest(".kanban-column")

                console.log(elt)
                if (elt != null) {
                    project.updateInProject(`Tasks[Id=${task.id}].KanbanColumn`, elt.id)
                }
            }}
            style={dragging ? `top:${mouse.y - startY}px;left:${mouse.x - startX}px` : ``}
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
        position: relative;
    }

    .date {
        position: absolute;
        top: 0px;
        right: 0.4em;
    }
</style>
