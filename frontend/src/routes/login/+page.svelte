<script>
    import { Input } from "$lib/components/ui/input/index"
    import { Button } from "$lib/components/ui/button/index"
    import { PUBLIC_API_HOST } from "$env/static/public"
    import {goto} from "$app/navigation";

    //TODO: implement password recovery link & page?

    let email = ""
    let password = ""
    let error = ""

    async function register() {
        const res = await fetch("http://" + PUBLIC_API_HOST + "/api/login", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            credentials: "include",
            body: JSON.stringify({ email, password }),
        })

        const data = await res.json()
        console.log(data) // Handle success or error messages
    }

    async function login() {
        //TODO: Implement email and password checking
        const res = await fetch("http://" + PUBLIC_API_HOST +
                                    `/login?email=${encodeURIComponent(email)}&password=${encodeURIComponent(password)}`, {
            method: "GET",
            mode: "cors",
            headers: { "Content-Type": "application/json" },
        })

        const data = await res.json()
        if (data.exists) {
            error = ""
            await goto("/home")
        }
        else {
            error = "Invalid Login"
        }

    }
</script>

<h1 class="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl">Login</h1>
<br />
<div class="grid w-full max-w-sm items-center gap-1.5">
    {#if error}
        <div class="p-2 bg-red-100 text-red-700 rounded-lg">
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
    <Button variant="link" onclick={() => goto("/registration")}>Register New Account</Button>
</div>
