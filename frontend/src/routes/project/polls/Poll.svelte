<script lang="ts">
    import type { ProjectState, Poll, Option } from "$lib/project_state.svelte"
    import { Button } from "$lib/components/ui/button"

    let { project, poll }: { project: ProjectState; poll: Poll } = $props()

    function voteFor(option: Option, level: string) {
        let alreadySelected = null
        if (option.likedUsers.includes(project.userId)) {
            alreadySelected = "LikedUsers"
        } else if (option.neutralUsers.includes(project.userId)) {
            alreadySelected = "NeutralUsers"
        } else if (option.dislikedUsers.includes(project.userId)) {
            alreadySelected = "DislikedUsers"
        }

        if (alreadySelected == level) {
            return
        }

        if (alreadySelected != null) {
            project.deleteInProject(
                `Polls[Id=${poll.id}].Options[Id=${option.id}].${alreadySelected}[$IT=${project.userId}]`,
            )
        }

        project.appendInProject(
            `Polls[Id=${poll.id}].Options[Id=${option.id}].${level}`,
            project.userId,
        )
    }
</script>

<div class="poll mt-[0.5em] border">
    <h1 class="mb-[0.3em] text-xl">{poll.title}</h1>
    {#if poll.description != ""}
        <p class="mb-[0.5em] opacity-80">{poll.description}</p>
    {/if}

    <div class="options-grid">
        {#each poll.options as option}
            <p class="option-name">{option.title}</p>
            <Button
                class="hover:border-transparent sctd:bg-green-400/80 sctd:hover:bg-green-400/80 not-sctd:hover:bg-green-400/40 sctd:dark:bg-green-700/80 sctd:dark:hover:bg-green-700/80 not-sctd:dark:hover:bg-green-700/40"
                data-selected={option.likedUsers.includes(project.userId)}
                onclick={() => voteFor(option, "LikedUsers")}
                variant="outline"
                size="icon">✅</Button
            >
            <Button
                class="hover:border-transparent sctd:bg-yellow-400/80 sctd:hover:bg-yellow-400/80 not-sctd:hover:bg-yellow-400/40 sctd:dark:bg-yellow-700/80 sctd:dark:hover:bg-yellow-700/80 not-sctd:dark:hover:bg-yellow-700/40"
                data-selected={option.neutralUsers.includes(project.userId)}
                onclick={() => voteFor(option, "NeutralUsers")}
                variant="outline"
                size="icon">🟡</Button
            >
            <Button
                class="hover:border-transparent sctd:bg-red-400/80 sctd:hover:bg-red-400/80 not-sctd:hover:bg-red-400/40 sctd:dark:bg-red-700/80 sctd:dark:hover:bg-red-700/80 not-sctd:dark:hover:bg-red-700/40"
                data-selected={option.dislikedUsers.includes(project.userId)}
                onclick={() => voteFor(option, "DislikedUsers")}
                variant="outline"
                size="icon">❌</Button
            >
        {/each}
    </div>
</div>

<style>
    .poll {
        padding: 0.5em;
        border-radius: 0.5em;
        width: fit-content;
    }

    .options-grid {
        display: grid;
        grid-template-columns: 1fr min-content min-content min-content;
        column-gap: 0.2em;
        row-gap: 0.4em;
    }

    .option-name {
        align-self: center;
        margin-right: 0.5em;
        text-align: right;
    }
</style>
