<script lang="ts">
    import * as Tabs from "$lib/components/ui/tabs"
    import type { ProjectState } from "$lib/project_state.svelte"
    import DueDate from "$lib/components/DueDate.svelte"
    import CreationModal from "./CreationModal.svelte"
    import type { PageData } from "./$types"
    import TimeEstimate from "$lib/components/TimeEstimate.svelte"

    let { project, data }: { project: ProjectState; data: PageData } = $props()

    //Testing
    const testDueDate = undefined
    const testTime = 0
</script>

<Tabs.Content value="tasks">
    <CreationModal
        {project}
        {data}
    />
    <div>
        <DueDate dueDate={testDueDate} />
        <TimeEstimate timeEstimate={testTime} />
    </div>
</Tabs.Content>
