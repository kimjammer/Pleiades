<script lang="ts">
    import type { ProjectState, Poll } from "$lib/project_state.svelte"
    import { Button } from "$lib/components/ui/button"

    let { project, poll }: { project: ProjectState; poll: Poll } = $props()
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
                class="hover:border-transparent hover:bg-green-400/40 dark:hover:bg-green-700/40"
                variant="outline"
                size="icon">✅</Button
            >
            <Button
                class="hover:border-transparent hover:bg-yellow-400/40 dark:hover:bg-yellow-700/40"
                variant="outline"
                size="icon">🟡</Button
            >
            <Button
                class="hover:border-transparent hover:bg-red-400/40 dark:hover:bg-red-700/40"
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
