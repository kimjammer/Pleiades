<script>
    import { Input } from "$lib/components/ui/input/index"
    import { Button } from "$lib/components/ui/button/index"
    import { PUBLIC_API_HOST } from "$env/static/public"

    let email = ""
    let password = ""
    let passwordConfirm = ""

    //TODO: create link to dashboard

    async function register() {
        const res = await fetch("http://" + PUBLIC_API_HOST + "/api/register", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ email, password }),
        })

        const data = await res.json()
        console.log(data) // Handle success or error messages
    }

    //This func is activated on submit, confirms validity of input to send to server
    //if input invalid, pops up error
    //NOTE: isValid() does NOT handle already existing accounts.
    function isValid() {
        if (password !== passwordConfirm) {
        }
        //TODO: Implement email and passoword checking

        //TODO: Create error messages for invalid input
    }

    //pop up error box with link to login page for already existing accoutns
    function accountExists() {
        //TODO: create pop-up box with link to login page
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
        type="confirm password"
        id="confirm password"
        placeholder="confirm Password"
        bind:value={passwordConfirm}
    />
    <Button on:click={isValid()}>Register</Button>
</div>
