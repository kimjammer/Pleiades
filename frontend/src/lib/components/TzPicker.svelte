<script lang="ts">
    import Check from "lucide-svelte/icons/check"
    import ChevronsUpDown from "lucide-svelte/icons/chevrons-up-down"
    import { tick } from "svelte"
    import * as Command from "$lib/components/ui/command/index.js"
    import * as Popover from "$lib/components/ui/popover/index.js"
    import { Button } from "$lib/components/ui/button/index.js"
    import { cn } from "$lib/utils.js"
    import {
        getAllTzNames,
        getLocalTzName,
        getTzOffset,
    } from "$lib/components/availability/timeutils"

    const timezones = getAllTzNames().map(tz => ({
        value: getTzOffset(tz),
        label: tz,
    }))

    let open = $state<boolean>(false)
    let selectedName = $state<string>(getLocalTzName())
    let { selectedValue = $bindable() }: { selectedValue: number } = $props()

    // We want to refocus the trigger button when the user selects
    // an item from the list so users can continue navigating the
    // rest of the form with the keyboard.
    let triggerRef = $state<HTMLButtonElement | null>(null)
    function closeAndFocusTrigger() {
        open = false
        tick().then(() => {
            triggerRef?.focus()
        })
    }
</script>

<Popover.Root bind:open>
    <Popover.Trigger bind:ref={triggerRef}>
        <Button
            variant="outline"
            role="combobox"
            aria-expanded={open}
            class="w-[200px] justify-between"
        >
            {selectedName}
            <ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 opacity-50" />
        </Button>
    </Popover.Trigger>
    <Popover.Content class="h-[500px] w-[200px] p-0">
        <Command.Root>
            <Command.Input placeholder="Search timezones..." />
            <Command.Empty>No timezones found.</Command.Empty>
            <Command.Group class="overflow-y-auto">
                {#each timezones as timezone}
                    <Command.Item
                        value={timezone.label}
                        onSelect={() => {
                            selectedName = timezone.label
                            selectedValue =
                                timezones.find(f => f.label === timezone.label)?.value ?? 0
                            closeAndFocusTrigger()
                        }}
                    >
                        <Check
                            class={cn(
                                "mr-2 h-4 w-4",
                                selectedName !== timezone.label && "text-transparent",
                            )}
                        />
                        {timezone.label}
                    </Command.Item>
                {/each}
            </Command.Group>
        </Command.Root>
    </Popover.Content>
</Popover.Root>
