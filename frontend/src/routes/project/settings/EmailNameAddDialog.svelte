<script lang="ts">
    import * as Command from "$lib/components/ui/command"
    import * as Dialog from "$lib/components/ui/dialog"

    let { description = "" }: { description?: string } = $props()
    let nameOrEmail = $state("")
</script>

<Dialog.Content class="sm:max-w-[425px]">
    <Dialog.Header>
        <Dialog.Title>Invite Teammates</Dialog.Title>
        <Dialog.Description>
            {description}
        </Dialog.Description>
    </Dialog.Header>

    <Command.Root>
        <Command.Input
            placeholder="Name or email"
            bind:value={nameOrEmail}
        />
        <Command.List>
            <Command.Empty>No results found.</Command.Empty>
            <Command.Item>Calendar</Command.Item>
            <Command.Item>Search Emoji</Command.Item>
            <Command.Item>Calculator</Command.Item>
        </Command.List>
    </Command.Root>

    <Dialog.Footer>
        <!-- TODO: s<Button type="submit">Done</Button> -->
    </Dialog.Footer>
</Dialog.Content>
