<script lang="ts">
    import { Button } from "bits-ui"
    import { Badge } from "$lib/components/ui/badge"
    import { Calendar } from "$lib/components/ui/calendar"
    import * as Popover from "$lib/components/ui/popover"
    import type { ProjectState, Task } from "$lib/project_state.svelte.js"
    import {
        type DateValue,
        fromAbsolute,
        getLocalTimeZone,
        now,
        toCalendarDate,
        toTime,
        ZonedDateTime,
    } from "@internationalized/date"
    import { CalendarDays, Plus, X } from "lucide-svelte"
    import { Input } from "$lib/components/ui/input"
    import { Separator } from "$lib/components/ui/separator"
    import { onMount } from "svelte"

    let { project, task }: { project: ProjectState; task: Task } = $props()

    let time = $state<string>("")
    let date = $state<DateValue | undefined>(undefined)
    let contentRef = $state<HTMLElement | null>(null)

    onMount(() => {
        if (task.dueDate === 0) return

        //Convert due date from timestamp to CalendarDate
        let serverDueDate = fromAbsolute(task.dueDate, getLocalTimeZone())
        console.log("Server timestamp", task.dueDate)
        console.log("Server due date", serverDueDate)
        date = toCalendarDate(serverDueDate)
        let dueTime = toTime(serverDueDate)
        console.log("Server due time", dueTime)
        time = `${dueTime.hour.toString().padStart(2, "0")}:${dueTime.minute.toString().padStart(2, "0")}`
        console.log(time)
    })

    async function handleEdit() {
        if (!date || time === "") return

        //Send to server
        //Component is updated when server updates the project state and replies
        const today = now(getLocalTimeZone())
        const hour = time.split(":").map(Number)[0]
        const minute = time.split(":").map(Number)[1]
        const dueDate = new ZonedDateTime(
            date.year,
            date.month,
            date.day,
            getLocalTimeZone(),
            today.offset,
            hour,
            minute,
        )

        const timestamp = dueDate.toDate().getTime()
        console.log("New timestamp ", timestamp)
        project.updateInProject(`Tasks[Id=${task.id}].DueDate`, timestamp)
    }

    async function handleDelete(e: Event) {
        e.stopPropagation()
        project.updateInProject(`Tasks[Id=${task.id}].DueDate`, 0)
    }
</script>

<Popover.Root>
    <Popover.Trigger>
        <Badge
            class="pl-2 pr-0.5"
            onclick={handleEdit}
        >
            <div class="flex items-center gap-1">
                <CalendarDays size="12" />
                {#if task.dueDate === 0}
                    Due Date
                    <Plus size="15" />
                {:else}
                    {new Date(task.dueDate).toLocaleString("default", { month: "long" })}

                    {new Date(task.dueDate).toLocaleString("default", { day: "numeric" })}
                    <Button.Root
                        class="
        inline-flex items-center justify-center whitespace-nowrap rounded-full text-sm font-medium ring-offset-background transition-colors hover:bg-accent hover:text-accent-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2
        disabled:pointer-events-none disabled:opacity-50"
                        onclick={handleDelete}
                    >
                        <X size="15" />
                    </Button.Root>
                {/if}
            </div>
        </Badge>
    </Popover.Trigger>
    <Popover.Content
        bind:ref={contentRef}
        class="w-auto p-0"
    >
        <Calendar
            type="single"
            bind:value={date}
            onValueChange={handleEdit}
            class="rounded-md border"
        />
        <Separator class="my-3" />
        <div class="m-3">
            <Input
                type="time"
                bind:value={time}
                onchange={handleEdit}
            ></Input>
        </div>
    </Popover.Content>
</Popover.Root>
