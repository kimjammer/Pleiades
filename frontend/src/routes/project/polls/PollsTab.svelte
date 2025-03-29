<script lang="ts">
    import * as Tabs from "$lib/components/ui/tabs"
    import type { ProjectState } from "$lib/project_state.svelte.ts"
    import type { PageData } from "./$types"
    import CreationModal from "./CreationModal.svelte"
    import { onMount } from "svelte"
    import {PUBLIC_API_HOST, PUBLIC_PROTOCOL} from "$env/static/public";
    import {goto} from "$app/navigation";
    import {base} from "$app/paths";
    import {toast} from "svelte-sonner";
    import type { PollsResponse } from "$lib/schema.js"


    let { project, data }: { project: ProjectState; data: PageData } = $props()

    let showForm = false;
    let polls = [] //string of poll titles
    let response: Promise<PollsResponse> = $state(new Promise(() => {}))

    //TODO: finish this so polls automatically populate
    onMount(() => {
        response = loadPolls()
    })

    //TODO: create endpoint in backend
    async function loadPolls() { //TODO: NOT DONE
        return new Promise<PollsResponse>(async (resolve, reject) => {
            const url = PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/polls?id=" + project.id
            const res = await fetch(url, { mode: "cors", credentials: "include" })
            if (res.status === 401) {
                goto(base + "/login")
            } else if (!res.ok) {
                toast.error("Failed to load polls")
                reject()
            }
            resolve((await res.json()) as PollsResponse)
            const data = await res.json();
            polls = data.polls;
        })
    }

    function handlePollClick(pollTitle: string) {
        console.log("Poll clicked:", pollTitle);
    }

    function toggleForm() {
        showForm = !showForm;
    }

</script>

<Tabs.Content value="polls">
    <CreationModal
            {project}
            {data}
    />
</Tabs.Content>

{#if polls.length > 0}
    {#each polls as poll}
        <button on:click={() => handlePollClick(poll.id)}>
            {poll.title}
        </button>
    {/each}
{:else}
    <p>No polls available.</p>
{/if}


