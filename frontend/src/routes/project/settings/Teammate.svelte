<script lang="ts">
    import UserAvatar from "$lib/components/UserAvatar.svelte"
    import { ProjectState, type UserInProject } from "$lib/project_state.svelte"

    let { details, project }: { details: UserInProject; project: ProjectState } = $props()
    $effect(() => console.log("DETAILS", details))
</script>

<div class="flex items-center justify-between space-x-4">
    <div class="flex items-center space-x-4">
        <UserAvatar
            {project}
            userID={details.id}
        />
        <div>
            <p class="text-sm font-medium leading-none">{details.firstName} {details.lastName}</p>
            <p class="text-sm text-blue-500">
                <a href="mailto:{details.email}">Contact</a>
            </p>
        </div>
    </div>
</div>
