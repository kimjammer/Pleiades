<script lang="ts">
    import ManualInput from "$lib/availability/ManualInput.svelte";
    import {weekdayDateRanges} from "$lib/availability/timeutils";
  
    let tzOffset = 0;
    const ranges = weekdayDateRanges();
  </script>
  
  <ManualInput
      {ranges}
      {tzOffset}
      shouldUseWeekdays={true}
  />
  