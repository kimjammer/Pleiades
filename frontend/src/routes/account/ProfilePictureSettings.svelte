<script lang="ts">
    import * as Card from "$lib/components/ui/card"
    import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"
    import { toast } from "svelte-sonner"
    import { Label } from "$lib/components/ui/label"
    import type { ChangeEventHandler } from "svelte/elements"
    import UserAvatar from "$lib/components/UserAvatar.svelte"
    import { Input } from "$lib/components/ui/input"

    let selectedFile

    // Handle file selection and upload
    async function handleFileSelect(event: Parameters<ChangeEventHandler<HTMLInputElement>>[0]) {
        const file = event.currentTarget.files![0]
        if (!file) return

        const reader = new FileReader()
        reader.onloadend = async () => {
            let base64Image = reader.result
            if (typeof base64Image === "string") {
                //can only use .split on string
                base64Image = base64Image.split(",")[1] // Extract base64 part for backend
            }
            selectedFile = base64Image //change to int64

            //Send image to the backend
            try {
                const res = await fetch(PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/profilepic", {
                    method: "POST",
                    mode: "cors",
                    credentials: "include",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify({ image: selectedFile }),
                })
                const data = await res.json()
                if (data.success) {
                    toast.success("Profile picture uploaded successfully!")
                } else {
                    toast.error(data.error)
                }
            } catch (error) {
                toast.error("Upload failed. Please try again.")
                console.error(error)
            }
        }
        reader.readAsDataURL(file)
    }
</script>

<Card.Root class="mt-2 w-fit">
    <Card.Header>
        <Card.Title>Profile Picture</Card.Title>
    </Card.Header>
    <Card.Content>
        <div class="flex flex-col gap-2">
            <UserAvatar
                project={null}
                userID={localStorage.myId}
            />
            <Label>Upload Profile Picture</Label>
            <Input
                id="fileInput"
                type="file"
                accept="image/png, image/jpeg"
                onchange={handleFileSelect}
            />
        </div>
    </Card.Content>
</Card.Root>
