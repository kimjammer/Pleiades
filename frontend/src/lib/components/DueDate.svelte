<script lang="ts">
    import { type DateValue, getLocalTimeZone, today } from "@internationalized/date"
    import { Button } from "bits-ui"
    import { Badge } from "$lib/components/ui/badge"
    import { Calendar } from "$lib/components/ui/calendar"
    import * as Popover from "$lib/components/ui/popover"
    import { CalendarDays, X } from "lucide-svelte"

    let { dueDate = new Date() }: { dueDate: Date } = $props()

    let value = $state<DateValue | undefined>(today(getLocalTimeZone()))
    let contentRef = $state<HTMLElement | null>(null)

    async function handleEdit() {
        //Send to server
        //Component is updated when server updates the project state and replies
    }
    async function handleDelete() {}
</script>

<Popover.Root>
    <Popover.Trigger>
        <Badge
            class="pl-2 pr-0.5"
            onclick={handleEdit}
        >
            <div class="flex items-center gap-1">
                <CalendarDays size="12" />
                {dueDate.toLocaleString("default", { month: "long" })}

                {dueDate.toLocaleString("default", { day: "numeric" })}
                <Button.Root
                    class="
        inline-flex items-center justify-center whitespace-nowrap rounded-full text-sm font-medium ring-offset-background transition-colors hover:bg-accent hover:text-accent-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2
        disabled:pointer-events-none disabled:opacity-50"
                    onclick={handleDelete}
                >
                    <X size="15" />
                </Button.Root>
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
