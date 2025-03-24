<script>
    import { Input } from "$lib/components/ui/input/index"
    import { Button } from "$lib/components/ui/button/index"
    import { PUBLIC_API_HOST } from "$env/static/public"
    import { goto } from "$app/navigation"
    import { onMount } from "svelte"
    import PleiadesNav from "$lib/components/PleiadesNav.svelte";


    //TODO: create image upload capabilities with file checking


    onMount(async () => {
        await getInfo()
    })

    async function getInfo() {
        //TODO implement getting the name and other relevant details
    }

</script>

<PleiadesNav></PleiadesNav>
<div>
    <h2
            class="scroll-m-20 pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0"
    >
        Hello [insert name here]
    </h2>
</div>

