<script lang="ts">
    import ManualInput from "$lib/components/availability/ManualInput.svelte"
    import { weekdayDateRanges } from "$lib/components/availability/timeutils"
    import TzPicker from "$lib/components/TzPicker.svelte"
    import { writable } from "svelte/store"

    let tzOffset = 0
    const ranges = weekdayDateRanges()
</script>

<TzPicker bind:selectedValue={tzOffset} />

<div class="flex justify-around">
    <!-- TODO: make tzOffset=-tzOffset consistent -->
    <ManualInput
        {ranges}
        tzOffset={-tzOffset}
        shouldUseWeekdays={true}
    />

    <!-- TODO: make `availablePeople` boolean, writable store value unused OR some way to delegate tooltip -->
    <ManualInput
        {ranges}
        tzOffset={-tzOffset}
        shouldUseWeekdays={true}
        availablePeople={writable([])}
    />
</div>
