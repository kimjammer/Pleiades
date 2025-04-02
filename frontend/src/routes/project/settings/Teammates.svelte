<script lang="ts">
    import * as Card from "$lib/components/ui/card"
    import type { ProjectState } from "$lib/project_state.svelte"
    import Teammate from "./Teammate.svelte"

    let { project }: { project: ProjectState } = $props()
</script>

<Card.Root class="mt-2 w-fit">
    <Card.Header>
        <Card.Title>Team Members</Card.Title>
        <Card.Description>Invite your team members to collaborate</Card.Description>
    </Card.Header>
    <Card.Content class="grid gap-6">
        {#each project.users as user}
            <Teammate details={user} />
        {/each}
    </Card.Content>
</Card.Root>
