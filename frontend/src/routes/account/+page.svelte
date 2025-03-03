<script lang="ts">
    import {ManualInput, weekdayDateRanges} from "availability-selector";
  
    let tzOffset = 0;
    const ranges = weekdayDateRanges();
  </script>
  
  <ManualInput
      {ranges}
      {tzOffset}
      shouldUseWeekdays={true}
  />
  