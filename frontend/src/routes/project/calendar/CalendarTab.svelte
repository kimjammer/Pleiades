<script lang="ts">
    import * as Tabs from "$lib/components/ui/tabs"
    import type { ProjectState } from "$lib/project_state.svelte"
    import Calendar from "./Calendar.svelte"

    let { project }: { project: ProjectState } = $props()

    let year = $state(new Date().getFullYear())
    let month = $state(new Date().getMonth() + 1)

    function handleYearChange(e: Event) {
        const input = e.target as HTMLInputElement
        year = parseInt(input.value) || new Date().getFullYear()
    }

    function handleMonthChange(e: Event) {
        const input = e.target as HTMLInputElement
        month = parseInt(input.value) || new Date().getMonth() + 1
    }
</script>

<Tabs.Content value="calendar">
    <div class="inputs">
        <label>
            Year:
            <input
                type="number"
                value={year}
                oninput={handleYearChange}
            />
        </label>
        <label>
            Month:
            <input
                type="number"
                min="1"
                max="12"
                value={month}
                oninput={handleMonthChange}
            />
        </label>
    </div>
    <Calendar
        {month}
        {year}
    />
</Tabs.Content>
