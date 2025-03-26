<script lang="ts">
    import type { Task } from "$lib/schema"

    let { year, month, tasks = [] }: { year: number; month: number; tasks?: Task[] } = $props()
    let calendar = $state<string[][]>([])

    const daysOfWeek = [
        "Sunday",
        "Monday",
        "Tuesday",
        "Wednesday",
        "Thursday",
        "Friday",
        "Saturday",
    ]

    const getDaysInMonth = (year: number, month: number) => new Date(year, month, 0).getDate()

    /**
     * Adapted from https://stackoverflow.com/a/2485172
     * @param month in range 1-12
     */
    function weekCount(year: number, month: number) {
        var firstOfMonth = new Date(year, month - 1, 1)
        var lastOfMonth = new Date(year, month, 0)

        var used = firstOfMonth.getDay() + lastOfMonth.getDate()

        return Math.ceil(used / 7)
    }

    function firstDayOfWeek(month: number, year: number, day: number) {
        const date = new Date(year, month - 1, day)
        return Math.max(1, day - date.getDay())
    }
</script>

<table>
    <thead>
        <tr>
            {#each daysOfWeek as day}
                <th>{day}</th>
            {/each}
        </tr>
    </thead>
    <tbody>
        {#each { length: weekCount(year, month) }, weekIndex}
            {@const representativeDate = weekIndex * 7 + 1}
            {@const sundayDate = firstDayOfWeek(month, year, representativeDate)}
            <tr>
                {#each { length: Math.min(getDaysInMonth(year, month) - sundayDate + 1, 7) }, dayIndex}
                    <td>{sundayDate + dayIndex}</td>
                {/each}
            </tr>
        {/each}
    </tbody>
</table>

<style>
    table {
        width: 100vw;
        height: 100vh;
        border-collapse: collapse;
    }

    th,
    td {
        border: 1px solid black;
        padding: 5px;
        width: calc(100vw / 7);
    }

    th {
        height: 5vh; /* Header takes about 5% of vertical space */
        text-align: right;
    }

    td {
        vertical-align: top;
    }
</style>
