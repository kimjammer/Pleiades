<script lang="ts">
    import * as Tabs from "$lib/components/ui/tabs"
    import type { ProjectState } from "$lib/project_state.svelte"
    import DueDate from "$lib/components/DueDate.svelte"
    import CreationModal from "./CreationModal.svelte"
    import type { PageData } from "./$types"

    let { project, data }: { project: ProjectState; data: PageData } = $props()

    //Testing
    const testDueDate = undefined
</script>

<Tabs.Content value="tasks">
    <CreationModal
        {project}
        {data}
    />
    <div>
        <DueDate dueDate={testDueDate} />
    </div>
</Tabs.Content>
