<script lang="ts">
    import { connectToProject, ProjectState } from "$lib/project_state.svelte"
    import { onMount } from "svelte"
    import * as Tabs from "$lib/components/ui/tabs/index.js"
    import PleiadesNav from "$lib/components/PleiadesNav.svelte"
    import SettingsTab from "./settings/SettingsTab.svelte"
    import TasksTab from "./tasks/TasksTab.svelte"
    import AvailabilityTab from "./availability/AvailabilityTab.svelte"
    import CalendarTab from "./calendar/CalendarTab.svelte"
    import * as Dialog from "$lib/components/ui/dialog"
    import { Button } from "$lib/components/ui/button"

    let projectId = $state("")

    let project: Promise<ProjectState> = $state(new Promise((_a, _b) => {}))

    let key = $state("Key")
    let value = $state("Value")

    onMount(() => {
        //Grab project ID from URL
        let params = new URLSearchParams(document.location.search)
        projectId = params.get("id") || ""

        project = connectToProject(projectId)
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

        <Tabs.Root
            value="tasks"
            class="w-[400px]"
        >
            <Tabs.List>
                <Tabs.Trigger value="tasks">Task Board</Tabs.Trigger>
                <Tabs.Trigger value="calendar">Calendar</Tabs.Trigger>
                <Tabs.Trigger value="availability">Availability</Tabs.Trigger>
                <Tabs.Trigger value="settings">Settings</Tabs.Trigger>
                <Tabs.Trigger value="debug">Debugging</Tabs.Trigger>
            </Tabs.List>

            <TasksTab {project} />

            <CalendarTab {project} />

            <AvailabilityTab {project} />

            <SettingsTab {project} />

            <Tabs.Content value="debug">
                <h1>Project page for {project.title}!</h1>
                <p>Description: {project.description}</p>
                <p>Project ID: {project.id}</p>
                <input
                    value="X"
                    oninput={e =>
                        project.updateInProject("reactive_testing.bruh", e.currentTarget.value)}
                />

                <p>{project.reactive_testing.bruh}</p>

                <button
                    onclick={() => {
                        project.appendInProject(
                            "reactive_testing.list",
                            project.reactive_testing.bruh,
                        )
                    }}>Append</button
                >

                <br />

                {#each project.reactive_testing.list as value}
                    <p>{value}</p>
                {/each}

                <button
                    onclick={() => {
                        project.deleteInProject(
                            `reactive_testing.list[${Math.floor(project.reactive_testing.list.length / 2)}]`,
                        )
                    }}>Delete Halfway</button
                >
                <br />
                <br />

                {#each Object.keys(project.reactive_testing.values) as value}
                    <p>{value}: {project.reactive_testing.values[value]}</p>
                {/each}

                <input bind:value={key} />
                <input bind:value /><br />

                <button
                    onclick={() => {
                        if (project.reactive_testing.values[key] == undefined) {
                            project.appendInProject(`reactive_testing.values.${key}`, value)
                        } else {
                            project.updateInProject(`reactive_testing.values.${key}`, value)
                        }
                    }}>Insert/Update</button
                ><br />
                <button
                    onclick={() => {
                        if (project.reactive_testing.values[key] != undefined) {
                            project.deleteInProject(`reactive_testing.values.${key}`)
                        }
                    }}>Delete</button
                >

                <br />
                <br />

                <button onclick={() => project.select("")}>Enable selector</button>
                <input
                    type="radio"
                    name="options"
                    id="A"
                    value="A"
                    checked={project.demoButtonState == "a"}
                    disabled={project.demoButtonState != ""}
                    oninput={() => project.select("a")}
                />
                <label for="A">A</label>
                <input
                    type="radio"
                    name="options"
                    id="B"
                    value="B"
                    checked={project.demoButtonState == "b"}
                    disabled={project.demoButtonState != ""}
                    oninput={() => project.select("b")}
                />
                <label for="B">B</label>
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
