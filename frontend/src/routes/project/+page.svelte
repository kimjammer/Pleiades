<script lang="ts">
    import PleiadesNav from "$lib/components/PleiadesNav.svelte"
    import { Button } from "$lib/components/ui/button"
    import * as Dialog from "$lib/components/ui/dialog"
    import * as Tabs from "$lib/components/ui/tabs/index.js"
    import { connectToProject, ProjectState } from "$lib/project_state.svelte.js"
    import { onMount } from "svelte"
    import type { PageData } from "./$types.js"
    import AvailabilityTab from "./availability/AvailabilityTab.svelte"
    import CalendarTab from "./calendar/CalendarTab.svelte"
    import PollsTab from "./polls/PollsTab.svelte"
    import SettingsTab from "./settings/SettingsTab.svelte"
    import StatsTab from "./stats/StatsTab.svelte"
    import TasksTab from "./tasks/TasksTab.svelte"
    import Notifications from "./Notifications.svelte"

    let { data }: { data: PageData } = $props()
    let project: Promise<ProjectState> = $state(new Promise((_a, _b) => {}))

    let projectId = $state("")
    let tab = $state("tasks")

    let now = $state(Date.now())

    onMount(() => {
        //Grab project ID from URL
        let params = new URLSearchParams(document.location.search)
        projectId = params.get("id") || ""

        project = connectToProject(projectId)

        tab = params.get("tab") || "tasks"

        const id = setInterval(() => {
            now = Date.now()
        }, 500)

        return () => clearInterval(id)
    })

    $effect(() => {
        let url = new URL(document.location.href)
        url.searchParams.set("tab", tab)
        window.history.pushState(null, "", url.toString())
    })

    type word = {
        content: string
        isLink: boolean
    }
    let words = $derived.by(async () => {
        return (await project).description.split(" ").map(word => {
            const section: word = {
                content: word + " ",
                isLink: word.startsWith("http"),
            }
            return section
        })
    })
</script>

<PleiadesNav></PleiadesNav>
{#await project}
    <p>Loading project</p>
{:then project}
    <Notifications {project} />
    <div class="p-5">
        <h2
            class="scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0"
        >
            {project.title}
        </h2>
        <p class="my-5 leading-7 [&:not(:first-child)]:mt-6">
            {#await words then words}
                {#each words as word}
                    {#if word.isLink}
                        <a
                            href={word.content}
                            target="_blank"
                            class="text-blue-500"
                        >
                            {word.content}
                        </a>
                    {:else}
                        {word.content}
                    {/if}
                {/each}
            {/await}
        </p>

        <Dialog.Root bind:open={project.showError}>
            <Dialog.Content>
                <Dialog.Header><p class="text-xl">Error</p></Dialog.Header>
                <div>
                    <p class="pb-8">{project.error}</p>
                    <Button
                        onclick={() => {
                            project.showError = false
                        }}>Close</Button
                    >
                </div>
            </Dialog.Content>
        </Dialog.Root>

        <Tabs.Root bind:value={tab}>
            <Tabs.List>
                <Tabs.Trigger value="tasks">Task Board</Tabs.Trigger>
                <Tabs.Trigger value="calendar">Calendar</Tabs.Trigger>
                <Tabs.Trigger value="availability">Availability</Tabs.Trigger>
                <Tabs.Trigger value="stats">Stats</Tabs.Trigger>
                <Tabs.Trigger value="polls">Polls</Tabs.Trigger>
                <Tabs.Trigger value="settings">Settings</Tabs.Trigger>
                <Tabs.Trigger value="debug">Debugging</Tabs.Trigger>
            </Tabs.List>

            <TasksTab
                {project}
                {data}
            />

            <CalendarTab {project} />

            <AvailabilityTab {project} />

            <PollsTab
                {project}
                {data}
                {now}
            />

            <SettingsTab {project} />

            <StatsTab {project} />

            <Tabs.Content value="debug">
                <h1>Project page for {project.title}!</h1>
                <p>Description: {project.description}</p>
                <p>Project ID: {project.id}</p>
                <input
                    value="X"
                    oninput={e => project.updateInProject("Title", e.currentTarget.value)}
                />
                <button onclick={() => project.select("")}>Enable selector</button>
                <input
                    type="radio"
                    name="options"
                    id="A"
                    value="A"
                    checked={project.demoButtonState === "a"}
                    disabled={project.demoButtonState !== ""}
                    oninput={() => project.select("a")}
                />
                <label for="A">A</label>
                <input
                    type="radio"
                    name="options"
                    id="B"
                    value="B"
                    checked={project.demoButtonState === "b"}
                    disabled={project.demoButtonState !== ""}
                    oninput={() => project.select("b")}
                />
                <label for="B">B</label>
                <button
                    onclick={() => {
                        project.notify(null, "Debug", "Hello", "world!")
                    }}>Notify!</button
                >
            </Tabs.Content>
        </Tabs.Root>
    </div>
{:catch err}
    <p>{err}</p>
{/await}

<style>
    input,
    button {
        border: 1px solid black;
        margin: 4px;
        padding-left: 4px;
        padding-right: 4px;
    }

    button {
        border-radius: 4px;
    }
</style>
