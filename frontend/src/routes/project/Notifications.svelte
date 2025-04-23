<script lang="ts">
    import type { ProjectState } from "$lib/project_state.svelte"
    import { flip } from "svelte/animate"
    import { backInOut, cubicInOut, expoInOut, quadInOut, sineInOut } from "svelte/easing"
    import { fade } from "svelte/transition"

    let { project }: { project: ProjectState } = $props()
</script>

<div class="container">
    {#each project.notifications as notif, i (notif.id)}
        <!-- svelte-ignore a11y_no_static_element_interactions -->
        <!-- svelte-ignore a11y_click_events_have_key_events -->
        <div
            class="notif mt-[0.5em] border bg-background/80"
            transition:fade={{ duration: 200 }}
            animate:flip={{ duration: 400 }}
            onclick={() => {
                project.notifications.splice(i, 1)
            }}
        >
            <h1 class="mb-[0.3em] text-xl">{notif.title}</h1>
            <p class="mb-[0.5em] opacity-80">{notif.message}</p>
        </div>
    {/each}
</div>

<style>
    .notif {
        padding: 0.5em;
        border-radius: 0.5em;
    }

    .container {
        position: absolute;
        right: 0px;
        width: min-content;
    }
</style>
