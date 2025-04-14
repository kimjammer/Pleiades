<script lang="ts">
    ;(window as any).handleCredentialResponse = (googleUser: any) => {
        debugger
        var profile = googleUser.getBasicProfile()
        console.log("ID: " + profile.getId()) // Do not send to your backend! Use an ID token instead.
        console.log("Name: " + profile.getName())
        console.log("Image URL: " + profile.getImageUrl())
        console.log("Email: " + profile.getEmail()) // This is null if the 'email' scope is not present.
    }
</script>

<svelte:head>
    <script
        src="https://accounts.google.com/gsi/client"
        async
        defer
    ></script>
</svelte:head>

<div
    id="g_id_onload"
    data-client_id="104619113313-21gij0dmhi857nk0c4tf7gjbgfc6jg82"
    data-ux_mode="redirect"
    data-login_uri="https://www.example.com/your_login_endpoint"
></div>
<div
    class="g_id_signin"
    data-type="standard"
></div>
