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
    class="kanban-column border"
    id={columnId}
    role="list"
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
    .kanban-column {
        text-align: center;
        flex-grow: 1;
        flex-basis: 0;
        display: flex;
        flex-direction: column;
        align-items: stretch;
        border-radius: 0.5em;
    }
</style>
