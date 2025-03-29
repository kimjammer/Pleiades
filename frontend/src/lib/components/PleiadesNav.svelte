<script>
    import { goto } from "$app/navigation"
    import { base } from "$app/paths"
    import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"
    import { Button } from "$lib/components/ui/button"
    import { LogOut, Moon, Sun } from "lucide-svelte"
    import { toggleMode } from "mode-watcher"
    import { onMount } from "svelte"
    import { toast } from "svelte-sonner"

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
        await goto(base + "/")
    }

    function login() {
        goto(base + "/login")
    }

    function account() {
        goto(base + "/account")
    }

    function home() {
        goto(base + "/home")
    }

    onMount(async () => {
        await verifySession()
    })
</script>

<div class="flex justify-between border-b p-5">
    <div>
        <h1 class="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl">Pleiades</h1>
    </div>
    <div class="flex gap-5">
        <Button
            onclick={toggleMode}
            variant="outline"
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
            <Button onclick={account}>Account</Button>
            <Button onclick={home}>Home</Button>
            <Button
                onclick={logout}
                size="icon"
            >
                <LogOut />
            </Button>
        {:else}
            <Button onclick={login}>Login</Button>
        {/if}
    </div>
</div>
