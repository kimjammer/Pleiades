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
        <div
            class="flex w-full flex-col items-center justify-center rounded-xl border-4
                    border-primary p-5"
        >
            <p class="leading-7 [&:not(:first-child)]:mt-6">
                Record sessions to see the time breakdown pie chart.
            </p>
        </div>
    {/if}
</div>
