<script lang="ts">
    import type { ProjectState } from "$lib/project_state.svelte"
    import KanbanColumn from "./KanbanColumn.svelte"

    let { project }: { project: ProjectState } = $props()
</script>

<div class="board">
    <KanbanColumn
        {project}
        columnName="Not started"
        columnId=""
    />
    <KanbanColumn
        {project}
        columnName="In progress"
        columnId="progress"
    />
    <KanbanColumn
        {project}
        columnName="Done"
        columnId="done"
    />
</div>

<style>
    .board {
        display: flex;
        align-items: stretch;
        min-height: 80vh;
    }
</style>
