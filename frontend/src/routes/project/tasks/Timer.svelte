<script lang="ts">
    import { Button } from "$lib/components/ui/button"
    import * as HoverCard from "$lib/components/ui/hover-card"
    import UserAvatar from "$lib/components/UserAvatar.svelte"
    import type { ProjectState, Session, Task } from "$lib/project_state.svelte.js"
    import { recordEvent } from "$lib/restApi"
    import { Play, Square } from "lucide-svelte"
    import { onMount } from "svelte"
    import { toast } from "svelte-sonner"

    let {
        project,
        task,
        progress = $bindable(),
    }: { project: ProjectState; task: Task; progress: number | undefined } = $props()

    //Infer correct session
    let crrSession: Session | undefined = $derived.by(() => {
        //Find session by this user without an endTime
        let crrUserSessions = task.sessions.filter(session => session.user === localStorage.myId)
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

    //Display string of current session duration
    let crrDuration = $state("")

    //Calculate total duration of all sessions by all users
    let totalDuration = $derived.by(() => {
        //No progress if there is no time estimate
        if (task.timeEstimate === 0) return 0

        let sum = 0
        for (const session of task.sessions) {
            if (session.endTime != 0) {
                sum += session.endTime - session.startTime
            }
        }
        return sum
    })

    //Time estimate, used as the max value in progress bar
    let totalEstimate = $derived.by(() => {
        //Don't divide by 0
        if (task.timeEstimate === 0) return 1
        return task.timeEstimate
    })

    type timeTotal = {
        id: string
        time: string
    }

    //Calculate total time for each user
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

    //Do not show list if there are no completed sessions
    let listDisabled = $derived(task.sessions.filter(session => session.endTime != 0).length <= 0)

    //Start session
    const handleStart = () => {
        project.appendInProject(`Tasks[Id=${task.id}].Sessions`, {
            id: crypto.randomUUID(),
            startTime: Date.now(),
            endTime: 0,
            user: localStorage.myId,
        })
    }

    //Stop session
    const handleStop = () => {
        if (crrSession) {
            project.updateInProject(
                `Tasks[Id=${task.id}].Sessions[Id=${crrSession.id}].EndTime`,
                Date.now(),
            )
            let secs = Math.floor((Date.now() - crrSession.startTime) / 1000)
            recordEvent("seconds", secs)
        } else {
            toast.error("Failed to stop session")
        }
    }

    onMount(() => {
        const interval = setInterval(() => {
            if (crrSession) {
                let secs = Math.floor((Date.now() - crrSession.startTime) / 1000)
                crrDuration = `${Math.floor(secs / 60)
                    .toString()
                    .padStart(2, "0")}:${(secs % 60).toString().padStart(2, "0")}`
            } else {
                crrDuration = "00:00"
            }
        }, 1000)

        return () => clearInterval(interval)
    })

    $effect(() => {
        progress = (100 * (totalDuration ?? 0)) / (totalEstimate ?? 1)
    })
</script>

{#snippet core()}
    <div class="inline-block">
        <div class="flex items-center justify-center gap-1 p-2">
            {#if status === "NotStarted"}
                <Button
                    class="rounded-full"
                    variant="outline"
                    size="icon"
                    onclick={handleStart}
                >
                    <Play />
                </Button>
                <small class="text-sm font-medium leading-none">Sessions</small>
            {:else if status === "Running"}
                <Button
                    class="rounded-full"
                    variant="outline"
                    size="icon"
                    onclick={handleStop}
                >
                    <Square />
                </Button>
                <small class="w-10 text-sm font-medium leading-none">
                    {crrDuration}
                </small>
            {/if}
        </div>
    </div>
{/snippet}

{#if listDisabled}
    {@render core()}
{:else}
    <HoverCard.Root>
        <HoverCard.Trigger>
            {@render core()}
        </HoverCard.Trigger>
        <HoverCard.Content
            side="bottom"
            align="start"
            sideOffset={0}
        >
            {#each timeTotals as total}
                <div class="flex items-center gap-1">
                    <div class="inline-block">
                        <UserAvatar
                            {project}
                            userID={total.id}
                        />
                    </div>
                    <p>
                        {total.time}
                    </p>
                </div>
            {/each}
        </HoverCard.Content>
    </HoverCard.Root>
{/if}
