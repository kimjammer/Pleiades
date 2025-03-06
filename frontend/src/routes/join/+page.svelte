<script lang="ts">
    import { goto } from "$app/navigation"
    import { PUBLIC_API_HOST } from "$env/static/public"
    import Button from "$lib/components/ui/button/button.svelte"
    import * as Card from "$lib/components/ui/card"

    const inviteInfo = fetch("http://" + PUBLIC_API_HOST + "/join/info" + location.search, {
        mode: "cors",
        credentials: "include",
    }).then(res => res.json())

    function decline() {
        window.close.bind(window)
        location.assign("/")
    }

    async function accept() {
        const projectId = (await inviteInfo).id
        if (
            (
                await fetch("http://" + PUBLIC_API_HOST + "/join" + location.search, {
                    mode: "cors",
                    credentials: "include",
                })
            ).status === 200
        ) {
            goto(location.origin + "/project?id=" + projectId)
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
                    <Button on:click={accept}>Accept</Button>
                    <Button
                        on:click={decline}
                        variant="outline">Decline</Button
                    >
                </Card.Footer>
            {:else}
                <Card.Header>
                    <Card.Title>Invalid or expired invite :(</Card.Title>
                </Card.Header>
                <Card.Footer>
                    <Button on:click={decline}>Close</Button>
                </Card.Footer>
            {/if}
        {/await}
    </Card.Root>
</div>
