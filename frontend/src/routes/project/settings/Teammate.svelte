<script lang="ts">
    import * as Avatar from "$lib/components/ui/avatar"
    import type { UserInProject } from "$lib/project_state.svelte"

    let { details }: { details: UserInProject } = $props()
    $effect(() => console.log("DETAILS", details))
</script>

<div class="flex items-center justify-between space-x-4">
    <div class="flex items-center space-x-4">
        <Avatar.Root>
            <Avatar.Image
                src="/avatars/02.png"
                alt={details.firstName}
            />
            <Avatar.Fallback>{details.firstName[0]}{details.lastName[0]}</Avatar.Fallback>
        </Avatar.Root>
        <div>
            <p class="text-sm font-medium leading-none">{details.firstName} {details.lastName}</p>
            <p class="text-muted-foreground text-sm">{details.email}</p>
        </div>
    </div>
</div>
