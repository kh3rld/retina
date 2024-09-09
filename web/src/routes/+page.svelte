<script lang="ts">
    import { SERVER_URL } from "$lib/config";
    import { onMount } from "svelte";

    let content = "Loading...";

    onMount(async () => {
        const res = await fetch(`${SERVER_URL}/api/hello`);
        const data = await res.text();

        content = data;
    });
</script>

<h1>Welcome to SvelteKit</h1>
<p>
    Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation
</p>
<p>{content}</p>
