<script>
    import { Input } from "$lib/components/ui/input/index"
    import { Button } from "$lib/components/ui/button/index"
    import { PUBLIC_API_HOST } from "$env/static/public"
    import {goto} from "$app/navigation";

    //TODO: implement password recovery link & page?

    let email = ""
    let password = ""

    async function register() {
        const res = await fetch("http://" + PUBLIC_API_HOST + "/api/login", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ email, password }),
        })

        const data = await res.json()
        console.log(data) // Handle success or error messages
    }

    function isValid() {
        //TODO: Implement email and password checking
        //TODO: Create error messages for invalid input
    }
</script>

<h1 class="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl">Login</h1>
<br />
<div class="grid w-full max-w-sm items-center gap-1.5">
    <Input
        type="email"
        id="email"
        placeholder="Email"
    />
    <Input
        type="password"
        id="password"
        placeholder="Password"
    />
    <Button on:click={isValid()}>Login</Button>
    <Button variant="link" onclick={() => goto("/registration")}>Register New Account</Button>
</div>
