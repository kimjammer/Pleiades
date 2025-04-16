<script lang="ts">
    import type { Poll, ProjectState, Option } from "$lib/project_state.svelte"

    let { project, poll, option }: { project: ProjectState; poll: Poll; option: Option } = $props()

    let sum = $derived(
        option.likedUsers.length + option.neutralUsers.length + option.dislikedUsers.length,
    )
</script>

<div
    class="chart"
    style={`grid-template-columns: ${option.likedUsers.length}fr ${option.neutralUsers.length}fr ${option.dislikedUsers.length}fr ${project.users.length - sum}fr`}
>
    <div class="bar bg-green-500/80"></div>
    <div class="bar bg-yellow-500/80"></div>
    <div class="bar bg-red-500/80"></div>
    <div class="bar bg-slate-500/30"></div>
</div>

<style>
    .chart {
        width: 10em;
        height: 2em;
        display: grid;
        clip-path: inset(0 0 round 0.2em);
    }
</style>
