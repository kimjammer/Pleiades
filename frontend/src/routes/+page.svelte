<script>
    import { Button } from "$lib/components/ui/button/index"
    import { goto } from "$app/navigation";
</script>

<div>
    <h1 class="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl">
        Pleiades
    </h1>

    <h2
        class="scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0"
    >
        Collaborate Easier...or else...
    </h2>
    <Button onclick={() => goto('/login')}>Login</Button>
    <Button onclick={() => goto('/registration')}>Create an account</Button>
</div>
