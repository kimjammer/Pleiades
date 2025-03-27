<script lang="ts">
    import * as Tabs from "$lib/components/ui/tabs"
    import type { ProjectState } from "$lib/project_state.svelte.ts"
    import { Button } from "$lib/components/ui/button";
    import { Input } from "$lib/components/ui/input";
    import { Textarea } from "$lib/components/ui/textarea";
    import { Label } from "$lib/components/ui/label";
    import { Card, CardContent, CardHeader, CardTitle } from "$lib/components/ui/card";
    import type { PageData } from "./$types"
    import CreationModal from "./CreationModal.svelte"

    let { project, data }: { project: ProjectState; data: PageData } = $props()


    async function createPoll() {
        //TODO: send info to create poll
    }

</script>

<Tabs.Content value="polls">
    <div>above</div>
    <CreationModal
            {project}
            {data}
    />

    <div>below</div>

</Tabs.Content>

