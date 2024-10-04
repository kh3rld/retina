<script lang="ts">
    import Sidebar from "$components/Sidebar.svelte";
    import { SERVER_URL } from "$lib/config";
    import { onMount } from "svelte";

    let content = "Loading...";

    onMount(async () => {
        const res = await fetch(`${SERVER_URL}/api/me`, {
            credentials: "include", // allows cookies to be sent to server
        });

        if (res.status === 401) {
            content = "Hello, world!";
            return;
        }

        const data = await res.json();

        console.log(data);

        content = `Hello ${data.RawData.email}!`;
    });
</script>

<div class="flex h-full">
    <Sidebar />

    <div class="flex-grow">
        <h1>Welcome to SvelteKit</h1>
        <p>
            Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the
            documentation
        </p>
        <p>{content}</p>
    </div>
</div>
