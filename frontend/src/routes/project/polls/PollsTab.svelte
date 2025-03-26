<script lang="ts">
    import * as Tabs from "$lib/components/ui/tabs"
    import type { ProjectState } from "$lib/project_state.svelte.ts"
    import { Button } from "$lib/components/ui/button/index"
    import CreationModal from "../tasks/CreationModal.svelte";
    import DueDate from "$lib/components/DueDate.svelte";
    import TimeEstimate from "$lib/components/TimeEstimate.svelte";

    let { project }: { project: ProjectState } = $props()

    async function createPoll() {
        //TODO: send info to create poll
    }

</script>

<Tabs.Content value="polls">
    <div>Todo: Team polls here!</div>
    <Button onclick={createPoll}>Create New Poll</Button>
</Tabs.Content>
