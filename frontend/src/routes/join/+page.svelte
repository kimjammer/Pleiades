<script lang="ts">
    import { PUBLIC_API_HOST } from "$env/static/public"
    import Button from "$lib/components/ui/button/button.svelte"
    import * as Card from "$lib/components/ui/card"

    const isValid = fetch("http://" + PUBLIC_API_HOST + "/join/validate" + location.search, {
        mode: "cors",
        credentials: "include",
    })
        .then(res => res.text())
        .then(validity => validity === "true")

    function decline() {
        window.close.bind(window)
        location.assign("/")
    }

    async function accept() {
        if (
            (
                await fetch("http://" + PUBLIC_API_HOST + "/join" + location.search, {
                    mode: "cors",
                    credentials: "include",
                })
            ).status === 200
        ) {
            location.assign(location.origin + "/project" + location.search)
        }
    }
</script>

<div class="flex items-center justify-center">
    <Card.Root>
        {#await isValid}
            <Card.Header>
                <Card.Title>Validating invite...</Card.Title>
            </Card.Header>
        {:then isValid}
            {#if isValid}
                <Card.Header>
                    <Card.Title>You've been invited to "project"</Card.Title>
                    <Card.Description>Card Description</Card.Description>
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
