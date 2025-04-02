<script lang="ts">
    import Chart from "$lib/components/Chart.svelte"
    import type { ProjectState } from "$lib/project_state.svelte"
    import { onMount } from "svelte"

    let { project }: { project: ProjectState } = $props()

    let data = {
        labels: [1, 2, 3, 4, 5],
        datasets: [
            {
                label: "test",
                data: [1, 2, 3, 4, 5],
            },
        ],
    }

    onMount(() => {})
</script>

<div>
    <Chart
        type="line"
        {data}
    />
</div>
