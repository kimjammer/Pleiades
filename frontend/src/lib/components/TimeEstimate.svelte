<script lang="ts">
    import { Button } from "bits-ui"
    import { Badge } from "$lib/components/ui/badge"
    import * as Popover from "$lib/components/ui/popover"
    import { Clock, X, Plus } from "lucide-svelte"
    import { Input } from "$lib/components/ui/input"
    import { Label } from "$lib/components/ui/label"
    import type { ProjectState } from "$lib/project_state.svelte"

    let {
        timeEstimate,
        project,
        taskID,
    }: { timeEstimate: number; project: ProjectState; taskID: string } = $props()

    let value = $state<string | undefined>()
    let contentRef = $state<HTMLElement | null>(null)
    let isOpen = $state(false)

    async function handleEdit() {
        isOpen = false

        //Attempt to convert value to number
        let valueNum = Number(value)
        if (isNaN(valueNum)) {
            valueNum = 0
        }
        //Round number to nearest tenth
        valueNum = Math.round(valueNum * 10) / 10

        value = valueNum.toString()
        //Send to server
        //Component is updated when server updates the project state and replies
        //project.updateInProject(`Tasks[Id=${taskID}].TimeEstimate`, value)

        //TODO:remove
        timeEstimate = valueNum
    }

    async function handleDelete(e: Event) {
        e.stopPropagation()
        isOpen = false

        //project.updateInProject(`Tasks[Id=${taskID}].TimeEstimate`, 0)

        //TODO:remove
        timeEstimate = 0
        value = undefined
    }
</script>

<Popover.Root bind:open={isOpen}>
    <Popover.Trigger>
        <Badge
            class="pl-2 pr-0.5"
            onclick={handleEdit}
        >
            <div class="flex items-center gap-1">
                <Clock size="12" />
                {#if timeEstimate === 0}
                    Add Estimated Time
                    <Plus size="15" />
                {:else}
                    {timeEstimate} Hrs
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
        <div class="flex items-center gap-3 p-2">
            <Input
                id="estimate"
                bind:value
                onchange={handleEdit}
            />
            <Label for="estimate">Hours</Label>
        </div>
    </Popover.Content>
</Popover.Root>
