<script lang="ts">
  import ManualInput from "$lib/ManualInput.svelte";
  import { weekdayDateRanges } from "$lib/timeutils.js"

  let tzOffset = 0;
  const ranges = weekdayDateRanges();
</script>

<ManualInput
    {ranges}
    {tzOffset}
    shouldUseWeekdays={true}
/>
