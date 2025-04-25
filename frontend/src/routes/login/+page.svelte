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

        postLogin(await res.json())
    }

    async function postLogin(data: any) {
        if (data.exists) {
            recordEvent("login")
            error = ""
            localStorage.myId = data.userId
            if (!(await tryJoinProject())) await goto(base + "/home")
        } else {
            error = "Invalid Login"
        }
    }

    window.googleSignin = async ({ credential }: any) => {
        const data = await (
            await fetch(PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/login/google", {
                method: "POST",
                mode: "cors",
                credentials: "include",
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded",
                },
                body: new URLSearchParams({ credential }).toString(),
            })
        ).json()
        postLogin(data)
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
        <h1 class="scroll-m-20 pb-5 text-4xl font-extrabold tracking-tight lg:text-5xl">Login</h1>
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

            <div
                id="g_id_onload"
                data-client_id={GOOGLE_OAUTH_CLIENT_ID}
                data-context="signin"
                data-ux_mode="popup"
                data-callback="googleSignin"
                data-auto_prompt="false"
            ></div>
            <div
                class="g_id_signin"
                data-type="standard"
                data-shape="rectangular"
                data-theme="outline"
                data-text="signin_with"
                data-size="large"
                data-logo_alignment="left"
            ></div>

            <Button
                variant="link"
                href="forgotPassword">Forgot Password?</Button
            >
            <Button
                variant="link"
                href="registration">Register New Account</Button
            >
        </div>
    </div>
</div>
