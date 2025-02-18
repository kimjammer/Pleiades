<script>
    import { ProjectState } from "../../lib/project_state.svelte"

    let state = $state(new ProjectState())
</script>

<h1>Project page!</h1>

<button onclick={() => state.enable()}>Enable selector</button>
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
