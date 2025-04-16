<script>
    import { goto } from "$app/navigation"
    import { base } from "$app/paths"
    import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"
    import { Button } from "$lib/components/ui/button"
    import UserAvatar from "$lib/components/UserAvatar.svelte"
    import { Home, LogOut, MessageCircleWarning, Moon, Sun } from "lucide-svelte"
    import { toggleMode } from "mode-watcher"
    import { onMount } from "svelte"
    import { toast } from "svelte-sonner"

    let loggedIn = $state(false)
    async function verifySession() {
        const url = PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/projects"
        const res = await fetch(url, { mode: "cors", credentials: "include" })
        loggedIn = res.ok
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
        // do not remove myId from localStorage because it is usedto determine
        // if user has an account and should be directed to register or login page
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
        <a href={base + (loggedIn ? "/home" : "/")}>
            <h1 class="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl">Pleiades</h1>
        </a>
    </div>
    <div class="flex gap-2 sm:gap-5">
        <a
            href="https://docs.google.com/forms/d/e/1FAIpQLSfPaVUSroOYBp4TMJ7jshZ1ShgtSos-P045BgL6TDW3FijxIA/viewform"
        >
            <Button
                variant="ghost"
                size="icon"><MessageCircleWarning /></Button
            >
        </a>
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
