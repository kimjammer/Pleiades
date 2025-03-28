<script lang="ts">
    import { Button } from "$lib/components/ui/button"
    import { Skeleton } from "$lib/components/ui/skeleton"
    import { Play, Square } from "lucide-svelte"
    import type { ProjectState, Task, Session } from "$lib/project_state.svelte"
    import { toast } from "svelte-sonner"
    import { onMount } from "svelte"

    let { project, task }: { project: ProjectState; task: Task } = $props()

    let crrSession: Session | undefined = $derived.by(() => {
        //Infer correct session
        let proxy = task.sessions
        let crrUserSessions = proxy.filter(session => session.user === localStorage.myId)
        return crrUserSessions.find(session => session.startTime != 0 && session.endTime === 0)
    })

    //Valid statuses are Unknown, NotStarted, Running
    let status = $derived.by(() => {
        if (crrSession == undefined) {
            return "NotStarted"
        } else {
            return "Running"
        }
    })

    let duration = $derived.by(() => {
        if (crrSession) {
            let secs = Math.floor((Date.now() - crrSession.startTime) / 1000)
            return `${Math.floor(secs / 60)}:${secs % 60}`
        } else {
            return "00:00"
        }
    })

    const handleStart = () => {
        //Start session
        const newSession: Session = {
            id: crypto.randomUUID(),
            startTime: Date.now(),
            endTime: 0,
            user: localStorage.myId,
        }
        task.sessions.push(newSession)
        console.log(task)
        //project.appendInProject(`Tasks`, newSession)
    }

    const handleStop = () => {
        //Stop session
        if (crrSession) {
            crrSession.endTime = Date.now()
            //project.updateInProject(`Sessions[Id=${crrSession.id}].endTime`, Date.now())
        } else {
            toast.error("Failed to stop session")
        }
    }

    onMount(() => {
        setInterval(() => {
            if (crrSession) {
                let secs = Math.floor((Date.now() - crrSession.startTime) / 1000)
                duration = `${Math.floor(secs / 60)}:${secs % 60}`
            } else {
                duration = "00:00"
            }
        }, 1000)
    })
</script>

<div
    class="inline-flex items-center justify-center gap-1 rounded-md border border-input bg-background p-2"
>
    {#if status === "NotStarted"}
        <small class="text-sm font-medium leading-none"> Start Session </small>
        <Button
            class="rounded-full"
            variant="outline"
            size="icon"
            onclick={handleStart}
        >
            <Play />
        </Button>
    {:else if status === "Running"}
        <small class="w-10 text-sm font-medium leading-none">
            {duration}
        </small>
        <Button
            class="rounded-full"
            variant="outline"
            size="icon"
            onclick={handleStop}
        >
            <Square />
        </Button>
    {:else}
        <Skeleton class="h-4 w-[100px]" />
    {/if}
</div>
Sessions: {task.sessions.length}
