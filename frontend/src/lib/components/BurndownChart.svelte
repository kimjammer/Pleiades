<script lang="ts">
    import Chart, { type ChartData } from "$lib/components/Chart.svelte"
    import type { ProjectState, Session } from "$lib/project_state.svelte"

    let { project }: { project: ProjectState } = $props()

    let dataAvailable = $state(false)

    let data: ChartData<"line"> = $state({
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

    $effect(() => {
        console.log("Updating chart")

        //TODO: Refactor this terrible code with sensible time zone handling

        //Get all tasks with due dates and time estimates
        let tasks = project.tasks.filter(task => task.dueDate && task.timeEstimate)

        //If no tasks, return
        if (tasks.length === 0) {
            dataAvailable = false
            return
        } else {
            dataAvailable = true
        }

        //Get all completed sessions spent on tasks with due dates and time estimates
        let sessions: Session[] = []
        tasks.forEach(task => {
            task.sessions.forEach(session => {
                if (session.startTime != 0 && session.endTime != 0) {
                    sessions.push(session)
                }
            })
        })

        //Get range for graph
        let start = project.created
        //Find last due date
        let end = tasks.reduce((latest, task) => {
            if (task.dueDate > latest) {
                return task.dueDate
            } else {
                return latest
            }
        }, tasks[0].dueDate)

        end += 1000 * 60 * 60 * 24 //Add one day to end date

        //Loop through each day in the range
        let labels = []
        let idealLine = []
        let actualLine = []
        let currentDate = new Date(start)
        let endDate = new Date(end)

        while (currentDate <= endDate) {
            labels.push(currentDate.toLocaleDateString())

            //For each task not past due date, calculate idea progress towards time estimate
            let idealTime = tasks.reduce((total, task) => {
                if (task.dueDate >= currentDate.getTime()) {
                    return (
                        total +
                        (task.timeEstimate / (task.dueDate - start)) *
                            (currentDate.getTime() - start)
                    )
                } else {
                    return total + task.timeEstimate
                }
            }, 0)
            idealLine.push(idealTime)

            //For each session, sum time spent on tasks
            let actualTime = sessions.reduce((total, session) => {
                //If session was completed on or before current date
                if (session.endTime <= currentDate.getTime()) {
                    return total + (session.endTime - session.startTime)
                }

                return total
            }, 0)
            actualLine.push(actualTime)

            //Increment date by 1 day
            currentDate.setDate(currentDate.getDate() + 1)
        }

        data = {
            labels: labels,
            datasets: [
                {
                    label: "ideal",
                    data: idealLine,
                },
                {
                    label: "actual",
                    data: actualLine,
                },
            ],
        }
    })
</script>

<div>
    {#if dataAvailable}
        <Chart
            type="line"
            data={$state.snapshot(data) as any}
        />
    {:else}
        <p class="leading-7 [&:not(:first-child)]:mt-6">
            Create a task with a due date and time estimate to see the burndown chart.
        </p>
    {/if}
</div>
