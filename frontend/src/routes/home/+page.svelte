<script lang="ts">
    import { onMount } from "svelte"
    import { Card } from "$lib/components/ui/card"
    import { Button } from "$lib/components/ui/button"
    import type { ProjectsResponse } from "$lib/schema.js"
    import { PUBLIC_API_HOST } from "$env/static/public"

    let response: ProjectsResponse = {
        projects: [],
    }

    onMount(async () => {
        const url = "http://" + PUBLIC_API_HOST + "/projects"
        const res = await fetch(url, { mode: "cors" })
        response = (await res.json()) as ProjectsResponse
    })

    async function createProject() {
        const url = "http://" + PUBLIC_API_HOST + "/projects/new"
        await fetch(url, {
            method: "POST",
            mode: "cors",
        })
    }
</script>

<h1 class="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl">Pleiades Home Page</h1>

{#each response.projects as project}
    <div class="p-10">
        <h2
            class="scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0"
        >
            {project.title}
        </h2>
        <h2
            class="scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0"
        >
            {project.description}
        </h2>
    </div>
{/each}

<Button onclick={createProject}>Create Project</Button>

<Card></Card>
