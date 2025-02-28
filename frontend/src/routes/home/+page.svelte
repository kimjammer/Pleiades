<script lang="ts">
    import { onMount } from "svelte"
    import { toggleMode } from "mode-watcher";
    import { Button } from "$lib/components/ui/button"
    import { Label } from "$lib/components/ui/label"
    import { Input } from "$lib/components/ui/input"
    import * as Dialog from "$lib/components/ui/dialog"
    import  { toast } from "svelte-sonner"
    import { Sun, Moon } from "lucide-svelte";
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
        if (!res.ok) {
            toast.error("Failed to load projects")
            return
        }
        response = (await res.json()) as ProjectsResponse
    }

    async function createProject() {
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

<div class="flex justify-between p-5 border-b">
    <div>
        <h1 class="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl">
            Pleiades
        </h1>
    </div>
    <div class="flex gap-5">
        <Button onclick={toggleMode} variant="outline" size="icon">
            <Sun
                class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
            />
            <Moon
                class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
            />
            <span class="sr-only">Toggle theme</span>
        </Button>
        <Button>
            Logout
        </Button>
    </div>
</div>

<div class="p-5">
    <div>
        <h2 class="scroll-m-20 pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0">
            Your Projects
        </h2>
    </div>
    <div class="flex flex-wrap mt-5 mb-5 gap-10">
        {#each response.projects as project}
            <a class="p-5 w-60 h-60 flex flex-col justify-end
            border-8 border-primary
            hover:bg-slate-300 hover:dark:bg-slate-800
            transition duration-300
            rounded-xl overflow-hidden"
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
                <div class="p-5 w-60 h-60 flex flex-col justify-center
                    border-8 border-dashed border-primary
                    hover:hover:bg-slate-300 hover:dark:bg-slate-800
                    transition duration-300
                    rounded-xl">
                    <h4 class="scroll-m-20 text-xl font-semibold tracking-tight">
                        Create a new project
                    </h4>
                </div>
            </Dialog.Trigger>
            <Dialog.Content>
                <Dialog.Header>
                    <Dialog.Title>
                        Create new Project
                    </Dialog.Title>
                    <Dialog.Description>
                        Choose a name and description for your project and click create!
                    </Dialog.Description>
                    <div class="flex flex-col gap-5 py-5">
                        <div>
                            <Label for="title" class="text-right">Title</Label>
                            <Input id="title" placeholder="Project Name" bind:value={title} />
                        </div>
                        <div>
                            <Label for="description" class="text-right">Description</Label>
                            <Input id="description" placeholder="Project Description" bind:value={description} />
                        </div>
                    </div>
                    <Dialog.Footer>
                        <Button onclick={createProject} onkeypress={createProject} type="submit">
                            Create!
                        </Button>
                    </Dialog.Footer>
                </Dialog.Header>
            </Dialog.Content>
        </Dialog.Root>

    </div>
</div>
