<script lang="ts">
    import { Icon, MagnifyingGlass } from "svelte-hero-icons";

    let search = "";
</script>

<nav
    class="flex h-12 items-center border-b-2 py-8 dark:border-violet-900 dark:bg-neutral-900"
>
    <span class="w-80 text-center text-2xl dark:text-violet-400">Retina</span>
    <div class="relative ml-8 w-96">
        <input
            class="w-full bg-neutral-800 pl-10"
            placeholder="Search"
            bind:value={search}
        />
        <Icon
            class="absolute left-0 top-0 ml-2 w-6 dark:text-neutral-500"
            src={MagnifyingGlass}
        />
    </div>
</nav>

<div class="fixed h-full w-80 border-r-2 dark:border-neutral-900"></div>
