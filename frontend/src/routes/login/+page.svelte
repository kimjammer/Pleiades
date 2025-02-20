

<script>
    let email = '';
    let password = '';

    async function register() {
        const res = await fetch('http://localhost:8080/api/register', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ email, password })
        });

        const data = await res.json();
        console.log(data); // Handle success or error messages
    }
</script>

<h1>Login</h1>
<br />
<input type="text" placeholder="Email" bind:value={email} />
<br />
<input type="password" placeholder="Password" bind:value={password} />