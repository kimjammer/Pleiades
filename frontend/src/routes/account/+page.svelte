<script lang="ts">
    import { Input } from "$lib/components/ui/input" // shadcn-svelte Input
    import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"
    import PleiadesNav from "$lib/components/PleiadesNav.svelte"
    import { Label } from "$lib/components/ui/label"
    import { toast } from "svelte-sonner"
    import type { Task } from "$lib/project_state.svelte"
    // Optional: for notifications
    import UserAvatar from "$lib/components/UserAvatar.svelte"
    import type { ChangeEventHandler } from "svelte/elements"
    import Calendar from "../project/calendar/Calendar.svelte"
    import {onMount} from "svelte";
    import {goto} from "$app/navigation";
    import {base} from "$app/paths";

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
    onMount(() => {
        getTasks()
    })
    let tasks: Task[]
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
                tasks = data.tasks as Task[]
            } else {
                toast.error(data.error)
            }
        } catch (error) {
            toast.error("Failed to get user tasks")
            console.error(error)
        }
        console.log("tasks " + tasks)
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
        <Calendar
                {month}
                {year}
                {tasks}
        />
    </div>
</div>
