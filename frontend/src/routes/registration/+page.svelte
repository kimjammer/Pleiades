

<script>
    import { Input } from "$lib/components/ui/input/index"
    import { Button } from "$lib/components/ui/button/index"

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

    function isValid() {
        //TODO: Implement email and passoword checking

        //TODO: Create error messages for invalid input

    }

</script>

<h1 class="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl">
    Register a New Account
</h1>

<p
        class="scroll-m-20 border-b pb-2 text-2xl font- tracking-tight transition-colors first:mt-0"
>
    Please enter a valid email and password
</p>
<br />
<div class="grid w-full max-w-sm items-center gap-1.5">
    <Input type="email" id="email" placeholder="Email" />
    <Input type="password" id="password" placeholder="Password" />
    <Input type="confirm password" id="confirm password" placeholder="confirm Password" />
    <Button on:click={isValid()}>Register</Button>
</div>
