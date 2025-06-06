<script lang="ts">
    import { CalendarDate, type DateValue, getLocalTimeZone } from "@internationalized/date"
    import { Button } from "bits-ui"
    import { Badge } from "$lib/components/ui/badge"
    import { Calendar } from "$lib/components/ui/calendar"
    import * as Popover from "$lib/components/ui/popover"
    import { CalendarDays, X, Plus } from "lucide-svelte"
    import type { ProjectState, Poll } from "$lib/project_state.svelte.js"

    let { project, poll }: { project: ProjectState; poll: Poll } = $props()

    let dueDateInit = new Date(poll.dueDate)
    let dueDayPreComp =
        poll.dueDate === 0
            ? undefined
            : new CalendarDate(
                dueDateInit.getFullYear(),
                dueDateInit.getMonth() + 1,
                dueDateInit.getDate(),
            )

    let value = $state<DateValue | undefined>(dueDayPreComp)
    let contentRef = $state<HTMLElement | null>(null)

    async function handleEdit() {
        //Send to server
        //Component is updated when server updates the project state and replies
        const timestamp = value?.toDate(getLocalTimeZone()).getTime() ?? 0
        project.updateInProject(`Polls[Id=${poll.id}].DueDate`, timestamp)
    }
    async function handleDelete(e: Event) {
        e.stopPropagation()
        project.updateInProject(`Polls[Id=${poll.id}].DueDate`, 0)
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
                {#if poll.dueDate === 0}
                    Due Date
                    <Plus size="15" />
                {:else}
                    {new Date(poll.dueDate).toLocaleString("default", { month: "long" })}

                    {new Date(poll.dueDate).toLocaleString("default", { day: "numeric" })}
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
                bind:value
                onValueChange={handleEdit}
                class="rounded-md border"
        />
    </Popover.Content>
</Popover.Root>
