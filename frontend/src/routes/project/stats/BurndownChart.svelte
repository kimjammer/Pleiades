<script lang="ts">
    import Chart, { type ChartData } from "$lib/components/Chart.svelte"
    import {
        type ProjectState,
        type Session,
        Task,
        UserInProject,
    } from "$lib/project_state.svelte.js"
    import {
        fromAbsolute,
        getLocalTimeZone,
        isSameDay,
        type ZonedDateTime,
    } from "@internationalized/date"
    import { Label } from "$lib/components/ui/label"
    import { Switch } from "$lib/components/ui/switch"

    let { project }: { project: ProjectState } = $props()

    let dataAvailable = $state(false)
    let userBreakout = $state(false)
    let taskBreakout = $state(false)

    let data: ChartData<"line"> = $state({
        labels: [],
        datasets: [],
    })

    //Absolutely horrible slow function, but is always timezone and time weirdness safe
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

    function aggregateTotalIdealTime(
        tasks: Task[],
        startDate: ZonedDateTime,
        crrDate: ZonedDateTime,
        endDate: ZonedDateTime,
    ) {
        let idealLine: number[] = []
        while (!isSameDay(crrDate, endDate)) {
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

            //Increment date by 1 day
            crrDate = crrDate.add({ days: 1 })
        }
        idealLine = idealLine.map(time => time / (1000 * 60 * 60))

        return {
            label: "ideal",
            data: idealLine,
        }
    }

    function aggregateTaskIdealTime(
        task: Task,
        startDate: ZonedDateTime,
        crrDate: ZonedDateTime,
        endDate: ZonedDateTime,
    ) {
        let taskLine: number[] = []
        while (!isSameDay(crrDate, endDate)) {
            //For each task not past due date, calculate ideal progress towards time estimate
            let taskDueDate = fromAbsolute(task.dueDate, getLocalTimeZone())
            let taskTime = 0
            if (taskDueDate.compare(crrDate) > 0) {
                taskTime =
                    (task.timeEstimate / daysBetweenDates(startDate, taskDueDate)) *
                    daysBetweenDates(startDate, crrDate)
            } else {
                taskTime = task.timeEstimate
            }
            taskLine.push(taskTime)

            //Increment date by 1 day
            crrDate = crrDate.add({ days: 1 })
        }
        taskLine = taskLine.map(time => time / (1000 * 60 * 60))

        return {
            label: task.title,
            stack: "tasks",
            data: taskLine,
        }
    }

    function aggregateTotalActualTime(
        sessions: Session[],
        startDate: ZonedDateTime,
        crrDate: ZonedDateTime,
        endDate: ZonedDateTime,
    ) {
        let actualLine: number[] = []
        while (!isSameDay(crrDate, endDate)) {
            //For each session, sum time spent on tasks
            let actualTime = sessions.reduce((total, session) => {
                let sessionEndDate = fromAbsolute(session.endTime, getLocalTimeZone())
                //If the session was completed on or before the current date
                if (sessionEndDate.compare(crrDate) < 0) {
                    return total + (session.endTime - session.startTime)
                }

                return total
            }, 0)
            actualLine.push(actualTime)

            //Increment date by 1 day
            crrDate = crrDate.add({ days: 1 })
        }
        actualLine = actualLine.map(time => time / (1000 * 60 * 60))

        return {
            label: "actual",
            data: actualLine,
        }
    }

    function aggregateUserActualTime(
        sessions: Session[],
        startDate: ZonedDateTime,
        crrDate: ZonedDateTime,
        endDate: ZonedDateTime,
        user: UserInProject,
    ) {
        let userLine: number[] = []
        while (!isSameDay(crrDate, endDate)) {
            //For each session, sum time spent on tasks
            let actualTime = sessions.reduce((total, session) => {
                let sessionEndDate = fromAbsolute(session.endTime, getLocalTimeZone())
                //If the session was completed on or before the current date
                if (session.user == user.id && sessionEndDate.compare(crrDate) < 0) {
                    return total + (session.endTime - session.startTime)
                }

                return total
            }, 0)
            userLine.push(actualTime)

            //Increment date by 1 day
            crrDate = crrDate.add({ days: 1 })
        }
        userLine = userLine.map(time => time / (1000 * 60 * 60))

        return {
            label: `${user.firstName} ${user.lastName}`,
            stack: "users",
            data: userLine,
        }
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

        //Aggregate data for each line
        let datasets = []

        if (taskBreakout) {
            for (let task of tasks) {
                datasets.push(aggregateTaskIdealTime(task, startDate, crrDate, endDate))
            }
        } else {
            datasets.push(aggregateTotalIdealTime(tasks, startDate, crrDate, endDate))
        }

        if (userBreakout) {
            for (let user of project.users) {
                datasets.push(aggregateUserActualTime(sessions, startDate, crrDate, endDate, user))
            }
        } else {
            datasets.push(aggregateTotalActualTime(sessions, startDate, crrDate, endDate))
        }

        //Iterate over days from start to end date
        let labels = []
        while (!isSameDay(crrDate, endDate)) {
            labels.push(crrDate.toDate().toLocaleDateString())
            //Increment date by 1 day
            crrDate = crrDate.add({ days: 1 })
        }

        data = {
            labels: labels,
            datasets: datasets,
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
        stack: true,
    }
</script>

<div class="flex aspect-video max-w-[500px]">
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
    <div class="flex flex-col justify-center gap-2">
        <div class="flex items-center gap-1">
            <Switch
                id="user-breakout"
                bind:checked={userBreakout}
            />
            <Label for="user-breakout">Breakout Users</Label>
        </div>
        <div class="flex items-center gap-1">
            <Switch
                id="task-breakout"
                bind:checked={taskBreakout}
            />
            <Label for="task-breakout">Breakout Tasks</Label>
        </div>
    </div>
</div>
