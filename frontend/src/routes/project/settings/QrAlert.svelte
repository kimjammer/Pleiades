<script lang="ts">
    import { base } from "$app/paths"
    import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"
    import { buttonVariants } from "$lib/components/ui/button/index.js"
    import * as Dialog from "$lib/components/ui/dialog/index.js"
    import { Skeleton } from "$lib/components/ui/skeleton"
    import { toast } from "svelte-sonner"

    let url = $state<string>()
    let qrDialogOpen = $state(false)

    async function generateLink(): Promise<string> {
        let response = await fetch(
            PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/invite" + location.search,
            {
                mode: "cors",
                credentials: "include",
            },
        )

        let token = await response.text()
        if (response.status !== 200) {
            toast(JSON.parse(token).error)
            qrDialogOpen = false
        }

        return location.origin + base + "/join?id=" + token
    }

    async function showLink() {
        url = undefined
        url = await generateLink()
    }
</script>

<Dialog.Root bind:open={qrDialogOpen}>
    <Dialog.Trigger
        class={buttonVariants({ variant: "outline" })}
        onclick={showLink}>QR/link</Dialog.Trigger
    >
    <Dialog.Content class="sm:max-w-[425px]">
        <Dialog.Header>
            <Dialog.Title>Invite Teammates</Dialog.Title>
            <Dialog.Description>
                Have everyone scan or visit this link, which expires in a week.
            </Dialog.Description>
        </Dialog.Header>
        {#if url === undefined}
            <Skeleton class="mx-auto h-[150px] w-[150px] columns-1" />
            loading...
        {:else}
            <div class="grid-columns-1 grid grid-rows-1">
                <Skeleton class="col-start-1 row-start-1 mx-auto h-[150px] w-[150px] columns-1" />
                <img
                    src={"https://quickchart.io/qr?text=" + url}
                    alt="loading..."
                    class="z-50 col-start-1 row-start-1 mx-auto h-[150px] w-[150px] columns-1"
                />
            </div>
            <a
                href={url}
                target="_blank">{url}</a
            >
        {/if}

        <Dialog.Footer>
            <!-- TODO: s<Button type="submit">Done</Button> -->
        </Dialog.Footer>
    </Dialog.Content>
</Dialog.Root>
