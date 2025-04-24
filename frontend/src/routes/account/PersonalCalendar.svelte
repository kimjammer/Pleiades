<script lang="ts">
    import { onMount } from "svelte"
    import { Task } from "$lib/project_state.svelte"
    import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"
    import { toast } from "svelte-sonner"
    import { Button } from "$lib/components/ui/button"
    import Calendar from "../project/calendar/Calendar.svelte"
    import * as DropdownMenu from "$lib/components/ui/dropdown-menu/index.js"
    import { Label } from "$lib/components/ui/label"
    import { Switch } from "$lib/components/ui/switch"

    let year = $state(new Date().getFullYear())
    let month = $state(new Date().getMonth() + 1)
    function handleYearChange(e: Event) {
        const input = e.target as HTMLInputElement
        year = parseInt(input.value) || new Date().getFullYear()
    }

    function handleMonthChange(e: Event) {
        const input = e.target as HTMLInputElement
        month = parseInt(input.value) || new Date().getMonth() + 1
    }
    onMount(async () => {
        tasks = getTasks()
        await tasks
        await mapTasks()
    })

    let tasks: Promise<Task[]> = $state(new Promise(() => {}))
    let projectNames: string[] = $state([])

    async function getTasks() {
        return new Promise<Task[]>((resolve, reject) => {
            fetch(PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/getUserTasks", {
                method: "GET",
                mode: "cors",
                credentials: "include",
                headers: { "Content-Type": "application/json" },
            })
                .then(res => {
                    return res.json()
                })
                .then(data => {
                    data = data as {
                        success: boolean
                        tasks: Task[]
                        projectNames: string[]
                    }
                    if (data.success) {
                        toast.success("User tasks fetched")
                        projectNames = [...data.projectNames]

                        resolve(data.tasks as Task[])
                    } else {
                        toast.error("Failed to load user tasks")
                        reject()
                    }
                })
        })
    }

    let projectTaskMap = new Map()
    let projectFilter = new Map()
    async function mapTasks() {
        for (let i = 0; i < projectNames.length; i++) {
            if (projectTaskMap.has(projectNames[i])) {
                let taskList = projectTaskMap.get(projectNames[i])
                taskList.push((await tasks)[i])
                projectTaskMap.set(projectNames[i], taskList)
            } else {
                projectTaskMap.set(projectNames[i], [(await tasks)[i]])
                projectFilter.set(projectNames[i], true)
            }
        }
    }

    async function filterTasks() {
        let newTasks: Task[] = []
        for (const [name, showing] of projectFilter.entries()) {
            if (showing && projectTaskMap.has(name)) {
                newTasks.push(...projectTaskMap.get(name))
            }
        }
        // reassign to trigger reactivity
        tasks = new Promise(resolve => {
            resolve([...newTasks])
        })
        console.log(tasks)
    }
</script>

<div>
    <h2 class="scroll-m-20 pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0">
        Your Calendar
    </h2>
    <div>
        <DropdownMenu.Root>
            <DropdownMenu.Trigger>
                <Button variant="outline">Filter</Button>
            </DropdownMenu.Trigger>
            <DropdownMenu.Content class="w-56">
                <DropdownMenu.Label>My Projects</DropdownMenu.Label>
                <DropdownMenu.Separator />
                <DropdownMenu.Group>
                    {#each Array.from(projectFilter.entries()) as [name, selected]}
                        <div>
                            <Switch
                                id="project"
                                checked={selected}
                                onCheckedChange={() =>
                                    projectFilter.set(name, !projectFilter.get(name))}
                            />
                            <Label for="project">{name}</Label>
                        </div>
                    {/each}
                </DropdownMenu.Group>
                <DropdownMenu.Separator />
                <DropdownMenu.Item>
                    <Button
                        onclick={filterTasks}
                        class="mt-2 w-full">Apply Filter</Button
                    >
                </DropdownMenu.Item>
            </DropdownMenu.Content>
        </DropdownMenu.Root>
    </div>
    <div class="inputs">
        <label>
            Year:
            <input
                type="number"
                value={year}
                oninput={handleYearChange}
            />
        </label>
        <label>
            Month:
            <input
                type="number"
                min="1"
                max="12"
                value={month}
                oninput={handleMonthChange}
            />
        </label>
    </div>
    {#await tasks}
        <p>Loading calendar...</p>
    {:then tasks}
        <Calendar
            {month}
            {year}
            {tasks}
        />
    {:catch error}
        <p class="text-red-500">Failed to load calendar.</p>
    {/await}
</div>
