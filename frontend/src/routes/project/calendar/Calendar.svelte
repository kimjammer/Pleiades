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

    function generateCalendar(year: number, month: number) {
        // Get the first day of the month and the number of days in the month
        const firstDay = new Date(year, month - 1, 1).getDay()
        const daysInMonth = new Date(year, month, 0).getDate()

        // Initialize calendar grid
        const calendarGrid: string[][] = []
        let currentRow: string[] = []
        let currentDay = 1

        // Fill empty cells before the first day
        for (let i = 0; i < firstDay; i++) {
            currentRow.push("")
        }

        // Generate days for the month
        while (currentDay <= daysInMonth) {
            // Complete current row
            while (currentRow.length < 7 && currentDay <= daysInMonth) {
                currentRow.push(currentDay.toString())
                currentDay++
            }

            // Add row to calendar and start a new row if needed
            calendarGrid.push(currentRow)
            currentRow = []
        }

        // Fill last row with empty cells if needed
        if (currentRow.length > 0) {
            while (currentRow.length < 7) {
                currentRow.push("")
            }
            calendarGrid.push(currentRow)
        }

        return calendarGrid
    }

    $effect(() => {
        calendar = generateCalendar(year, month)
    })
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
        {#each calendar as week}
            <tr>
                {#each week as day}
                    {@const date = `${year}-${month}-${day}`}
                    <td>
                        <div>{day}</div>
                        {#each tasks.filter(task => task.due === date) as task}
                            <div>{task}</div>
                        {/each}
                    </td>
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
