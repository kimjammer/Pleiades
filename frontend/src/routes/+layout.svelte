<script lang="ts">
    import { base } from "$app/paths"
    import { Toaster } from "$lib/components/ui/sonner/index"
    import { ModeWatcher } from "mode-watcher"
    import "../app.css"
</script>

<nav>
    <a href={base}>Landing</a>
    <a href={base + "/login"}>Login</a>
    <a href={base + "/registration"}>Register</a>
    <a href={base + "/home"}>User_Home</a>
    <a href={base + "/account"}>Account</a>
</nav>

<ModeWatcher />
<Toaster richColors />
<slot></slot>
