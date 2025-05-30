<script lang="ts">
    import { ProjectState, Task, mouse } from "$lib/project_state.svelte"
    import * as ContextMenu from "$lib/components/ui/context-menu"
    import { Accordion } from "bits-ui"
    import DueDate from "./DueDate.svelte"
    import TimeEstimate from "./TimeEstimate.svelte"
    import Timer from "./Timer.svelte"
    import { Button } from "$lib/components/ui/button"

    let { project, task }: { project: ProjectState; task: Task } = $props()

    let progress: number | undefined = $state()

    let assignees = $derived(
        task.assignees
            .map(id => project.users.find(user => user.id === id))
            .map(user => `${user?.firstName} ${user?.lastName}`),
    )

    let assigned = $derived(task.assignees.includes(localStorage.myId))

    let dragging = $state(false)
    let startX: number = $state(0)
    let startY: number = $state(0)

    let card: HTMLDivElement

    function changeTaskMembership() {
        if (task.assignees.includes(localStorage.myId)) {
            //if an assignee of the task, leave
            task.assignees = task.assignees.filter(id => id !== localStorage.myId)
        } else {
            //if not an assignee, join
            task.assignees.push(localStorage.myId)
        }
        //update in project
        project.updateInProject(`Tasks[Id=${task.id}].Assignees`, task.assignees)
    }

    //for color coding the progress bar
    function getProgressColor(progress: number) {
        if (progress < 20) return "#FF8C00"
        if (progress < 50) return "#facc15"
        if (progress < 70) return "#9ACD32"
        return "#22c55e" // green-500
    }

    function getTitleColor(column: string) {
        if (column == "" && progress == 0) {
            return null
        } else if (column == "") return "#cc7a00"
        else if (column == "progress") return "#99cc00"
        else return "#008000"
    }
</script>

{#snippet content()}
    <Accordion.Root type="single">
        <Accordion.Item
            value="1"
            class="group"
        >
            <Accordion.Header>
                <Accordion.Trigger class="w-full">
                    <h1
                        class="text-[1.2em] font-bold dark:text-white"
                        style={`color: ${getTitleColor(task.kanbanColumn)}`}
                    >
                        {task.title}
                    </h1>
                    <p>
                        {task.description}
                    </p>
                </Accordion.Trigger>
            </Accordion.Header>
            <Accordion.Content
                class="overflow-hidden data-[state=closed]:animate-accordion-up data-[state=open]:animate-accordion-down"
                onmousedown={e => {
                    e.stopPropagation()
                }}
                onmouseup={e => {
                    e.stopPropagation()
                }}
            >
                <div class="grid grid-cols-3">
                    <div class="col-span-1">
                        <Timer
                            {project}
                            {task}
                            bind:progress
                        />
                    </div>
                    <div class="col-span-1">
                        {#each assignees as assignee}
                            <p>{assignee}</p>
                        {/each}
                    </div>
                    <div class="col-span-1 flex flex-col items-end px-2">
                        <DueDate
                            {project}
                            {task}
                        />
                        <TimeEstimate
                            {project}
                            {task}
                        />
                    </div>
                    <div>
                        {#if assigned}
                            <Button
                                variant="destructive"
                                onclick={changeTaskMembership}>Leave Task</Button
                            >
                        {:else}
                            <Button onclick={changeTaskMembership}>Join Task</Button>
                        {/if}
                    </div>
                </div>
            </Accordion.Content>
        </Accordion.Item>
    </Accordion.Root>

    <div class="relative h-2 w-full overflow-hidden">
        <div
            class="absolute h-full w-full transition-all"
            style={`transform: translateX(-${100 - Math.min(progress ?? 0, 100)}%);
                                    background-color: ${getProgressColor(progress)}`}
        ></div>
    </div>
{/snippet}

<ContextMenu.Root>
    <ContextMenu.Trigger>
        <div
            bind:this={card}
            class="task bg-muted"
            role="presentation"
            onmousedown={() => {
                startX = mouse.x
                startY = mouse.y
                dragging = true
            }}
            onmouseup={e => {
                setTimeout(() => {
                    dragging = false
                }, 100)

                card.hidden = true
                let elemBelow = document.elementFromPoint(e.clientX, e.clientY)
                card.hidden = false

                let elt = elemBelow?.closest(".kanban-column")

                console.log(elt)
                if (elt != null) {
                    project.updateInProject(`Tasks[Id=${task.id}].KanbanColumn`, elt.id)
                }
            }}
            style={dragging ? `top:${mouse.y - startY}px;left:${mouse.x - startX}px` : ``}
        >
            {@render content()}
        </div>
    </ContextMenu.Trigger>
    <ContextMenu.Content>
        <ContextMenu.Item
            onclick={() => {
                project.deleteInProject(`Tasks[Id=${task.id}]`)
            }}>Delete</ContextMenu.Item
        >
    </ContextMenu.Content>
</ContextMenu.Root>

<style>
    .task {
        margin: 0.5em;
        border-radius: 0.5em;
        position: relative;
        overflow: hidden;
    }
</style>
