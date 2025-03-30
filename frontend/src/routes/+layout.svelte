<script lang="ts">
    import { Toaster } from "$lib/components/ui/sonner/index"
    import { ModeWatcher } from "mode-watcher"
    import "../app.css"
</script>

<ModeWatcher />
<Toaster richColors />
<slot></slot>
