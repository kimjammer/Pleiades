<script lang="ts">
    import { onMount } from "svelte"
    import { toggleMode } from "mode-watcher";
    import { Button } from "$lib/components/ui/button"
    import { Sun, Moon } from "lucide-svelte";
    import type { ProjectsResponse } from "$lib/schema.js"
    import { PUBLIC_API_HOST } from "$env/static/public"


    let response: ProjectsResponse = $state({
        projects: [],
    })

    onMount(async () => {
        await loadProjects()
    })

    async function loadProjects() {
        const url = "http://" + PUBLIC_API_HOST + "/projects"
        const res = await fetch(url, { mode: "cors" })
        response = (await res.json()) as ProjectsResponse
    }

    async function createProject() {
        const url = "http://" + PUBLIC_API_HOST + "/projects/new"
        await fetch(url, {
            method: "POST",
            mode: "cors",
        })
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
        <div class="p-5 w-60 h-60 flex flex-col justify-center
        border-8 border-dashed border-primary
        hover:hover:bg-slate-300 hover:dark:bg-slate-800
        transition duration-300
        rounded-xl"
             onclick={createProject}
             onkeyup={createProject}
             role="button"
             tabindex="0"
        >
            <h4 class="scroll-m-20 text-xl font-semibold tracking-tight">
                Create a new project
            </h4>
        </div>
    </div>
</div>
