<script lang="ts">
    import BurndownChart from "./BurndownChart.svelte"
    import * as Tabs from "$lib/components/ui/tabs/index.js"
    import type { ProjectState } from "$lib/project_state.svelte"
    import TimePiechart from "./TimePiechart.svelte"

    let { project }: { project: ProjectState } = $props()
</script>

<Tabs.Content value="stats">
    <div class="flex flex-col gap-5">
        <BurndownChart {project} />
        <TimePiechart {project} />
    </div>
</Tabs.Content>
