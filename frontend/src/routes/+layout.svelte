<script lang="ts">
    import "../app.css"
    import { ModeWatcher } from "mode-watcher"
    import { Toaster } from "$lib/components/ui/sonner/index"
</script>

<nav>
    <a href="/">Landing</a>
    <a href="/login">Login</a>
    <a href="/registration">Register</a>
    <a href="/home">User_Home</a>
    <a href="/account">Account</a>
</nav>

<ModeWatcher />
<Toaster richColors />
<slot></slot>
