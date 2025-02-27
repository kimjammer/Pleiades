<script>
    import { Input } from "$lib/components/ui/input/index"
    import { Button } from "$lib/components/ui/button/index"
    import { PUBLIC_API_HOST } from "$env/static/public"
    import {goto} from "$app/navigation";

    let email = ""
    let password = ""
    let passwordConfirm = ""
    let error = ""
    let firstname = ""
    let lastname = ""

    //TODO: create link to dashboard

    async function register() {
        const user = {
            email: email,
            firstname: firstname,
            lastname: lastname,
            password: password,
        };
        const res = await fetch("http://" + PUBLIC_API_HOST + "/register/new", {
            method: "POST",
            mode: "cors",
            body: JSON.stringify(user)
        })
        console.log("registration data sent")

        const data = await res.json()
        console.log(data) // Handle success or error messages
        if (data.success) {
            await goto("/home")
        }
    }

    function hasUpperandLowerCase() {
        return /[A-Z]/.test(password) && /[a-z]/.test(password)
    }

    //This func is activated on submit, confirms validity of input to send to server
    //if input invalid, pops up error
    //NOTE: isValid() does NOT handle already existing accounts.
    async function isValid() {

        if (!email.includes("@")) {
            error = "Invalid Email"
            return
        }
        //FIRST: Check for if account exists
        console.log("account Exists checking")
        const res = await fetch("http://" + PUBLIC_API_HOST + `/register/check?email=${encodeURIComponent(email)}`, {
            method: "GET",
            mode: "cors",
            headers: {
                'Content-Type': 'application/json'
            },
            //body: JSON.stringify({email: email})
        })
        console.log("awaiting")
        const data = await res.json()
        console.log(data)
        if (data.exists) {
            error = "There is already an account registered with this email!"
        }
        else if (password.length < 8 || !hasUpperandLowerCase() || !/[`!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~]/.test(password)) {
            error = "Your password must be at least 8 characters, " +
                        "have at least one upper case and lower case letter, and contain at least one special character."
        }
        else if (password !== passwordConfirm) {
            error = "Your passwords do not match"
        }
        else {
            error = ""
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
            type="First Name"
            id="First name"
            placeholder="First Name"
            bind:value={firstname}
    />
    <Input
            type="Last Name"
            id="Last Name"
            placeholder= "Last Name"
            bind:value={lastname}
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
    <Button onclick={isValid}>Register</Button>
    <Button variant="link" onclick={() => goto("/login")}>Login</Button>

</div>
