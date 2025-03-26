<script lang="ts">
    import { goto } from "$app/navigation"
    import { base } from "$app/paths"
    import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"
    import Button from "$lib/components/ui/button/button.svelte"
    import * as Card from "$lib/components/ui/card"
    import { joinProject } from "$lib/restApi"

    const inviteInfo = fetch(PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/join/info" + location.search, {
        mode: "cors",
        credentials: "include",
    }).then(res => res.json())

    function decline() {
        window.close.bind(window)
        goto(base + "/")
    }

    async function accept() {
        const projectId = (await inviteInfo).id
        const resp = await joinProject(projectId)
        if (resp.status !== 200) {
            goto(`${location.origin}/registration${location.search}&project=${projectId}`)
        }
    }
</script>

<div class="flex items-center justify-center">
    <Card.Root>
        {#await inviteInfo}
            <Card.Header>
                <Card.Title>Validating invite...</Card.Title>
            </Card.Header>
        {:then inviteInfo}
            {#if inviteInfo.exists}
                <Card.Header>
                    <Card.Title>You've been invited to "{inviteInfo.title}"</Card.Title>
                    <Card.Description>{inviteInfo.description}</Card.Description>
                </Card.Header>
                <Card.Footer class="flex justify-between">
                    <Button onclick={accept}>Accept</Button>
                    <Button
                        onclick={decline}
                        variant="outline">Decline</Button
                    >
                </Card.Footer>
            {:else}
                <Card.Header>
                    <Card.Title>Invalid or expired invite :(</Card.Title>
                </Card.Header>
                <Card.Footer>
                    <Button onclick={decline}>Close</Button>
                </Card.Footer>
            {/if}
        {/await}
    </Card.Root>
</div>
