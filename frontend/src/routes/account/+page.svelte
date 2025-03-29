<script>
    import { Input } from "$lib/components/ui/input/index"
    import { Button } from "$lib/components/ui/button/index"
    import { PUBLIC_PROTOCOL, PUBLIC_API_HOST } from "$env/static/public"
    import { goto } from "$app/navigation"
    import { onMount } from "svelte"
    import PleiadesNav from "$lib/components/PleiadesNav.svelte"

    //TODO: create image upload capabilities with file checking

    let firstname = ""

    onMount(async () => {
        //await getInfo()
    })

    async function getInfo() {
        //TODO implement getting the name and other relevant details

        const res = await fetch(PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/projects", {
            mode: "cors",
            credentials: "include",
        })
        const data = await res.json()
        firstname = data.firstname
    }

    async function uploadProfilePic() {

    }
</script>

<PleiadesNav></PleiadesNav>
<div>
    <h2 class="scroll-m-20 pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0">
        Hello {firstname}
    </h2>
</div>

