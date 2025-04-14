<script lang="ts">
    import * as Tabs from "$lib/components/ui/tabs"
    import type { ProjectState } from "$lib/project_state.svelte"
    import type { PageData } from "./$types"
    import CreationModal from "./CreationModal.svelte"
    import PollComponent from "./Poll.svelte"

    let { project, data }: { project: ProjectState; data: PageData } = $props()
    console.log("data in PollsTab: " + data)

    let polls = project.polls //string of poll titles
</script>

<Tabs.Content value="polls">
    <CreationModal
        {project}
        {data}
    />

    <div>
        {#if polls.length > 0}
            {#each polls as poll}
                <PollComponent
                    {project}
                    {poll}
                />
            {/each}
        {:else}
            <p>No polls available.</p>
        {/if}
    </div>
</Tabs.Content>
