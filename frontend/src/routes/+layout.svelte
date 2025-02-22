<script lang="ts">
    import "../app.css"
    let { children } = $props()
</script>

<nav>
    <a href="/">Landing</a>
    <a href="/login">Login</a>
    <a href="/registration">Register</a>
    <a href="/home">User_Home</a>
</nav>

{@render children()}
