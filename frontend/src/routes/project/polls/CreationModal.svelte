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

    let { project }: { project: ProjectState } = $props()
    let createDialogOpen = $state(false)

    let question = ""
    let options = ""
    let dueDate = ""

    async function createPoll() {
        console.log("out")
        console.log(question)
        console.log(options)
        console.log(dueDate)

        const pollInfo = {
            question: question,
            options: options,
            dueDate: dueDate,
        }

        const res = await fetch(PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/newPoll", {
            method: "POST",
            mode: "cors",
            credentials: "include",
            body: JSON.stringify(pollInfo),
        })

        const data = await res.json()
        console.log(data) // Handle success or error messages
        if (data.success) {
            localStorage.myId = data.userId
            if (!(await tryJoinProject())) await goto(base + "/home")
        }
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
                <Label>Subject</Label>
                <Input bind:value={question}></Input>
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
            <Button onclick={createPoll}>Create</Button>
        </Dialog.Footer>
    </Dialog.Content>
</Dialog.Root>
