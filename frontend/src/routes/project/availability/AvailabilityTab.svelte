<script lang="ts">
    import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"
    import {
        loadAvailability,
        loadAvailabilityOne,
        type Availability,
    } from "$lib/components/availability/Availability"
    import ManualInput from "$lib/components/availability/ManualInput.svelte"
    import { weekdayDateRanges } from "$lib/components/availability/timeutils"
    import TzPicker from "$lib/components/TzPicker.svelte"
    import * as Tabs from "$lib/components/ui/tabs"
    import type { ProjectState } from "$lib/project_state.svelte"
    import { GOOGLE_OAUTH_CLIENT_ID } from "$lib/utils"
    import { writable } from "svelte/store"
    import { availabilityToDateMap, dateMapToAvailability } from "./adapter"

    let { project }: { project: ProjectState } = $props()

    let tzOffset = $state(0)
    const ranges = weekdayDateRanges()
    const myAvailability = $derived(
        availabilityToDateMap(
            project.users.find(user => user.id === localStorage.myId)?.availability ?? [],
            "2017-02-27T00:00:00.000Z",
            8,
        ),
    )

    async function save(ev: CustomEvent<Availability>) {
        const availability = dateMapToAvailability(ev.detail)
        console.log("db availability", availability)

        await fetch(PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/availability", {
            method: "POST",
            mode: "cors",
            credentials: "include",
            body: JSON.stringify(availability),
        })
    }

    let groupAvailabilities = $derived(
        project.users.map(user => ({
            availability: availabilityToDateMap(user.availability, "2017-02-27T00:00:00.000Z", 8),
            username: user.firstName + " " + user.lastName,
        })),
    )
    $effect(() => {
        console.log("group avail", $state.snapshot(groupAvailabilities))
    })
    $effect(() => {
        console.log("my avail", $state.snapshot(myAvailability))
    })

    function googleImport() {
        const tokenClient = google.accounts.oauth2.initTokenClient({
            client_id: GOOGLE_OAUTH_CLIENT_ID,
            scope: "https://www.googleapis.com/auth/calendar.events.freebusy",
            callback: fetchEvents,
        })
        tokenClient.requestAccessToken() // no popup if already authorized
    }

    async function fetchEvents(tokenResponse: any) {
        const { accessToken } = tokenResponse
        const res = await (
            await fetch("https://www.googleapis.com/calendar/v3/calendars/primary/events", {
                mode: "cors",
                headers: {
                    "Authorization": `Bearer ${accessToken}`,
                    "Content-Type": "application/json",
                },
            })
        ).json()
    }
</script>

<Tabs.Content value="availability">
    <TzPicker bind:selectedValue={tzOffset} />

    <div class="flex justify-around">
        <!-- TODO: make tzOffset=-tzOffset consistent -->
        <ManualInput
            {ranges}
            tzOffset={-tzOffset}
            shouldUseWeekdays={true}
            on:save={save}
            availability={loadAvailabilityOne(myAvailability)}
        />

        <!-- TODO: make `availablePeople` boolean, writable store value unused OR some way to delegate tooltip -->
        <ManualInput
            {ranges}
            tzOffset={-tzOffset}
            shouldUseWeekdays={true}
            availablePeople={writable([])}
            allParticipants={groupAvailabilities.map(person => person.username)}
            availability={loadAvailability(...groupAvailabilities)}
        />
    </div>
</Tabs.Content>
