<script>
    import { goto } from "$app/navigation"
    import { base } from "$app/paths"
    import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"
    import { Button } from "$lib/components/ui/button/index"
    import { Input } from "$lib/components/ui/input/index"
    import { tryJoinProject } from "$lib/restApi"

    //TODO: implement password recovery link & page?

    let email = ""
    let password = ""
    let error = ""

    async function login() {
        const res = await fetch(
            PUBLIC_PROTOCOL +
                PUBLIC_API_HOST +
                `/login?email=${encodeURIComponent(email)}&password=${encodeURIComponent(password)}`,
            {
                method: "GET",
                mode: "cors",
                credentials: "include", //for cookies
                headers: { "Content-Type": "application/json" },
            },
        )

        const data = await res.json()
        if (data.exists) {
            error = ""
            localStorage.myId = data.userId
            if (!(await tryJoinProject())) await goto(base + "/home")
        } else {
            error = "Invalid Login"
        }
    }
</script>

<h1 class="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl">Login</h1>
<br />
<div class="grid w-full max-w-sm items-center gap-1.5">
    {#if error}
        <div class="rounded-lg bg-red-100 p-2 text-red-700">
            {error}
        </div>
    {/if}
    <Input
        type="email"
        id="email"
        placeholder="Email"
        bind:value={email}
    />
    <Input
        type="password"
        id="password"
        placeholder="Password"
        bind:value={password}
    />
    <Button onclick={login}>Login</Button>
    <Button
        variant="link"
        onclick={() => goto(base + "/registration")}>Register New Account</Button
    >
</div>
