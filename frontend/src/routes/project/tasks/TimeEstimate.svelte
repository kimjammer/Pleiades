<script lang="ts">
    import { Button } from "bits-ui"
    import { Badge } from "$lib/components/ui/badge"
    import * as Popover from "$lib/components/ui/popover"
    import { Clock, X, Plus } from "lucide-svelte"
    import { Input } from "$lib/components/ui/input"
    import { Label } from "$lib/components/ui/label"
    import type { ProjectState, Task } from "$lib/project_state.svelte.js"

    let { project, task }: { project: ProjectState; task: Task } = $props()

    let value = $state<string | undefined>()
    let contentRef = $state<HTMLElement | null>(null)
    let isOpen = $state(false)

    $effect(() => {
        value = task.timeEstimate === 0 ? "" : (task.timeEstimate / 1000 / 60 / 60).toString()
    })

    async function handleEdit() {
        isOpen = false

        //Attempt to convert value to number
        let estHours = Number(value)
        if (isNaN(estHours)) {
            estHours = 0
        }
        //Round number to nearest tenth
        estHours = Math.round(estHours * 10) / 10
        let estMillis = estHours * 60 * 60 * 1000

        value = estHours.toString()
        //Send to server
        //Component is updated when server updates the project state and replies
        project.updateInProject(`Tasks[Id=${task.id}].TimeEstimate`, estMillis)
    }

    async function handleDelete(e: Event) {
        e.stopPropagation()
        isOpen = false

        project.updateInProject(`Tasks[Id=${task.id}].TimeEstimate`, 0)
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
                {#if task.timeEstimate === 0}
                    Add Estimated Time
                    <Plus size="15" />
                {:else}
                    {task.timeEstimate / 1000 / 60 / 60} Hrs
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
