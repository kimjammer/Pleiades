<script lang="ts">
    import { PUBLIC_API_HOST, PUBLIC_PROTOCOL } from "$env/static/public"
    import PleiadesNav from "$lib/components/PleiadesNav.svelte"
    import NotificationSettings from "./NotificationSettings.svelte"
    import PersonalCalendar from "./PersonalCalendar.svelte"
    import ProfilePictureSettings from "./ProfilePictureSettings.svelte"

    const userID = localStorage.getItem("myId")

    const name = fetch(PUBLIC_PROTOCOL + PUBLIC_API_HOST + "/fetchName?id=" + userID, {
        method: "GET",
        mode: "cors",
        credentials: "include",
        headers: { "Content-Type": "application/json" },
    })
        .then(res => res.json())
        .then(({ firstName, lastName }) => firstName + " " + lastName)
</script>

<PleiadesNav></PleiadesNav>
<div class="p-5">
    <h2 class="scroll-m-20 pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0">
        Settings
    </h2>

    <p>
        {#await name then name}
            Hello {name}!
        {/await}
    </p>

    <div class="mb-10 flex gap-5">
        <ProfilePictureSettings />
        <NotificationSettings />
    </div>
    <PersonalCalendar />
</div>
