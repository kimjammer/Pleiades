import type { Availability } from "$lib/components/availability/Availability"

type DbAvailability = {
    dayOfWeek: number
    startOffset: number
    endOffset: number
}

const dateToDayOfWeek = (date: string): number => new Date(date).getDay()

export const availabilityToDateMap = (
    data: DbAvailability[],
    startDate: string,
    numDays: number,
): Availability => {
    const dateMap: Availability = {}
    const start = new Date(startDate)

    for (let i = 0; i < numDays; i++) {
        const currentDate = new Date(start)
        currentDate.setDate(start.getDate() + i)
        const dateStr = currentDate.toISOString().split("T")[0]
        const dayOfWeek = currentDate.getDay()

        const slots = data
            .filter(({ dayOfWeek: d }) => d === dayOfWeek)
            .flatMap(({ startOffset, endOffset }) =>
                Array.from({ length: endOffset - startOffset + 1 }, (_, j) => startOffset + j),
            )

        dateMap[dateStr] = slots
    }

    return dateMap
}

export const dateMapToAvailability = (dateMap: Availability): DbAvailability[] => {
    const availability: DbAvailability[] = []

    Object.entries(dateMap).forEach(([date, slots]) => {
        if (slots.length > 0) {
            const dayOfWeek = dateToDayOfWeek(date)
            let startOffset = slots[0]
            let endOffset = slots[0]

            for (let i = 1; i < slots.length; i++) {
                if (slots[i] === endOffset + 1) {
                    endOffset = slots[i]
                } else {
                    availability.push({ dayOfWeek, startOffset, endOffset })
                    startOffset = slots[i]
                    endOffset = slots[i]
                }
            }

            availability.push({ dayOfWeek, startOffset, endOffset })
        }
    })

    return availability
}
