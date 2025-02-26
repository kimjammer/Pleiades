<script>
    import { Input } from "$lib/components/ui/input/index"
    import { Button } from "$lib/components/ui/button/index"
    import { PUBLIC_API_HOST } from "$env/static/public"
    import {goto} from "$app/navigation";

    let email = ""
    let password = ""
    let passwordConfirm = ""
    let error = ""

    //TODO: create link to dashboard

    async function register() {
        const res = await fetch("http://" + PUBLIC_API_HOST + "/register", {
            method: "POST",
            mode: "cors",
            body: JSON.stringify({ email })
        })

        const data = await res.json()
        console.log(data) // Handle success or error messages
    }

    //This func is activated on submit, confirms validity of input to send to server
    //if input invalid, pops up error
    //NOTE: isValid() does NOT handle already existing accounts.
    async function isValid() {
        //TODO: Implement email and password checking
        //TODO: Create error messages for invalid input
        //FIRST: Check for if account exists
        console.log("account Exists checking")
        //TODO: check server to see if account exists
        const res = await fetch("http://" + PUBLIC_API_HOST + `/api/register/check?email=${encodeURIComponent(email)}`, {
            mode: "cors",
        })

        const data = await res.json()
        console.log(data)
        if (data.exists) {
            error = "There is already an account registered with this email!"
        }
        else if (password.length < 8) {
            error = "Your password must be at least 8 characters"
        }
        else if (password !== passwordConfirm) {
            error = "Your passwords do not match"
        }
        else {
            await register()
        }

    }
    
</script>

<h1 class="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl">
    Register a New Account
</h1>

<p class="font- scroll-m-20 border-b pb-2 text-2xl tracking-tight transition-colors first:mt-0">
    Please enter a valid email and password
</p>
<br />
<div class="grid w-full max-w-sm items-center gap-1.5">
    {#if error}
        <div class="p-2 bg-red-100 text-red-700 rounded-lg">
            {error}
        </div>
    {/if}
    <!--Create three input fields and then bind them to three variables-->
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
    <Input
        type="password"
        id="confirm password"
        placeholder="confirm Password"
        bind:value={passwordConfirm}
    />
    <Button on:click={isValid()}>Register</Button>
</div>
