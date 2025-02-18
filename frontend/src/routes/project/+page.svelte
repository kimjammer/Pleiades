<script>
    import { ProjectState } from "../../lib/project_state.svelte"

    let state = $state(new ProjectState())

    let key = "Key"
    let value = "Value"
</script>

<h1>Project page!</h1>

<input
    value="X"
    oninput={e => state.updateInProject("reactive_testing.bruh", e.currentTarget.value)}
/>

<p>{state.reactive_testing.bruh}</p>

<button
    onclick={() => {
        state.appendInProject("reactive_testing.list", state.reactive_testing.bruh)
    }}>Append</button
>

<br />

{#each state.reactive_testing.list as value}
    <p>{value}</p>
{/each}

<button
    onclick={() => {
        state.deleteInProject(
            `reactive_testing.list[${Math.floor(state.reactive_testing.list.length / 2)}]`,
        )
    }}>Delete Halfway</button
>
<br />
<br />

{#each Object.keys(state.reactive_testing.values) as value}
    <p>{value}: {state.reactive_testing.values[value]}</p>
{/each}

<input bind:value={key} />
<input bind:value /><br />

<button
    onclick={() => {
        if (state.reactive_testing.values[key] == undefined) {
            state.appendInProject(`reactive_testing.values.${key}`, value)
        } else {
            state.updateInProject(`reactive_testing.values.${key}`, value)
        }
    }}>Insert/Update</button
><br />
<button
    onclick={() => {
        if (state.reactive_testing.values[key] != undefined) {
            state.deleteInProject(`reactive_testing.values.${key}`)
        }
    }}>Delete</button
>

<br />
<br />

<button onclick={() => state.updateInProject("button_state", "enabled")}>Enable selector</button>
<input
    type="radio"
    name="options"
    id="A"
    value="A"
    disabled={state.button_state != "enabled"}
    oninput={() => state.select("a")}
/>
<label for="A">A</label>
<input
    type="radio"
    name="options"
    id="B"
    value="B"
    disabled={state.button_state != "enabled"}
    oninput={() => state.select("b")}
/>
<label for="B">B</label>

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
