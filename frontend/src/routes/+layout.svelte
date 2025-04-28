<script lang="ts">
    import { Toaster } from "$lib/components/ui/sonner/index"
    import { recordEvent } from "$lib/restApi"
    import { ModeWatcher } from "mode-watcher"
    import "../app.css"
    let { children } = $props()

    recordEvent("visit")
</script>

<ModeWatcher />
<Toaster richColors />
{@render children()}
