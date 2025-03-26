<script lang="ts">
    import { type DateValue, getLocalTimeZone } from "@internationalized/date"
    import { Button } from "bits-ui"
    import { Badge } from "$lib/components/ui/badge"
    import { Calendar } from "$lib/components/ui/calendar"
    import * as Popover from "$lib/components/ui/popover"
    import { CalendarDays, X, Plus } from "lucide-svelte"
    import type { ProjectState } from "$lib/project_state.svelte"

    let { dueDate, project, taskID }: { dueDate: number; project: ProjectState; taskID: string } =
        $props()

    let value = $state<DateValue | undefined>()
    let contentRef = $state<HTMLElement | null>(null)

    async function handleEdit() {
        //Send to server
        //Component is updated when server updates the project state and replies
        const timestamp = value?.toDate(getLocalTimeZone()).getTime()
        //project.updateInProject(`Tasks[Id=${taskID}].DueDate`, timestamp)

        //TODO:remove
        dueDate = timestamp || 0
    }
    async function handleDelete(e: Event) {
        e.stopPropagation()
        //project.updateInProject(`Tasks[Id=${taskID}].DueDate`, 0)

        //TODO:remove
        dueDate = 0
        value = undefined
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
                {#if dueDate === 0}
                    Add Due Date
                    <Plus size="15" />
                {:else}
                    {new Date(dueDate).toLocaleString("default", { month: "long" })}

                    {new Date(dueDate).toLocaleString("default", { day: "numeric" })}
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
