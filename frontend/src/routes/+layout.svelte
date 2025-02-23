<script lang="ts">
    import "../app.css"
    import { ModeWatcher } from "mode-watcher"
    let { children } = $props()
</script>

<nav>
    <a href="/">Landing</a>
    <a href="/login">Login</a>
    <a href="/registration">Register</a>
    <a href="/home">User_Home</a>
</nav>

<ModeWatcher />
{@render children()}
