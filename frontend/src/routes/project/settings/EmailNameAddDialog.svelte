<script lang="ts">
    import * as Command from "$lib/components/ui/command"
    import * as Dialog from "$lib/components/ui/dialog"
    import { debounce } from "$lib/utils"
    import type { ChangeEventHandler } from "svelte/elements"
    import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"

    let { description = "" }: { description?: string } = $props()

    let isLoading = $state(false)
    let suggestions = $state([])

    const suggestCompletions: ChangeEventHandler<HTMLInputElement> = debounce(async event => {
        const nameOrEmail = (event.target as HTMLInputElement).value
        isLoading = true
        const url = PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/directory/purdue.edu?name=" +
            encodeURIComponent(nameOrEmail)
        const res = await (await fetch(url, { mode: "cors", credentials: "include" })).json()
        isLoading = false
    }, 500)
</script>

<Dialog.Content class="sm:max-w-[425px]">
    <Dialog.Header>
        <Dialog.Title>Invite Teammates</Dialog.Title>
        <Dialog.Description>
            {description}
        </Dialog.Description>
    </Dialog.Header>

    <Command.Root>
        <Command.Input
            placeholder="Name or email"
            oninput={suggestCompletions}
        />
        <Command.List>
            <Command.Empty>{isLoading ? "Loading..." : "No results found."}</Command.Empty>
            <Command.Item>Calendar</Command.Item>
            <Command.Item>Search Emoji</Command.Item>
            <Command.Item>Calculator</Command.Item>
        </Command.List>
    </Command.Root>

    <Dialog.Footer>
        <!-- TODO: s<Button type="submit">Done</Button> -->
    </Dialog.Footer>
</Dialog.Content>
