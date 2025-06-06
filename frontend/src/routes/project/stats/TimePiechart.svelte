<script lang="ts">
    import Chart, { type ChartData } from "$lib/components/Chart.svelte"
    import type { ProjectState } from "$lib/project_state.svelte"
    import { onMount } from "svelte"

    let { project }: { project: ProjectState } = $props()

    let dataAvailable = $state(false)

    let data = $state({
        labels: project.users.map(user => user.firstName + " " + user.lastName),
        datasets: [
            {
                data: [],
            },
        ],
    } as ChartData<"pie">)

    onMount(() => {
        // Mapping of user id to hours spent on all tasks in this project
        const accumulatedUserTime = Object.fromEntries(project.users.map(user => [user.id, 0]))

        for (const task of project.tasks) {
            for (const session of task.sessions) {
                //Ignore sessions currently in progress
                if (session.endTime === 0) continue

                accumulatedUserTime[session.user] +=
                    (session.endTime - session.startTime) / 1000 / 60 / 60
            }
        }

        // If no tasks, return
        if (Object.values(accumulatedUserTime).every(time => !time)) {
            dataAvailable = false
            return
        } else {
            dataAvailable = true
        }

        data.datasets[0].data = Object.values(accumulatedUserTime)
    })
</script>

<div style="max-width: 500px; aspect-ratio: 1;">
    {#if dataAvailable}
        <Chart
            type="pie"
            data={$state.snapshot(data) as any}
            options={{
                plugins: {
                    tooltip: {
                        callbacks: { label: ({ formattedValue }) => formattedValue + " hr" },
                    },
                },
            }}
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
