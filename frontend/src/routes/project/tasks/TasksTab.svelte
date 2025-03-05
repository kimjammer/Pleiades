<script lang="ts">
    import * as Tabs from "$lib/components/ui/tabs"
    import type { ProjectState } from "$lib/project_state.svelte"

    let { project }: { project: ProjectState } = $props()
</script>

<Tabs.Content value="tasks">
    <div></div>
</Tabs.Content>
