<script lang="ts">
    import Sidebar from "$components/Sidebar.svelte";
    import { SERVER_URL } from "$lib/config";

    const signOut = () => {
        window.location.href = `${SERVER_URL}/logout/google`;
    };
</script>

<div class="flex h-full">
    <Sidebar>
        <button class="h-12 dark:hover:bg-violet-800" on:click={signOut}
            >Sign out</button
        >
    </Sidebar>

    <div class="flex-grow">
        <h1>Profile</h1>
    </div>
</div>
