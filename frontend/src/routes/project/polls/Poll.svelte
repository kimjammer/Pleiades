<script lang="ts">
    import {
        type ProjectState,
        type Poll,
        type Option,
        Notification,
    } from "$lib/project_state.svelte"
    import PollVoting from "./PollVoting.svelte"
    import PollResults from "./PollResults.svelte"
    import * as ContextMenu from "$lib/components/ui/context-menu"

    let { project, poll, now }: { project: ProjectState; poll: Poll; now: number } = $props()

    let date = $derived(new Date(poll.dueDate).toLocaleDateString())

    let done = $derived(now > poll.dueDate)

    let prevDone = done

    // The poll component exists globally so this runs no matter the tab
    $effect(() => {
        if (done && !prevDone) {
            prevDone = true
            // Sending to just the current user is OK because this will happen for all users
            project.notify(
                project.userId,
                "poll",
                `Poll ${poll.title} finished!`,
                "Go check it out!",
            )
        }
    })

    function countVotes(option: Option): number {
        return option.likedUsers.length + option.neutralUsers.length + option.dislikedUsers.length
    }

    function optionScore(option: Option): number {
        let count = countVotes(option)

        if (count == 0) {
            return -Infinity
        }

        return (option.likedUsers.length - option.dislikedUsers.length) / count
    }

    let orderedOptions = $derived.by(() => {
        if (!done) {
            return poll.options
        }

        return poll.options.toSorted((a, b) => {
            let aScore = optionScore(a)
            let bScore = optionScore(b)

            if (aScore != bScore) {
                return bScore - aScore
            }

            let aVotes = countVotes(a)
            let bVotes = countVotes(b)

            return bVotes - aVotes
        })
    })
</script>

<ContextMenu.Root>
    <ContextMenu.Trigger class="w-fit">
        <div class="poll mt-[0.5em] border">
            <h1 class="mb-[0.3em] text-xl">{poll.title}</h1>
            {#if poll.description != ""}
                <p class="mb-[0.5em] opacity-80">{poll.description}</p>
            {/if}
            <p class="mb-[0.5em] opacity-80">Over at {date}</p>

            <div class="options-grid">
                {#each orderedOptions as option}
                    <p class="option-name">{option.title}</p>

                    {#if done}
                        <PollResults
                            {project}
                            {poll}
                            {option}
                        />
                    {:else}
                        <PollVoting
                            {project}
                            {poll}
                            {option}
                        />
                    {/if}
                {/each}
            </div>
        </div>
    </ContextMenu.Trigger>
    <ContextMenu.Content>
        {#if !done}
            <ContextMenu.Item
                onclick={() => {
                    let now = Date.now()
                    project.updateInProject(`Polls[Id=${poll.id}].DueDate`, now)
                }}>Finish Now</ContextMenu.Item
            >
        {/if}
        <ContextMenu.Item
            onclick={() => {
                project.deleteInProject(`Polls[Id=${poll.id}]`)
            }}>Delete</ContextMenu.Item
        >
    </ContextMenu.Content>
</ContextMenu.Root>

<style>
    .poll {
        padding: 0.5em;
        border-radius: 0.5em;
        width: fit-content;
    }

    .options-grid {
        display: grid;
        grid-template-columns: 1fr min-content;
        column-gap: 0.2em;
        row-gap: 0.4em;
    }

    .option-name {
        align-self: center;
        margin-right: 0.5em;
        text-align: right;
    }
</style>
