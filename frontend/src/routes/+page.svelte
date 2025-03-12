<script>
    import { Button } from "$lib/components/ui/button/index"
    import { goto } from "$app/navigation"
    import { PUBLIC_API_HOST } from "$env/static/public"
    import { toast } from "svelte-sonner"

    //TODO: Remove me when not needed
    const fakeLogin = async () => {
        const url = "http://" + PUBLIC_API_HOST + "/fakelogin"
        const res = await fetch(url, { mode: "cors", credentials: "include" })
        if (res.ok) {
            toast.success("Successfully logged in as first found user")
        } else {
            toast.error("Failed to login")
        }
    }
</script>

<div>
    <h1 class="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl">Pleiades</h1>

    <h2
        class="scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0"
    >
        Collaborate Easier...or else...
    </h2>
    <Button onclick={() => goto("/login")}>Login</Button>
    <Button onclick={() => goto("/registration")}>Create an account</Button>

    <Button onclick={fakeLogin}>Login as some user</Button>
</div>
