<script lang="ts">
    import { connectToProject, ProjectState } from "../../lib/project_state.svelte"
    import { onMount } from "svelte"

    let projectId = $state("")

    let project: Promise<ProjectState> = $state(new Promise((_a, _b) => {}))

    let key = $state("Key")
    let value = $state("Value")

    onMount(() => {
        //Grab project ID from URL
        let params = new URLSearchParams(document.location.search)
        projectId = params.get("id") || ""
        //TODO: Handle error if no project ID

        project = connectToProject("BOOTSTRAPPER", projectId)
    })
</script>

{#await project}
    <p>Loading project</p>
{:then project}
    <h1>Project page for {project.title}!</h1>
    <p>Description: {project.description}</p>
    <p>Project ID: {project.id}</p>
    <input
        value="X"
        oninput={e => project.updateInProject("reactive_testing.bruh", e.currentTarget.value)}
    />

    <p>{project.reactive_testing.bruh}</p>

    <button
        onclick={() => {
            project.appendInProject("reactive_testing.list", project.reactive_testing.bruh)
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

    <button onclick={() => project.updateInProject("button_state", "enabled")}
        >Enable selector</button
    >
    <input
        type="radio"
        name="options"
        id="A"
        value="A"
        disabled={project.button_state != "enabled"}
        oninput={() => project.select("a")}
    />
    <label for="A">A</label>
    <input
        type="radio"
        name="options"
        id="B"
        value="B"
        disabled={project.button_state != "enabled"}
        oninput={() => project.select("b")}
    />
    <label for="B">B</label>
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
