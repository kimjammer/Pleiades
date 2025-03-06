<script lang="ts">
    import ManualInput from "$lib/components/availability/ManualInput.svelte"
    import { weekdayDateRanges } from "$lib/components/availability/timeutils"
    import TzPicker from "$lib/components/TzPicker.svelte"
    import { writable } from "svelte/store"
    import { dummy1, dummy2, dummy3 } from "./dummy"
    import {
        loadAvailability,
        loadAvailabilityOne,
        type Availability,
        type UserAvailability,
    } from "$lib/components/availability/Availability"
    import * as Tabs from "$lib/components/ui/tabs"
    import type { ProjectState } from "$lib/project_state.svelte"
    import { Availability as DbAvailability } from "$lib/project_state.svelte.js"

    let { project }: { project: ProjectState } = $props()

    let tzOffset = $state(0)
    const ranges = weekdayDateRanges()

    function save(ev: CustomEvent<Availability>) {
        const detail = ev.detail
        const myIndex = project.users.findIndex(user => user.id === localStorage.myId)
        console.log("TODO: websocket", detail)
        project.updateInProject(`users.${myIndex}.availability`, [
            {
                dayOfWeek: 0,
                startOffset: 10,
                endOffset: 11,
            },
        ] satisfies DbAvailability[])
        project.appendInProject(`users.${myIndex}.availability`, {
            dayOfWeek: 0,
            startOffset: 10,
            endOffset: 11,
        } satisfies DbAvailability)
        project.users[myIndex].firstName = "joe"
        project.updateInProject("users.0.firstName", "joe")
    }

    let groupAvailabilities: UserAvailability[] = [
        { availability: dummy1, username: "Mr. Rectangle" },
        { availability: dummy2, username: ":)" },
        { availability: dummy3, username: "unemployed" },
    ]
</script>

<Tabs.Content value="availability">
    <TzPicker bind:selectedValue={tzOffset} />

    <div class="flex justify-around">
        <!-- TODO: make tzOffset=-tzOffset consistent -->
        <ManualInput
            {ranges}
            tzOffset={-tzOffset}
            shouldUseWeekdays={true}
            availability={loadAvailabilityOne(dummy2)}
            on:save={save}
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
