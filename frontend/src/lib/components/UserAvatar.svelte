<script lang="ts">
    import * as Avatar from "$lib/components/ui/avatar"
    import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"
    import type { ProjectState } from "$lib/project_state.svelte"

    let { project, userID }: { project: ProjectState | null; userID: string } = $props()

    let image: Promise<string | null | undefined> = $derived.by(async () => {
        const res = await fetch(PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/getprofilepic?id=" + userID, {
            method: "GET",
            mode: "cors",
            credentials: "include",
            headers: { "Content-Type": "application/json" },
        })

        const contentType = res.headers.get("Content-Type")
        if (contentType == null) {
            return null
        }
        if (contentType.includes("application/json")) {
            const data = await res.json()
            if (!data.found) {
                console.log("User has no profile picture")
                return null // Set to empty or default avatar
            }
        } else {
            // Response is an image, process it
            const blob = await res.blob()
            return URL.createObjectURL(blob) ?? null
        }
    })

    let initial = $derived.by(() => {
        if (project == null) return ""
        let user = project.users.find(user => user.id === userID)
        if (user) {
            return (user.firstName[0] + user.lastName[0]).toUpperCase()
        } else {
            return ""
        }
    })
</script>

<Avatar.Root>
    {#await image then image}
        <Avatar.Image src={image} />
    {/await}
    <Avatar.Fallback>{initial}</Avatar.Fallback>
</Avatar.Root>
