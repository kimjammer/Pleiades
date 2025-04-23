<script lang="ts">
    import { Input } from "$lib/components/ui/input" // shadcn-svelte Input
    import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"
    import PleiadesNav from "$lib/components/PleiadesNav.svelte"
    import { Label } from "$lib/components/ui/label"
    import { Switch } from "$lib/components/ui/switch"
    import * as Card from "$lib/components/ui/card"
    import { toast } from "svelte-sonner"
    import type { Task } from "$lib/project_state.svelte"
    // Optional: for notifications
    import UserAvatar from "$lib/components/UserAvatar.svelte"
    import type { ChangeEventHandler } from "svelte/elements"
    import Calendar from "../project/calendar/Calendar.svelte"
    import {onMount} from "svelte";
    import {goto} from "$app/navigation";
    import {base} from "$app/paths";
    import Teammate from "../project/settings/Teammate.svelte";


    /*TODO: Create boolean array in users
            update user websocket
            bind boolean values to switches
            create function to switch booleans
    */
    let notifUserJoin = $state(false)
    let notifPollEnd = $state(false)
    let notifTaskAssign = $state(false)
    /*
        TODO: Create hovercard for each task
              Create filtering UI for personal calendar
              Bind UI to functions updating tasks[]
     */


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
        getTasks()
        mapTasks()
        await getNotifSettings()
    })
    let tasks: Task[] = $state([])
    let projectNames: string[] = []
    async function getTasks() {
        try {
            const res = await fetch(PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/getUserTasks", {
                method: "GET",
                mode: "cors",
                credentials: "include",
                headers: { "Content-Type": "application/json" },
            })
            const data = await res.json()
            if (data.success) {
                toast.success("User tasks fetched")
                projectNames = data.projectNames
                tasks = data.tasks as Task[]
            } else {
                toast.error(data.error)
            }
        } catch (error) {
            toast.error("Failed to get user tasks")
            console.error(error)
        }
        console.log("tasks " + tasks)
        return []
    }
    let projectTaskMap = new Map()
    let projectFilter = new Map()
    async function mapTasks() {
        for (let i = 0; i < projectNames.length; i++) {
            if (projectTaskMap.has(projectNames[i])) {
                let taskList = projectTaskMap.get(projectNames[i])
                taskList.push(tasks[i])
                projectTaskMap.set(projectNames[i], taskList)
            } else {
                projectTaskMap.set(projectNames[i], [tasks[i]])
                projectFilter.set(projectNames[i], true)
            }
        }
    }

    async function filterTasks() {
        //Create function that goes through this projectFilter and uses projectTaskMap to redefine tasks
        tasks = []
        for (const [name, showing] of projectFilter) {
            if (showing) {
                tasks.push(...projectTaskMap.get(name))
            }
        }
    }


    let selectedFile

    // Handle file selection and upload
    async function handleFileSelect(event: Parameters<ChangeEventHandler<HTMLInputElement>>[0]) {
        const file = event.currentTarget.files![0]
        if (!file) return

        const reader = new FileReader()
        reader.onloadend = async () => {
            let base64Image = reader.result
            if (typeof base64Image === "string") {
                //can only use .split on string
                base64Image = base64Image.split(",")[1] // Extract base64 part for backend
            }
            selectedFile = base64Image //change to int64

            //Send image to the backend
            try {
                const res = await fetch(PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/profilepic", {
                    method: "POST",
                    mode: "cors",
                    credentials: "include",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify({ image: selectedFile }),
                })
                const data = await res.json()
                if (data.success) {
                    toast.success("Profile picture uploaded successfully!")
                } else {
                    toast.error(data.error)
                }
            } catch (error) {
                toast.error("Upload failed. Please try again.")
                console.error(error)
            }
        }
        reader.readAsDataURL(file)
    }

    async function flipNotifSetting(index: number) {
        console.log("Flipping notification setting")
        const res = await fetch(PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/flipNotif", {
            method: "POST",
            mode: "cors",
            credentials: "include",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ notifIndex: index }),
        })
        const data = await res.json()
        if (data.success) {
            toast.success("Changed notification preference")
            getNotifSettings()
        } else {
            toast.error("Error changing notification preference")
        }
    }

    async function getNotifSettings() {
        const res = await fetch(PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/notifSettings", {
            method: "GET",
            mode: "cors",
            credentials: "include",
            headers: { "Content-Type": "application/json" },
        })
        const data = await res.json()
        if (data.success) {
            [notifUserJoin, notifPollEnd, notifTaskAssign] = data.notifSettings
            console.log(notifUserJoin, notifPollEnd, notifTaskAssign)
        } else {
            toast.error("Failed to load notification settings")
        }
    }


</script>

<PleiadesNav></PleiadesNav>
<div>
    <h2 class="scroll-m-20 pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0">
        Hello
    </h2>
</div>

<div class="grid w-full max-w-sm items-center gap-1.5">
    <UserAvatar
        project={null}
        userID={localStorage.myId}
    />
    <Label>Upload Profile Picture</Label>
    <!-- Trigger file input -->
    <Input
        id="fileInput"
        type="file"
        accept="image/png, image/jpeg"
        onchange={handleFileSelect}
    />
</div>

<Card.Root class="mt-2 w-fit">
    <Card.Header>
        <Card.Title>Notification Settings</Card.Title>
    </Card.Header>
    <Card.Content class="grid gap-6">
        <div class="flex items-center space-x-2">
            <Switch
                    className="user joining"
                    bind:checked={notifUserJoin}
                    onCheckedChange={() => flipNotifSetting(0)}
            />
            <Label className="user joining">New users joining projects</Label>
        </div>
        <div class="flex items-center space-x-2">
            <Switch
                    className="ending polls"
                    bind:checked={notifPollEnd}
                    onCheckedChange={() => flipNotifSetting(1)}
            />
            <Label className="ending polls">Polls ending soon</Label>
        </div>
        <div class="flex items-center space-x-2">
            <Switch
                    className="task assignments"
                    bind:checked={notifTaskAssign}
                    onCheckedChange={() => flipNotifSetting(2)}
            />
            <Label className="task assignments">New task assignments</Label>
        </div>
    </Card.Content>
</Card.Root>

<div class="p-5">
    <div>
        <h2
            class="scroll-m-20 pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0"
        >
            Your Calendar
        </h2>
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
            <p>Loading calendar...</p> <!-- You can also use a Skeleton or spinner here -->
        {:then tasks}
            <Calendar {month} {year} {tasks} />
        {:catch error}
            <p class="text-red-500">Failed to load calendar.</p>
        {/await}
    </div>
</div>
