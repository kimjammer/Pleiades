<script>
    import { onMount } from "svelte"
    import { Card } from "$lib/components/ui/card"
    import { Button } from "$lib/components/ui/button"

    onMount(async () => {
        const url = "http://localhost:8080/projects"
        const response = await fetch(url, { mode: "cors" })
        console.log(response)
    })

    let response = [
        {
            title: "project",
            description: "project 428282",
        },
        {
            title: "project2",
            description: "project descirpit",
        },
    ]

    async function createProject() {
        const url = "http://localhost:8080/projects/new"
        await fetch(url, {
            method: "POST",
            mode: "cors",
        })
    }
</script>

<h1 class="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl">Pleiades Home Page</h1>

{#each response as project}
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
