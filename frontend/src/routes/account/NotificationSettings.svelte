<script lang="ts">
    import { Switch } from "$lib/components/ui/switch"
    import * as Card from "$lib/components/ui/card"
    import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"
    import { toast } from "svelte-sonner"
    import { Label } from "$lib/components/ui/label"
    import { onMount } from "svelte"

    let notifUserJoin = $state(false)
    let notifPollEnd = $state(false)
    let notifTaskAssign = $state(false)

    onMount(async () => {
        await getNotifSettings()
    })

    async function flipNotifSetting(index: number) {
        console.log("Flipping notification setting")
        const res = await fetch(PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/flipNotif", {
            method: "POST",
            mode: "cors",
            credentials: "include",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ notifIndex: index }),
        })
        const data = await res.json()
        if (data.success) {
            toast.success("Changed notification preference")
            await getNotifSettings()
        } else {
            toast.error("Error changing notification preference")
        }
    }

    async function getNotifSettings() {
        const res = await fetch(PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/notifSettings", {
            method: "GET",
            mode: "cors",
            credentials: "include",
            headers: { "Content-Type": "application/json" },
        })
        const data = await res.json()
        if (data.success) {
            ;[notifUserJoin, notifPollEnd, notifTaskAssign] = data.notifSettings
            console.log(notifUserJoin, notifPollEnd, notifTaskAssign)
        } else {
            toast.error("Failed to load notification settings")
        }
    }
</script>

<Card.Root class="mt-2 w-fit">
    <Card.Header>
        <Card.Title>Notification Settings</Card.Title>
    </Card.Header>
    <Card.Content class="grid gap-6">
        <div class="flex items-center space-x-2">
            <Switch
                id="user-joining"
                bind:checked={notifUserJoin}
                onCheckedChange={() => flipNotifSetting(0)}
            />
            <Label for="user-joining">New users joining projects</Label>
        </div>
        <div class="flex items-center space-x-2">
            <Switch
                id="ending-polls"
                bind:checked={notifPollEnd}
                onCheckedChange={() => flipNotifSetting(1)}
            />
            <Label for="ending-polls">Polls ending soon</Label>
        </div>
        <div class="flex items-center space-x-2">
            <Switch
                id="task-assignments"
                bind:checked={notifTaskAssign}
                onCheckedChange={() => flipNotifSetting(2)}
            />
            <Label for="task-assignments">New task assignments</Label>
        </div>
    </Card.Content>
</Card.Root>
