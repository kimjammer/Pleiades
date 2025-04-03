<script>
    import { goto } from "$app/navigation"
    import { base } from "$app/paths"
    import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"
    import { Button } from "$lib/components/ui/button"
    import { Home, LogOut, Moon, Sun, UserRound } from "lucide-svelte"
    import { toggleMode } from "mode-watcher"
    import { onMount } from "svelte"
    import { toast } from "svelte-sonner"
    import UserAvatar from "$lib/components/UserAvatar.svelte";

    let loggedIn = $state(false)
    async function verifySession() {
        const url = PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/projects"
        const res = await fetch(url, { mode: "cors", credentials: "include" })
        if (res.ok) {
            loggedIn = true
        } else {
            loggedIn = false
        }
    }

    async function logout() {
        const url = PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/logout"
        const res = await fetch(url, {
            method: "POST",
            mode: "cors",
            credentials: "include",
            headers: { "Content-Type": "application/json" },
        })

        //Error handling
        if (!res.ok) {
            toast.error("Failed to logout")
            return
        }
        localStorage.removeItem("myId")
        toast.success("Logged out")
        loggedIn = false
        await goto(base + "/")
    }

    onMount(async () => {
        //Load cached login status to prevent flicker
        loggedIn = localStorage.myId !== undefined

        await verifySession()
    })
</script>

<div class="flex justify-between border-b p-5">
    <div>
        <a href={base + "/"}>
            <h1 class="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl">Pleiades</h1>
        </a>
    </div>
    <div class="flex gap-2 sm:gap-5">
        <Button
            onclick={toggleMode}
            variant="ghost"
            size="icon"
        >
            <Sun
                class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
            />
            <Moon
                class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
            />
            <span class="sr-only">Toggle theme</span>
        </Button>
        {#if loggedIn}
            <Button
                href={base + "/home"}
                size="icon"
                variant="secondary"
            >
                <Home />
            </Button>
            <Button
                href={base + "/account"}
                size="icon"
                variant="secondary"
            >
<!--                <UserRound />-->
                <UserAvatar
                        project={null}
                        userID={localStorage.myId}
                />
            </Button>
            <Button
                onclick={logout}
                size="icon"
            >
                <LogOut />
            </Button>

        {:else}
            <Button href={base + "/login"}>Login</Button>
        {/if}
    </div>
</div>
