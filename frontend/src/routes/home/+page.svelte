<script lang="ts">
    import { goto } from "$app/navigation"
    import { base } from "$app/paths"
    import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"
    import PleiadesNav from "$lib/components/PleiadesNav.svelte"
    import { Button } from "$lib/components/ui/button"
    import * as Dialog from "$lib/components/ui/dialog"
    import { Input } from "$lib/components/ui/input"
    import { Label } from "$lib/components/ui/label"
    import { Skeleton } from "$lib/components/ui/skeleton"
    import { recordEvent } from "$lib/restApi"
    import type { ProjectsResponse, newProjectRequest } from "$lib/schema.js"
    import { onMount } from "svelte"
    import { toast } from "svelte-sonner"

    let title = $state("")
    let description = $state("")
    let response: Promise<ProjectsResponse> = $state(new Promise(() => {}))
    let createDialogOpen = $state(false)

    onMount(() => {
        response = loadProjects()
    })

    async function loadProjects() {
        return new Promise<ProjectsResponse>(async (resolve, reject) => {
            const url = PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/projects"
            const res = await fetch(url, { mode: "cors", credentials: "include" })
            if (res.status === 401) {
                goto(base + "/login")
            } else if (!res.ok) {
                toast.error("Failed to load projects")
                reject()
            }
            resolve((await res.json()) as ProjectsResponse)
        })
    }

    async function createProject() {
        //Basic validation
        if (title === "") {
            toast.error("Title is required")
            return
        }

        createDialogOpen = false

        const url = PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/projects/new"
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

        recordEvent("projects")
        response = Promise.resolve(await loadProjects())
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
    {#await response}
        <div class="mb-5 mt-5 flex flex-wrap gap-10">
            <Skeleton class="h-60 w-60" />
        </div>
    {:then response}
        <div class="mb-5 mt-5 flex flex-wrap gap-10">
            {#each response.projects as project}
                <a
                    class="border-primary flex h-60 w-60 flex-col justify-end
            overflow-hidden rounded-xl
            border-8 p-5
            transition duration-300
            hover:bg-slate-300 hover:dark:bg-slate-800"
                    href="{base}/project?id={project.id}"
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
    {/await}
</div>
