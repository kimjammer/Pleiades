<script lang="ts">
    import * as Tabs from "$lib/components/ui/tabs"
    import type { ProjectState } from "$lib/project_state.svelte"

    let { project }: { project: ProjectState } = $props()
</script>

<Tabs.Content value="calendar">
    <div>Todo: Monthly calendar view here!!!</div>
</Tabs.Content>
