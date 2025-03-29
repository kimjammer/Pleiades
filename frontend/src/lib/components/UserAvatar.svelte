<script lang="ts">
    import * as Avatar from "$lib/components/ui/avatar"
    import type { ProjectState } from "$lib/project_state.svelte"

    let { project, userID }: { project: ProjectState; userID: string } = $props()

    let image = $derived.by(() => {
        //TODO: Convert userPhoto to correct url/format if applicable
        return project.users.find(user => user.id === userID)?.userPhoto ?? ""
    })

    let initial = $derived.by(() => {
        let user = project.users.find(user => user.id === userID)
        if (user) {
            return (user.firstName[0] + user.lastName[0]).toUpperCase()
        } else {
            return ""
        }
    })
</script>

<Avatar.Root>
    <Avatar.Image src={image} />
    <Avatar.Fallback>{initial}</Avatar.Fallback>
</Avatar.Root>
