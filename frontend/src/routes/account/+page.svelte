<script>
    import { Button } from '$lib/components/ui/button'; // shadcn-svelte Button
    import { Input } from '$lib/components/ui/input'; // shadcn-svelte Input
    import { toast } from 'svelte-sonner';
    import PleiadesNav from "$lib/components/PleiadesNav.svelte";
    import {Label} from "$lib/components/ui/label";
    import {PUBLIC_API_HOST, PUBLIC_PROTOCOL} from "$env/static/public"; // Optional: for notifications
    import { onMount } from "svelte"

    let selectedFile;
    let imageUrl = ''; // Store the image URL once it is fetched from the backend

    // Handle file selection and upload
    async function handleFileSelect(event) {
        const file = event.target.files[0];
        if (!file) return;

        const reader = new FileReader();
        reader.onloadend = async () => {
            var base64Image = reader.result
            if (typeof base64Image === "string") { //can only use .split on string
                base64Image = base64Image.split(',')[1]; // Extract base64 part for backend
            }
            selectedFile = base64Image;

            // Send image to the backend
            try {
                const res = await fetch(PUBLIC_PROTOCOL +
                    PUBLIC_API_HOST + "/profilepic", {
                    method: 'POST',
                    mode: 'cors',
                    credentials: "include",
                    headers: {'Content-Type': 'application/json'},
                    body: JSON.stringify({ image: selectedFile }),
                });
                const data = await res.json()
                if (data.success) {
                    toast.success('Profile picture uploaded successfully!');
                } else {
                    toast.error(data.error);
                }
            } catch (error) {
                toast.error('Upload failed. Please try again.');
                console.error(error);
            }
        };
        reader.readAsDataURL(file);
    }



    onMount(() => {
        fetchImage();
    });

</script>
<PleiadesNav></PleiadesNav>
<div>
    <h2 class="scroll-m-20 pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0">
        Hello
    </h2>
</div>

<div class="grid w-full max-w-sm items-center gap-1.5">
    <Label>Upload Profile Picture</Label>
    <!-- Trigger file input -->
    <Input
            id="fileInput"
            type="file"
            accept="image/png, image/jpeg"
            onchange={handleFileSelect}
    />
</div>

<!-- Optionally display the uploaded image -->
{#if imageUrl}
    <img src={fetchImage()} alt="Uploaded Profile Picture" class="mt-4" />
{/if}

