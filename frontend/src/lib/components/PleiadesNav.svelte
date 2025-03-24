<script>
    import { PUBLIC_API_HOST } from "$env/static/public"
    import { toast } from "svelte-sonner"
    import { toggleMode } from "mode-watcher"
    import { onMount } from "svelte"
    import { Button } from "$lib/components/ui/button"
    import { Sun, Moon } from "lucide-svelte"
    import { goto } from "$app/navigation"

    let loggedIn = $state(false)
    async function verifySession() {
        const url = "http://" + PUBLIC_API_HOST + "/projects"
        const res = await fetch(url, { mode: "cors", credentials: "include" })
        if (res.ok) {
            loggedIn = true
        } else {
            loggedIn = false
        }
    }

    async function logout() {
        const url = "http://" + PUBLIC_API_HOST + "/logout"
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
        goto("/")
    }

    function login() {
        goto("/login")
    }

    function account() {
        goto("/account")
    }

    function home() {
        goto("/home")
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
            <Button onclick={logout}>Logout</Button>
        {:else}
            <Button onclick={login}>Login</Button>
        {/if}
    </div>
</div>
