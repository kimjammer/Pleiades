<script lang="ts">
    import Chart, { type ChartData } from "$lib/components/Chart.svelte"
    import type { ProjectState, Session } from "$lib/project_state.svelte.js"
    import {
        fromAbsolute,
        getLocalTimeZone,
        isSameDay,
        type ZonedDateTime,
    } from "@internationalized/date"

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

    //Absolutely horrible slow function, but is always timezone & time weirdness safe
    function daysBetweenDates(start: ZonedDateTime, end: ZonedDateTime): number {
        if (start.compare(end) > 0) {
            return 0
        }

        let counter = 0
        let crrDate = start
        while (!isSameDay(crrDate, end)) {
            counter++
            crrDate = crrDate.add({ days: 1 })
        }
        return counter
    }

    $effect(() => {
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
        let startTimestamp = project.created
        //Find last due date
        let endTimeStamp = tasks.reduce((latest, task) => {
            if (task.dueDate > latest) {
                return task.dueDate
            } else {
                return latest
            }
        }, tasks[0].dueDate)

        //Loop through each day in the range
        let labels = []
        let idealLine = []
        let actualLine = []

        let startDate = fromAbsolute(startTimestamp, getLocalTimeZone())
        startDate = startDate.set({
            hour: 0,
            minute: 0,
            second: 0,
        })
        let crrDate = fromAbsolute(startTimestamp, getLocalTimeZone())
        crrDate = crrDate.set({
            hour: 11,
            minute: 59,
            second: 59,
        })
        let endDate = fromAbsolute(endTimeStamp, getLocalTimeZone())
        endDate = endDate.add({ days: 1 })

        //Iterate over days from start to end date
        while (!isSameDay(crrDate, endDate)) {
            labels.push(crrDate.toDate().toLocaleDateString())

            //For each task not past due date, calculate ideal progress towards time estimate
            let idealTime = tasks.reduce((total, task) => {
                let taskDueDate = fromAbsolute(task.dueDate, getLocalTimeZone())
                if (taskDueDate.compare(crrDate) > 0) {
                    return (
                        total +
                        (task.timeEstimate / daysBetweenDates(startDate, taskDueDate)) *
                            daysBetweenDates(startDate, crrDate)
                    )
                } else {
                    return total + task.timeEstimate
                }
            }, 0)
            idealLine.push(idealTime)

            //For each session, sum time spent on tasks
            let actualTime = sessions.reduce((total, session) => {
                let sessionEndDate = fromAbsolute(session.endTime, getLocalTimeZone())
                //If session was completed on or before current date
                if (sessionEndDate.compare(crrDate) < 0) {
                    return total + (session.endTime - session.startTime)
                }

                return total
            }, 0)
            actualLine.push(actualTime)

            //Increment date by 1 day
            crrDate = crrDate.add({ days: 1 })
        }

        //Transform time from millis to hours
        idealLine = idealLine.map(time => time / (1000 * 60 * 60))
        actualLine = actualLine.map(time => time / (1000 * 60 * 60))

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

    let options = {
        scales: {
            y: {
                title: {
                    display: true,
                    text: "Hours",
                },
            },
            x: {
                title: {
                    display: true,
                    text: "Date",
                },
            },
        },
        plugins: {
            title: {
                display: true,
                text: "Progress Chart",
            },
        },
    }
</script>

<div style="max-width: 500px; aspect-ratio: 2;">
    {#if dataAvailable}
        <Chart
            type="line"
            data={$state.snapshot(data) as any}
            {options}
        />
    {:else}
        <div
            class="flex w-full flex-col items-center justify-center rounded-xl border-4
                    border-primary p-5"
        >
            <p class="leading-7 [&:not(:first-child)]:mt-6">
                Create a task with a due date and time estimate to see the burndown chart.
            </p>
        </div>
    {/if}
</div>
