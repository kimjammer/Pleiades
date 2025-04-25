<script lang="ts">
    import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"
    import { DAY, TIME_STEP } from "$lib/components/availability"
    import {
        loadAvailability,
        loadAvailabilityOne,
        type Availability,
    } from "$lib/components/availability/Availability"
    import ManualInput from "$lib/components/availability/ManualInput.svelte"
    import {
        getTodayWeek,
        minutesSinceUTCMidnight,
        weekdayDateRanges,
    } from "$lib/components/availability/timeutils"
    import TzPicker from "$lib/components/TzPicker.svelte"
    import { Button } from "$lib/components/ui/button"
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
        if (!confirm("This will overwrite your availability. Do you want to continue?")) {
            return
        }

        const tokenClient = google.accounts.oauth2.initTokenClient({
            client_id: GOOGLE_OAUTH_CLIENT_ID,
            scope: "https://www.googleapis.com/auth/calendar.events.freebusy",
            callback: fetchEvents,
        })
        tokenClient.requestAccessToken() // no popup if already authorized
    }

    function blockWithinEvent(day: number, block: number, event: any) {
        const startDate = new Date(event.start.dateTime)
        const endDate = new Date(event.end.dateTime)
        const startBlock = minutesSinceUTCMidnight(startDate) / TIME_STEP
        const endBlock = minutesSinceUTCMidnight(endDate) / TIME_STEP
        return (
            day <= endDate.getUTCDay() &&
            day >= startDate.getUTCDay() &&
            block <= endBlock &&
            block >= startBlock
        )
    }

    async function fetchEvents(tokenResponse: any) {
        const { access_token } = tokenResponse
        const thisWeek = getTodayWeek()
        const query = new URLSearchParams({
            singleEvents: "true",
            timeMin: thisWeek[0].toISOString(),
            timeMax: thisWeek[6].toISOString(),
        })
        const data = await (
            await fetch(
                "https://www.googleapis.com/calendar/v3/calendars/primary/events?" +
                    query.toString(),
                {
                    mode: "cors",
                    headers: {
                        "Authorization": "Bearer " + access_token,
                        "Content-Type": "application/json",
                    },
                },
            )
        ).json()
        const { items: events }: { items: any[] } = data

        const availability = structuredClone($state.snapshot(myAvailability))
        for (let date in availability) {
            availability[date] = []
            const targetDayOfWeek = new Date(date).getUTCDay()
            for (let blockIdx = 0; blockIdx < DAY / TIME_STEP; blockIdx++) {
                if (!events.some(event => blockWithinEvent(targetDayOfWeek, blockIdx, event)))
                    availability[date].push(blockIdx)
            }
        }
        // Iterating through all possible blocks. If this causes issues
        // 1: rewrite the entire availability component (it needs it)
        // 2: only start within range. I believe ((24804660 * MILLISECOND) % DAY / TIME_STEP) should work. Runs from 7am to 10pm
        debugger

        await fetch(PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/availability", {
            method: "POST",
            mode: "cors",
            credentials: "include",
            body: JSON.stringify(dateMapToAvailability(availability)),
        })
    }
</script>

<svelte:head>
    <script
        src="https://accounts.google.com/gsi/client"
        async
        defer
    ></script>
</svelte:head>

<Tabs.Content value="availability">
    <TzPicker bind:selectedValue={tzOffset} />
    <Button
        variant="outline"
        onclick={googleImport}>Import from Google Calendar</Button
    >

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
