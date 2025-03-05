<script lang="ts">
    import { onMount } from "svelte"
    import { Button } from "$lib/components/ui/button"
    import { Label } from "$lib/components/ui/label"
    import { Input } from "$lib/components/ui/input"
    import PleiadesNav from "$lib/components/PleiadesNav.svelte"
    import * as Dialog from "$lib/components/ui/dialog"
    import { toast } from "svelte-sonner"
    import type { ProjectsResponse, newProjectRequest } from "$lib/schema.js"
    import { PUBLIC_API_HOST } from "$env/static/public"

    let title = $state("")
    let description = $state("")
    let response: ProjectsResponse = $state({
        projects: [],
    })
    let createDialogOpen = $state(false)

    onMount(async () => {
        await loadProjects()
    })

    async function loadProjects() {
        const url = "http://" + PUBLIC_API_HOST + "/projects"
        const res = await fetch(url, { mode: "cors", credentials: "include" })
        if (res.status === 401) {
            location.assign("/")
        } else if (!res.ok) {
            toast.error("Failed to load projects")
            return
        }
        response = (await res.json()) as ProjectsResponse
    }

    async function createProject() {
        //Basic validation
        if (title === "") {
            toast.error("Title is required")
            return
        }

        createDialogOpen = false

        const url = "http://" + PUBLIC_API_HOST + "/projects/new"
        const body: newProjectRequest = { title, description }
        const res = await fetch(url, {
            method: "POST",
            mode: "cors",
            credentials: "include",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(body),
        })

        //Error handling
        if (!res.ok) {
            toast.error("Failed to create project")
            return
        }

        await loadProjects()
    }
</script>

<PleiadesNav></PleiadesNav>
<div class="p-5">
    <div>
        <h2
            class="scroll-m-20 pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0"
        >
            Your Projects
        </h2>
    </div>
    <div class="mb-5 mt-5 flex flex-wrap gap-10">
        {#each response.projects as project}
            <a
                class="border-primary flex h-60 w-60 flex-col justify-end
            overflow-hidden rounded-xl
            border-8 p-5
            transition duration-300
            hover:bg-slate-300 hover:dark:bg-slate-800"
                href="/project?id={project.id}"
            >
                <h3 class="scroll-m-20 text-2xl font-semibold tracking-tight">
                    {project.title}
                </h3>
                <p class="leading-7 [&:not(:first-child)]:mt-6">
                    {project.description}
                </p>
            </a>
        {/each}
        <Dialog.Root bind:open={createDialogOpen}>
            <Dialog.Trigger>
                <div
                    class="border-primary flex h-60 w-60 flex-col justify-center
                    rounded-xl border-8 border-dashed
                    p-5 transition
                    duration-300 hover:hover:bg-slate-300
                    hover:dark:bg-slate-800"
                >
                    <h4 class="scroll-m-20 text-xl font-semibold tracking-tight">
                        Create a new project
                    </h4>
                </div>
            </Dialog.Trigger>
            <Dialog.Content>
                <Dialog.Header>
                    <Dialog.Title>Create new Project</Dialog.Title>
                    <Dialog.Description>
                        Choose a name and description for your project and click create!
                    </Dialog.Description>
                    <div class="flex flex-col gap-5 py-5">
                        <div>
                            <Label
                                for="title"
                                class="text-right">Title</Label
                            >
                            <Input
                                id="title"
                                placeholder="Project Name"
                                bind:value={title}
                            />
                        </div>
                        <div>
                            <Label
                                for="description"
                                class="text-right">Description</Label
                            >
                            <Input
                                id="description"
                                placeholder="Project Description"
                                bind:value={description}
                            />
                        </div>
                    </div>
                    <Dialog.Footer>
                        <Button
                            onclick={createProject}
                            onkeypress={createProject}
                            type="submit"
                        >
                            Create!
                        </Button>
                    </Dialog.Footer>
                </Dialog.Header>
            </Dialog.Content>
        </Dialog.Root>
    </div>
</div>
