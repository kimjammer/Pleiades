<script lang="ts">
    import * as Tabs from "$lib/components/ui/tabs"
    import { type ProjectState } from "$lib/project_state.svelte"
    import type { PageData } from "./$types"
    import CreationModal from "./CreationModal.svelte"
    import KanbanBoard from "./KanbanBoard.svelte"

    let { project, data }: { project: ProjectState; data: PageData } = $props()
</script>

<Tabs.Content value="tasks">
    <CreationModal
        {project}
        {data}
    />
    <KanbanBoard {project} />
</Tabs.Content>
