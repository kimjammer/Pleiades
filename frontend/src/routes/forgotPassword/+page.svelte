<script lang="ts">
    import { base } from "$app/paths"
    import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"
    import { Button } from "$lib/components/ui/button/index"
    import { Input } from "$lib/components/ui/input/index"
    import { toast } from "svelte-sonner"
    import type { ForgotPasswordRequest } from "$lib/schema.js"

    let email = ""

    async function forgotPassword() {
        const url = PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/forgotPassword"
        const body: ForgotPasswordRequest = { email }
        const res = await fetch(url, {
            method: "POST",
            mode: "cors",
            credentials: "include", //for cookies
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(body),
        })

        if (res.status !== 200) {
            toast.error("Failed to request password reset.")
        } else {
            toast.success("Password reset email sent.")
        }
    }
</script>

<div class="grid grid-cols-1 grid-rows-3 sm:h-dvh sm:grid-cols-3 sm:grid-rows-1">
    <div class="row-span-1 bg-primary p-10 dark:bg-secondary sm:col-span-1">
        <a href={base + "/"}>
            <h1
                class="scroll-m-20 text-4xl font-extrabold tracking-tight text-primary-foreground dark:text-secondary-foreground lg:text-5xl"
            >
                Pleiades
            </h1>
        </a>
    </div>
    <div
        class="row-span-2 flex flex-col items-center p-10 sm:col-span-2 sm:items-start sm:justify-center"
    >
        <h1 class="scroll-m-20 pb-5 text-4xl font-extrabold tracking-tight lg:text-5xl">
            Forgot Password?
        </h1>
        <small class="mb-5 text-sm font-medium leading-none text-muted-foreground">
            Enter the email address you used to register, and we will send you a link to reset your
            password.
        </small>
        <div class="grid w-full max-w-sm items-center gap-1.5">
            <Input
                type="email"
                id="email"
                placeholder="Email"
                bind:value={email}
            />
            <Button onclick={forgotPassword}>Send Email</Button>
        </div>
    </div>
</div>
