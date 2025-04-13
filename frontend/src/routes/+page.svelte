<script>
    import { goto } from "$app/navigation"
    import { base } from "$app/paths"
    import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"
    import PleiadesNav from "$lib/components/PleiadesNav.svelte"
    import { Button } from "$lib/components/ui/button/index"

    const res = fetch(PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/stats", {
        method: "GET",
        mode: "cors",
    }).then(res => res.json())
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
    <p class="scroll-m-20 pb-5 text-center tracking-tight">
        {#await res then stats}
            Our {stats.users ?? 0} users have logged {Math.floor((stats.seconds ?? 0) / 60 / 60)} hours
            to {stats.tasks ?? 0} tasks across {stats.projects ?? 0} projects!
        {/await}
    </p>
    <div class="flex content-center justify-center">
        <Button
            onclick={() => goto(base + (localStorage.myId ? "/login" : "/registration"))}
            size="lg"
        >
            Get Started
        </Button>
    </div>
</div>
