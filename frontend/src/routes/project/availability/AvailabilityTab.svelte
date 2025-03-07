<script lang="ts">
    import ManualInput from "$lib/components/availability/ManualInput.svelte"
    import { weekdayDateRanges } from "$lib/components/availability/timeutils"
    import TzPicker from "$lib/components/TzPicker.svelte"
    import { writable } from "svelte/store"
    import {
        loadAvailability,
        loadAvailabilityOne,
        type Availability,
        type UserAvailability,
    } from "$lib/components/availability/Availability"
    import * as Tabs from "$lib/components/ui/tabs"
    import type { ProjectState } from "$lib/project_state.svelte"
    import { Availability as DbAvailability } from "$lib/project_state.svelte.js"
    import { availabilityToDateMap, dateMapToAvailability } from "./adapter"
    import { PUBLIC_API_HOST } from "$env/static/public"

    let { project }: { project: ProjectState } = $props()

    let tzOffset = $state(0)
    const ranges = weekdayDateRanges()
    const myAvailability = $derived(
        project.users.find(user => user.id === localStorage.myId)?.availability ?? [],
    )

    async function save(ev: CustomEvent<Availability>) {
        const availability = dateMapToAvailability(ev.detail)
        console.log("db availability", availability)

        await fetch("http://" + PUBLIC_API_HOST + "/availability", {
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
        />
        <!-- availability={loadAvailabilityOne(dummy2)} -->

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
