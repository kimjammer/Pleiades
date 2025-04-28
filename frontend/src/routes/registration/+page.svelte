<script lang="ts">
    import { goto } from "$app/navigation"
    import { base } from "$app/paths"
    import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"
    import { Button } from "$lib/components/ui/button/index"
    import { Input } from "$lib/components/ui/input/index"
    import { recordEvent, tryJoinProject } from "$lib/restApi"
    import { GOOGLE_OAUTH_CLIENT_ID } from "$lib/utils"

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
        }
        const res = await fetch(PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/register/new", {
            method: "POST",
            mode: "cors",
            credentials: "include",
            body: JSON.stringify(user),
        })
        console.log("registration data sent")

        const data = await res.json()
        console.log(data) // Handle success or error messages
        postRegister(data)
    }

    async function postRegister(data: any) {
        if (data.success) {
            recordEvent("users")
            localStorage.myId = data.userId
            if (!(await tryJoinProject())) await goto(base + "/home")
        }
    }

    function hasUpperandLowerCase() {
        return /[A-Z]/.test(password) && /[a-z]/.test(password)
    }

    async function checkAccountExists(email: string) {
        // TODO: this check only occurs on client, thus duplicate accounts can be made. It should happen on server
        console.log("account Exists checking")

        const res = await fetch(
            PUBLIC_PROTOCOL +
                PUBLIC_API_HOST +
                `/register/check?email=${encodeURIComponent(email)}`,
            {
                method: "GET",
                mode: "cors",
                credentials: "include",
                headers: {
                    "Content-Type": "application/json",
                },
            },
        )
        console.log("awaiting")
        return await res.json()
    }

    //This func is activated on submit, confirms validity of input to send to server
    //if input invalid, pops up error
    //NOTE: isValid() does NOT handle already existing accounts.
    async function isValid(ev: MouseEvent) {
        ev.preventDefault()
        if (!email.includes("@")) {
            error = "Invalid Email"
            return
        }
        //FIRST: Check for if account exists
        const data = await checkAccountExists(email)
        console.log(data)
        if (data.exists) {
            error = "There is already an account registered with this email!"
        } else if (
            password.length < 8 ||
            !hasUpperandLowerCase() ||
            !/[`!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~]/.test(password)
        ) {
            error =
                "Your password must be at least 8 characters, " +
                "have at least one upper case and lower case letter, and contain at least one special character."
        } else if (password !== passwordConfirm) {
            error = "Your passwords do not match"
        } else {
            error = ""
            await register()
        }
    }

    window.googleRegister = async ({ credential }: any) => {
        const data = await (
            await fetch(PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/register/google", {
                method: "POST",
                mode: "cors",
                credentials: "include",
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded",
                },
                body: new URLSearchParams({ credential }).toString(),
            })
        ).json()
        postRegister(data)
    }
</script>

<svelte:head>
    <script
        src="https://accounts.google.com/gsi/client"
        async
        defer
    ></script>
</svelte:head>

<div class="grid grid-cols-1 grid-rows-3 sm:h-dvh sm:grid-cols-3 sm:grid-rows-1">
    <div class="bg-primary dark:bg-secondary row-span-1 p-10 sm:col-span-1">
        <a href={base + "/"}>
            <h1
                class="text-primary-foreground dark:text-secondary-foreground scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl"
            >
                Pleiades
            </h1>
        </a>
    </div>
    <div
        class="row-span-2 flex flex-col items-center p-10 sm:col-span-2 sm:items-start sm:justify-center"
    >
        <h1 class="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl">
            Register a New Account
        </h1>
        <p class="scroll-m-20 border-b pb-2 text-2xl tracking-tight transition-colors first:mt-0">
            Please enter a valid email and password
        </p>
        <br />
        <form class="grid w-full max-w-sm items-center gap-1.5">
            {#if error}
                <div class="rounded-lg bg-red-100 p-2 text-red-700">
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
                type="text"
                id="First name"
                placeholder="First Name"
                bind:value={firstname}
            />
            <Input
                type="text"
                id="Last Name"
                placeholder="Last Name"
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
                placeholder="Confirm password"
                bind:value={passwordConfirm}
            />
            <Button onclick={isValid}>Register</Button>

            <div
                id="g_id_onload"
                data-client_id={GOOGLE_OAUTH_CLIENT_ID}
                data-context="signup"
                data-ux_mode="popup"
                data-callback="googleRegister"
                data-auto_prompt="false"
            ></div>
            <div
                class="g_id_signin"
                data-type="standard"
                data-shape="rectangular"
                data-theme="outline"
                data-text="signup_with"
                data-size="large"
                data-logo_alignment="left"
            ></div>

            <Button
                variant="link"
                onclick={() => goto(base + "/login" + location.search)}>Login</Button
            >
        </form>
    </div>
</div>
