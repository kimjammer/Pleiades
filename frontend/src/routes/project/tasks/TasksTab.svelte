<script lang="ts">
    import * as Tabs from "$lib/components/ui/tabs"
    import { type ProjectState, Task } from "$lib/project_state.svelte"
    import DueDate from "$lib/components/DueDate.svelte"
    import CreationModal from "./CreationModal.svelte"
    import type { PageData } from "./$types"
    import TimeEstimate from "$lib/components/TimeEstimate.svelte"
    import Timer from "$lib/components/Timer.svelte"

    let { project, data }: { project: ProjectState; data: PageData } = $props()
    console.log("data in TasksTab: " + data)

    //Testing
    const task: Task = new Task()
    task.id = crypto.randomUUID()
    task.title = "My epic task"
    task.description = "super cool description"
    task.dueDate = 0
    task.kanbanColumn = ""
    task.timeEstimate = 0
    task.completed = false
    task.sessions = []
    task.assignees = []
</script>

<Tabs.Content value="tasks">
    <CreationModal
        {project}
        {data}
    />
    <div>
        <DueDate
            {project}
            {task}
        />
        <TimeEstimate
            {project}
            {task}
        />
        <Timer
            {project}
            {task}
        />
    </div>
</Tabs.Content>
