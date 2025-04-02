<script lang="ts">
    import BurndownChart from "$lib/components/BurndownChart.svelte"
    import * as Tabs from "$lib/components/ui/tabs/index.js"
    import type { ProjectState } from "$lib/project_state.svelte"

    let { project }: { project: ProjectState } = $props()
</script>

<Tabs.Content value="stats">
    <BurndownChart {project} />
</Tabs.Content>
