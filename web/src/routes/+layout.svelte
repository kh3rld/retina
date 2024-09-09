<script lang="ts">
    import AppBar from "$lib/AppBar.svelte";
    import { onMount } from "svelte";
    import "../app.css";

    // tailwindcss dark mode
    onMount(() => {
        if (
            localStorage.theme === "dark" ||
            (!("theme" in localStorage) &&
                window.matchMedia("(prefers-color-scheme: dark)").matches)
        ) {
            document.documentElement.classList.add("dark");
        } else {
            document.documentElement.classList.remove("dark");
        }
    });
</script>

<AppBar />

<div class="ml-80">
    <slot />
</div>
