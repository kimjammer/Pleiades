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

    let { project }: { project: ProjectState } = $props()

    let tzOffset = $state(0)
    const ranges = weekdayDateRanges()

    function save(ev: CustomEvent<Availability>) {
        const availability = dateMapToAvailability(ev.detail)
        console.log("db availability", availability)
        const myIndex = project.users.findIndex(user => user.id === localStorage.myId)

        // TODO: how to save to db?
        project.users[myIndex].availability = availability
        project.updateInProject(`users.${myIndex}.availability`, availability)
    }

    let groupAvailabilities = $derived(
        project.users.map(user => ({
            availability: availabilityToDateMap(user.availability, "2017-02-27T00:00:00.000Z", 8),
            username: user.firstName + " " + user.lastName,
        })),
    )
    $effect(() => {
        console.log("group avail", groupAvailabilities)
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
