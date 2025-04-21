<script lang="ts">
    import * as Card from "$lib/components/ui/card"
    import type { ProjectState } from "$lib/project_state.svelte"
    import EmailNameAdd from "./EmailNameAdd.svelte"
    import QrAlert from "./QrAlert.svelte"
    import Teammate from "./Teammate.svelte"

    let { project }: { project: ProjectState } = $props()
</script>

<Card.Root class="mt-2 w-fit">
    <Card.Header>
        <Card.Title>Team Members</Card.Title>
    </Card.Header>
    <Card.Content class="grid gap-6">
        {#each project.users as user}
            <Teammate
                details={user}
                {project}
            />
        {/each}
        <div>
            Invite by <QrAlert /><EmailNameAdd />
        </div>
    </Card.Content>
</Card.Root>
