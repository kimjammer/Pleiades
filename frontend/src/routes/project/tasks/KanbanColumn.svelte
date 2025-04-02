<script lang="ts">
    import type { ProjectState } from "$lib/project_state.svelte"
    import TaskCard from "./TaskCard.svelte"

    let {
        project,
        columnName,
        columnId,
    }: { project: ProjectState; columnName: string; columnId: string } = $props()
</script>

<div
    class="column border"
    ondragover={e => {
        if (e.dataTransfer === null) {
            throw "Bruh"
        }
        e.preventDefault()
        console.log("over")

        e.dataTransfer.dropEffect = "move"
    }}
    ondrop={e => {
        if (e.dataTransfer === null) {
            throw "Bruh"
        }

        e.preventDefault()

        let id = e.dataTransfer.getData("text/plain")
        console.log(id)

        project.updateInProject(`Tasks[Id=${id}].KanbanColumn`, columnId)
    }}
>
    <h1 class="text-[1.5em]">{columnName}</h1>
    {#each project.tasks as task}
        {#if task.kanbanColumn === columnId}
            <TaskCard
                {project}
                {task}
            />
        {/if}
    {/each}
</div>

<style>
    .column {
        text-align: center;
        flex-grow: 1;
        display: flex;
        flex-direction: column;
        align-items: stretch;
        margin: 0.5em;
        border-radius: 0.5em;
    }
</style>
