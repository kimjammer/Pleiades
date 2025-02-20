<script lang="ts">
	import '../app.css';
	let { children } = $props();
</script>

<nav>
    <a href="/">Home</a>
    <a href="/login">Login</a>
    <a href="/registration">Register</a>
    <a href="/dashboard">Dashboard</a>
</nav>

{@render children()}
