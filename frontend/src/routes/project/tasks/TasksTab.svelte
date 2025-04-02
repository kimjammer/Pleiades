<script lang="ts">
    import * as Tabs from "$lib/components/ui/tabs"
    import { type ProjectState, Task } from "$lib/project_state.svelte"
    import DueDate from "$lib/components/DueDate.svelte"
    import CreationModal from "./CreationModal.svelte"
    import type { PageData } from "./$types"
    import TimeEstimate from "$lib/components/TimeEstimate.svelte"
    import Timer from "$lib/components/Timer.svelte"
    import Chart from "$lib/components/Chart.svelte"
    import BurndownChart from "$lib/components/BurndownChart.svelte"

    let { project, data }: { project: ProjectState; data: PageData } = $props()
    console.log("data in TasksTab: " + data)

    //Testing
    const task: Task = project.tasks[0] ?? new Task()
</script>

<Tabs.Content value="tasks">
    <CreationModal
        {project}
        {data}
    />
    <div>
        {task.title}
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
    <BurndownChart {project} />
</Tabs.Content>
