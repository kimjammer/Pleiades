<script lang="ts">
    import PleiadesNav from "$lib/components/PleiadesNav.svelte"
    import PersonalCalendar from "./PersonalCalendar.svelte"
    import NotificationSettings from "./NotificationSettings.svelte"
    import ProfilePictureSettings from "./ProfilePictureSettings.svelte"
</script>

<PleiadesNav></PleiadesNav>
<div class="p-5">
    <h2 class="scroll-m-20 pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0">
        Settings
    </h2>

    <div class="mb-10 flex gap-5">
        <ProfilePictureSettings />
        <NotificationSettings />
    </div>
    <PersonalCalendar />
</div>
