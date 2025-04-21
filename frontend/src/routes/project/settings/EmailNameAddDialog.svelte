<script lang="ts">
    import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"
    import { Button } from "$lib/components/ui/button"
    import * as Dialog from "$lib/components/ui/dialog"
    import { Input } from "$lib/components/ui/input"
    import { debounce } from "$lib/utils"
    import { toast } from "svelte-sonner"
    import type { ChangeEventHandler } from "svelte/elements"

    let { description = "" }: { description?: string } = $props()

    // 1st is name, 2nd is email
    type Suggestion = [string, string]

    let isLoading = $state(false)
    let suggestions = $state([] as Suggestion[])

    const suggestCompletions: ChangeEventHandler<HTMLInputElement> = debounce(async event => {
        const nameOrEmail = (event.target as HTMLInputElement).value
        isLoading = true
        const url =
            PUBLIC_PROTOCOL +
            PUBLIC_API_HOST +
            "/directory/purdue.edu?name=" +
            encodeURIComponent(nameOrEmail)
        const res = await fetch(url, { mode: "cors", credentials: "include" })
        isLoading = false
        if (!res.ok) {
            toast.error("Could not fetch suggestions.")
            return
        }
        suggestions = await res.json()
        if (suggestions.length === 0 && nameOrEmail.includes("@")) {
            suggestions = [[nameOrEmail.split("@")[0], nameOrEmail]]
            return
        }
    }, 500)

    async function goInvite(ev: MouseEvent) {
        const elem = ev.target as HTMLButtonElement
        const email = elem.value
        const name = elem.innerText
        const params = new URLSearchParams(document.location.search)
        const projectId = params.get("id") || ""
        const url =
            PUBLIC_PROTOCOL +
            PUBLIC_API_HOST +
            `/invite/email?name=${name}&email=${email}&id=${projectId}`
        const res = await fetch(url, { mode: "cors", credentials: "include" })
        if (res.ok) {
            toast.success("Invite email sent!")
        } else {
            toast.error("Could not invite user.")
        }
    }
</script>

<Dialog.Content class="sm:max-w-[425px]">
    <Dialog.Header>
        <Dialog.Title>Invite Teammates</Dialog.Title>
        <Dialog.Description>
            {description}
        </Dialog.Description>
    </Dialog.Header>

    <!-- Command makes sense for this, but it didn't like me dynamically changing the items -->
    <Input
        type="text"
        placeholder="Name or email"
        class="w-full"
        oninput={suggestCompletions}
    />

    {#if suggestions.length}
        <div class="max-h-96">
            {#each suggestions as [name, email]}
                <Button
                    variant="ghost"
                    value={email}
                    onclick={goInvite}>{name}</Button
                >
            {/each}
        </div>
    {:else}
        {isLoading ? "Loading..." : "No results found."}
    {/if}

    <Dialog.Footer>
        <!-- TODO: s<Button type="submit">Done</Button> -->
    </Dialog.Footer>
</Dialog.Content>
