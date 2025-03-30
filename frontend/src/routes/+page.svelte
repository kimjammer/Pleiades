<script>
    import { goto } from "$app/navigation"
    import { base } from "$app/paths"
    import { Button } from "$lib/components/ui/button/index"
    import PleiadesNav from "$lib/components/PleiadesNav.svelte"
</script>

<PleiadesNav />

<div class="p-5">
    <div class="flex content-center justify-center pb-5 pt-20">
        <h2
            class="max-w-4xl scroll-m-20 text-center text-4xl font-semibold tracking-tight transition-colors first:mt-0 sm:text-6xl"
        >
            Bring your team together, instantly
        </h2>
    </div>

    <div class="flex content-center justify-center pb-5">
        <h3 class="max-w-lg scroll-m-20 text-center text-lg font-normal tracking-tight sm:text-2xl">
            Real-time collaborative task boards for students. Share tasks, track progress, and
            succeed together.
        </h3>
    </div>
    <div class="flex content-center justify-center">
        <Button
            onclick={() => goto(base + "/registration")}
            size="lg"
        >
            Get Started
        </Button>
    </div>
</div>
