<script lang="ts">
    import * as Tabs from "$lib/components/ui/tabs"
    import type { ProjectState } from "$lib/project_state.svelte.ts"
    import type { PageData } from "./$types"
    import CreationModal from "./CreationModal.svelte"
    import { onMount } from "svelte"
    import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"
    import { goto } from "$app/navigation"
    import { base } from "$app/paths"
    import { toast } from "svelte-sonner"
    import type { PollsResponse } from "$lib/schema.js"
    import { Button } from "$lib/components/ui/button"

    let { project, data }: { project: ProjectState; data: PageData } = $props()
    console.log("data in PollsTab: " + data)

    let showForm = false
    let polls = project.polls //string of poll titles

    function handlePollClick(pollTitle: string) {
        console.log("Poll clicked:", pollTitle)
    }
</script>

<Tabs.Content value="polls">
    <CreationModal
        {project}
        {data}
    />

    {#if polls.length > 0}
        {#each polls as poll}
            <button onclick={() => handlePollClick(poll.id)}>
                {poll.title}
            </button>
        {/each}
    {:else}
        <p>No polls available.</p>
    {/if}

</Tabs.Content>

