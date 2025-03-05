<script lang="ts">
    import ManualInput from "$lib/components/availability/ManualInput.svelte";
    import {weekdayDateRanges} from "$lib/components/availability/timeutils";
    import TzPicker from "$lib/components/TzPicker.svelte"
  
    let tzOffset = 0;
    const ranges = weekdayDateRanges();
</script>
  
<TzPicker />

<ManualInput
    {ranges}
    {tzOffset}
    shouldUseWeekdays={true}
/>
  