

<script>
    let email = '';
    let password = '';
    let passwordConfirm = '';

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

<h1>Register a New Account</h1>
<p>You must enter a valid, unregistered email to create an account.</p>
<br />
<p>Email</p>
<input type="text" placeholder="Email" bind:value={email} />
<br />
<p>Password</p>
<input type="password" placeholder="Password" bind:value={password} />
<br />
<input type="confirm password" placeholder="Confirm Password" bind:value={passwordConfirm} />
