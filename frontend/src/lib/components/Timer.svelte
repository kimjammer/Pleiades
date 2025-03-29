<script lang="ts">
    import { Button } from "$lib/components/ui/button"
    import { Skeleton } from "$lib/components/ui/skeleton"
    import * as HoverCard from "$lib/components/ui/hover-card"
    import { Play, Square } from "lucide-svelte"
    import type { ProjectState, Task, Session } from "$lib/project_state.svelte"
    import { toast } from "svelte-sonner"
    import { onMount } from "svelte"
    import UserAvatar from "$lib/components/UserAvatar.svelte"

    let { project, task }: { project: ProjectState; task: Task } = $props()

    let crrSession: Session | undefined = $derived.by(() => {
        //Infer correct session
        let proxy = task.sessions
        let crrUserSessions = proxy.filter(session => session.user === localStorage.myId)
        return crrUserSessions.find(session => session.startTime != 0 && session.endTime === 0)
    })

    //Valid statuses are NotStarted, Running
    let status = $derived.by(() => {
        if (crrSession == undefined) {
            return "NotStarted"
        } else {
            return "Running"
        }
    })

    let duration = $state("")

    type timeTotal = {
        id: string
        time: string
    }

    let timeTotals: timeTotal[] = $derived.by(() => {
        let completedSessions = task.sessions.filter(session => session.endTime != 0)
        let userSessions = Object.groupBy(completedSessions, ({ user }) => user)
        let result = []
        for (const [userID, sessions] of Object.entries(userSessions)) {
            let sum = 0
            if (sessions == undefined) continue
            for (const session of sessions) {
                if (session.endTime != 0) sum += session.endTime - session.startTime
            }
            let secs = Math.floor(sum / 1000)
            result.push({
                id: userID,
                time: `${Math.floor(secs / 60)
                    .toString()
                    .padStart(2, "0")}:${(secs % 60).toString().padStart(2, "0")}`,
            })
        }
        return result
    })

    let listDisabled = $derived(task.sessions.filter(session => session.endTime != 0).length <= 0)

    const handleStart = () => {
        //Start session
        const newSession: Session = {
            id: crypto.randomUUID(),
            startTime: Date.now(),
            endTime: 0,
            user: localStorage.myId,
        }
        task.sessions.push(newSession)
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
        const interval = setInterval(() => {
            if (crrSession) {
                let secs = Math.floor((Date.now() - crrSession.startTime) / 1000)
                duration = `${Math.floor(secs / 60)
                    .toString()
                    .padStart(2, "0")}:${(secs % 60).toString().padStart(2, "0")}`
            } else {
                duration = "00:00"
            }
        }, 1000)

        return () => clearInterval(interval)
    })
</script>

{#if listDisabled}
    <!--TODO: When inside task card, probably don't want all this padding-->
    <div
        class="inline-flex items-center justify-center gap-1 rounded-md border border-input bg-background p-2"
    >
        {#if status === "NotStarted"}
            <small class="text-sm font-medium leading-none">Sessions</small>
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
{:else}
    <HoverCard.Root>
        <HoverCard.Trigger>
            <div
                class="inline-flex items-center justify-center gap-1 rounded-md border border-input bg-background p-2"
            >
                {#if status === "NotStarted"}
                    <small class="text-sm font-medium leading-none">Sessions</small>
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
        </HoverCard.Trigger>
        <HoverCard.Content side="bottom">
            {#each timeTotals as total}
                <div class="inline-block">
                    <UserAvatar
                        {project}
                        userID={total.id}
                    />
                </div>
                {total.time}
            {/each}
        </HoverCard.Content>
    </HoverCard.Root>
{/if}
