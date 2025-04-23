<script lang="ts">
    import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"
    import * as Avatar from "$lib/components/ui/avatar"
    import type { ProjectState } from "$lib/project_state.svelte"
    import type { UserId } from "$lib/schema"

    let { project, userID }: { project: ProjectState | null; userID: UserId } = $props()

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
                //console.log("User has no profile picture")
                return null // Use default avatar
            }
        } else {
            // Response is an image, process it
            const blob = await res.blob()
            return URL.createObjectURL(blob) ?? null
        }
    })

    let initial = $derived.by(async () => {
        if (project == null) {
            const res = await fetch(PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/fetchName?id=" + userID, {
                method: "GET",
                mode: "cors",
                credentials: "include",
                headers: { "Content-Type": "application/json" },
            })

            if (!res.ok) {
                throw new Error("User not found")
            }
            const data = await res.json()
            const firstName = data.firstName || ""
            const lastName = data.lastName || ""
            //console.log(firstName.charAt(0) + lastName.charAt(0))
            return firstName.charAt(0) + lastName.charAt(0)
        }
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
    {#await initial then initial}
        <Avatar.Fallback>{initial}</Avatar.Fallback>
    {/await}
</Avatar.Root>
