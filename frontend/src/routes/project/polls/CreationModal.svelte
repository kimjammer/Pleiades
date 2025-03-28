<script lang="ts">
    import { buttonVariants } from "$lib/components/ui/button"
    import * as Dialog from "$lib/components/ui/dialog"
    import * as Form from "$lib/components/ui/form"
    import { Input } from "$lib/components/ui/input"
    import type { ProjectState } from "$lib/project_state.svelte"
    import { Label } from "$lib/components/ui/label";
    import { Button } from "$lib/components/ui/button";
    import { pollformSchema, type PollFormSchema } from "$lib/schema"
    import { type Infer, superForm, type SuperValidated } from "sveltekit-superforms"
    import { zodClient } from "sveltekit-superforms/adapters"
    import {PUBLIC_API_HOST, PUBLIC_PROTOCOL} from "$env/static/public";
    import {tryJoinProject} from "$lib/restApi";
    import {goto} from "$app/navigation";
    import {base} from "$app/paths";
    import { toast } from "svelte-sonner"

    let { project }: { project: ProjectState } = $props()
    let createDialogOpen = $state(false)

    let title = ""
    let description = ""
    let options = ""
    let dueDate = ""

    let error = false

    function validateInput() { //does error checking
        error = false
        if (title == "") {
            error = true
            toast.error("Title must be at least one character")
        }
        if (options.split(",").length < 2) {
            error = true
            toast.error("The poll must have at least two options")
        }
        if (dueDate == "") {
            error = true
            toast.error("The poll must have an end date")
        }
        if (!error) {
            createPoll()
        }

    }

    async function createPoll() {
        console.log("out")
        console.log(title)
        console.log(options)
        console.log(dueDate)

        //TODO: checking for title and at least two options

        const pollInfo = {
            title: title,
            description: description,
            options: options,
            dueDate: dueDate,
        }
        const url = PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/newPoll?id=" + project.id
        console.log(url)
        const res = await fetch(url, {
            method: "POST",
            mode: "cors",
            credentials: "include",
            body: JSON.stringify(pollInfo),
        })
        console.log("awaiting")
        const data = await res.json()
        if (data.success) {
            console.log("SUCCESSFULLY ADDED POLL")
        }

        //project.polls = //todo finsih this wtf

    }

    //TODO: display all poll titles as buttons to expand the poll
    //      note: look at project page

</script>

<Dialog.Root bind:open={createDialogOpen}>
    <Dialog.Trigger class={buttonVariants({ variant: "outline" })}>
        Create a new Poll
    </Dialog.Trigger>
    <Dialog.Content>
        <Dialog.Header>
            <Dialog.Title>Create new poll</Dialog.Title>
            <Dialog.Description>
                Enter poll question and options
            </Dialog.Description>
            <div>
                <Label>Title</Label>
                <Input bind:value={title}></Input>
                <Label>Description</Label>
                <Input bind:value={description}></Input>
                <Label>Options (comma separated)</Label>
                <Input bind:value={options}></Input>
                <Label>End Date</Label>
                <Input
                        type="date"
                       bind:value={dueDate}>
                </Input>
            </div>
        </Dialog.Header>
        <Dialog.Footer>
            <Button onclick={validateInput}>Create</Button>
        </Dialog.Footer>
    </Dialog.Content>
</Dialog.Root>
