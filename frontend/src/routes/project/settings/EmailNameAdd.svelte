<script lang="ts">
    import { buttonVariants } from "$lib/components/ui/button"
    import * as Dialog from "$lib/components/ui/dialog"
    import EmailNameAddDialog from "./EmailNameAddDialog.svelte"
</script>

<Dialog.Root>
    <Dialog.Trigger class={buttonVariants({ variant: "outline" })}>
        <img
            src="purdue_logo.png"
            alt="Purdue Logo"
            class="h-3"
        />
        Name
        <EmailNameAddDialog description="Enter the name of the Purdue student you want to invite" />
    </Dialog.Trigger>
</Dialog.Root>

<Dialog.Root>
    <Dialog.Trigger class={buttonVariants({ variant: "outline" })}>
        Email
        <EmailNameAddDialog description="Enter the email of the person you want to invite" />
    </Dialog.Trigger>
</Dialog.Root>
