<script lang="ts">
    import * as Avatar from "$lib/components/ui/avatar"
    import type { ProjectState } from "$lib/project_state.svelte"
    import {PUBLIC_API_HOST, PUBLIC_PROTOCOL} from "$env/static/public";

    let { project, userID }: { project: ProjectState; userID: string } = $props()

    let image = $derived.by(() => {
        let user = project.users.find(user => user.id === userID)
        //TODO: Fetch image from backend
        return null ?? ""
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
