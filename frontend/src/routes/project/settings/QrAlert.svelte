<script lang="ts">
    import { PUBLIC_API_HOST } from "$env/static/public"
    import { buttonVariants } from "$lib/components/ui/button/index.js"
    import * as Dialog from "$lib/components/ui/dialog/index.js"

    let url = new Promise<string>(() => {})

    async function generateLink(): Promise<string> {
        let response = await fetch("http://" + PUBLIC_API_HOST + "/invite" + location.search, {
            mode: "cors",
            credentials: "include",
        })

        let token = await response.text()

        return location.origin + "/join?id=" + token
    }
</script>

<Dialog.Root>
    <Dialog.Trigger
        class={buttonVariants({ variant: "outline" })}
        onclick={() => (url = generateLink())}>Invite</Dialog.Trigger
    >
    <Dialog.Content class="sm:max-w-[425px]">
        <Dialog.Header>
            <Dialog.Title>Invite Teammates</Dialog.Title>
            <Dialog.Description>
                Have everyone scan or visit this link, which expires in a week.
            </Dialog.Description>
        </Dialog.Header>
        {#await url}
            loading...
        {:then url}
            <img
                src={"https://quickchart.io/qr?text=" + url}
                alt="loading..."
                class="mx-auto"
            />
            <a
                href={url}
                target="_blank">{url}</a
            >
        {/await}

        <Dialog.Footer>
            <!-- TODO: s<Button type="submit">Done</Button> -->
        </Dialog.Footer>
    </Dialog.Content>
</Dialog.Root>
