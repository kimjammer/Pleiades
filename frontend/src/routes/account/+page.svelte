<script lang="ts">
    import ManualInput from "$lib/availability/ManualInput.svelte";
    import {weekdayDateRanges} from "$lib/availability/timeutils";
    import TzPicker from "$lib/components/ui/TzPicker.svelte"
  
    let tzOffset = 0;
    const ranges = weekdayDateRanges();
</script>
  
<TzPicker />

<ManualInput
    {ranges}
    {tzOffset}
    shouldUseWeekdays={true}
/>
  