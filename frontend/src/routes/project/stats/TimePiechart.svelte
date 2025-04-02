<script lang="ts">
    import Chart, { type ChartData } from "$lib/components/Chart.svelte"
    import type { ProjectState } from "$lib/project_state.svelte"
    import { onMount } from "svelte"

    let { project }: { project: ProjectState } = $props()

    let dataAvailable = $state(false)

    let data: ChartData<"pie"> = $state({
        labels: [],
        datasets: [
            {
                label: "ideal",
                data: [],
            },
            {
                label: "actual",
                data: [],
            },
        ],
    })

    onMount(() => {
        // Mapping of user id to hours spent on all tasks in this project
        const accumulatedUserTime = Object.fromEntries(project.users.map(user => [user.id, 0]))

        for (const task of project.tasks) {
            for (const session of task.sessions) {
                accumulatedUserTime[session.user] +=
                    (session.endTime - session.endTime) / 1000 / 60 / 60
            }
        }

        /* TODO: If no tasks, return
        if (tasks.length === 0) {
            dataAvailable = false
            return
        } else {
            dataAvailable = true
        }*/
    })
</script>

<div>
    {#if dataAvailable}
        <Chart
            type="pie"
            {data}
        />
    {:else}
        <p class="leading-7 [&:not(:first-child)]:mt-6">
            Create a task with a due date and time estimate to see the burndown chart.
        </p>
    {/if}
</div>
