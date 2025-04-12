<script lang="ts">
    import { base } from "$app/paths"
    import { goto } from "$app/navigation"
    import { Button } from "$lib/components/ui/button"
    import { Input } from "$lib/components/ui/input"
    import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"
    import type { ResetPasswordRequest } from "$lib/schema"
    import { toast } from "svelte-sonner"

    let password = $state("")
    let passwordConfirm = $state("")

    function hasUpperandLowerCase() {
        return /[A-Z]/.test(password) && /[a-z]/.test(password)
    }

    function isValid(): boolean {
        if (
            password.length < 8 ||
            !hasUpperandLowerCase() ||
            !/[`!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~]/.test(password)
        ) {
            toast.error(
                "Your password must be at least 8 characters, have at least one upper case and " +
                    "lower case letter, and contain at least one special character.",
            )
            return false
        } else if (password !== passwordConfirm) {
            toast.error("Your passwords do not match.")
            return false
        } else {
            return true
        }
    }

    async function passwordReset() {
        if (!isValid()) return

        let urlParams = new URLSearchParams(document.location.search)
        let token = urlParams.get("token") ?? ""

        const url = PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/resetPassword"
        const body: ResetPasswordRequest = { token, newPassword: password }
        const res = await fetch(url, {
            method: "POST",
            mode: "cors",
            credentials: "include",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(body),
        })

        if (res.status !== 200) {
            toast.error("Failed to reset password.")
        } else {
            toast.success("Password reset successfully.")
            await goto(base + "/login")
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
            Reset Password
        </h1>
        <small class="mb-5 text-sm font-medium leading-none text-muted-foreground">
            Enter a new password.
        </small>
        <form class="grid w-full max-w-sm items-center gap-1.5">
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
            <Button onclick={passwordReset}>Reset password</Button>
        </form>
    </div>
</div>
