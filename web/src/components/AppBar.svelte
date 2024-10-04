<script lang="ts">
    import { SERVER_URL } from "$lib/config";
    import type { User } from "$lib/types/user";
    import { Icon, MagnifyingGlass } from "svelte-hero-icons";

    export let user: User | null;

    let search = "";

    const handleSignIn = () => {
        window.location.href = `${SERVER_URL}/auth/google`;
    };
</script>

<nav
    class="flex h-12 items-center border-b-2 py-8 pr-8 dark:border-violet-900 dark:bg-neutral-900"
>
    <a class="w-80 text-center text-2xl dark:text-violet-400" href="/">Retina</a
    >
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
    <div class="grow"></div>
    {#if user}
        <a class="h-10 w-10 rounded-full p-0 hover:shadow" href="/profile">
            <img
                class="rounded-full"
                width="100%"
                height="100%"
                src={user.picture}
                alt="profile_picture"
            />
        </a>
    {:else}
        <button
            class="h-10 rounded-full px-4 hover:shadow dark:bg-violet-700"
            on:click={handleSignIn}
        >
            Sign in
        </button>
    {/if}
</nav>
