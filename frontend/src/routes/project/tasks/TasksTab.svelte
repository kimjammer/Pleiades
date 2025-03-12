<script lang="ts">
    import * as Tabs from "$lib/components/ui/tabs"
    import type { ProjectState } from "$lib/project_state.svelte"
    import DueDate from "$lib/components/DueDate.svelte"
    import CreationModal from "./CreationModal.svelte"

    let { project }: { project: ProjectState } = $props()

    //Testing
    const testDueDate = new Date()
</script>

<Tabs.Content value="tasks">
    <CreationModal {project} />
    <div>
        <DueDate dueDate={testDueDate} />
    </div>
</Tabs.Content>
