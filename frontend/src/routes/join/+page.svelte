<script lang="ts">
    import { PUBLIC_API_HOST } from "$env/static/public"

    const isValid = fetch("http://" + PUBLIC_API_HOST + "/join/validate" + location.search, {
        mode: "cors",
        credentials: "include",
    })
        .then(res => res.text())
        .then(validity => validity === "true")
</script>

{#await isValid}
    validating...
{:then isValid}
    {isValid}
{/await}
