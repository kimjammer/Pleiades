<script lang="ts">
    import { Input } from "$lib/components/ui/input" // shadcn-svelte Input
    import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"
    import PleiadesNav from "$lib/components/PleiadesNav.svelte"
    import PersonalCalendar from "./PersonalCalendar.svelte"
    import { Label } from "$lib/components/ui/label"
    import { Button } from "$lib/components/ui/button/index.js";
    import NotificationSettings from "./NotificationSettings.svelte"
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
        await getTasks()
        await mapTasks()
    })
    let tasks: Task[] = $state([])
    let projectNames: string[] = $state([])

    async function getTasks() {
        try {
            const res = await fetch(PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/getUserTasks", {
                method: "GET",
                mode: "cors",
                credentials: "include",
                headers: { "Content-Type": "application/json" },
            })
            const data = await res.json() as {
                success: boolean;
                tasks: Task[];
                projectNames: string[];
            }
            console.log("Fetched data:", data)
            if (data.success) {
                toast.success("User tasks fetched")
                projectNames = [...data.projectNames]


                tasks = data.tasks as Task[]
            } else {
                toast.error("Failed to load user tasks")
            }
        } catch (error) {
            toast.error("Failed to get user tasks")
            console.error(error)
        }
        //console.log("tasks " + tasks)
        //console.log("projectNames " + projectNames)
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
        //console.log(projectNames)
        //console.log(projectFilter)
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

<NotificationSettings></NotificationSettings>

<PersonalCalendar></PersonalCalendar>
