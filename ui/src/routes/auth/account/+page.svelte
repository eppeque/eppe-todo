<script lang="ts">
  import Card from "$lib/components/Card.svelte";
  import Title from "$lib/components/Title.svelte";
  import { PROVIDER_CTX, type StoreProvider } from "$lib/provider";
  import { checkAuth } from "$lib/redirect";
  import { getContext, onMount } from "svelte";

  const provider = getContext<StoreProvider>(PROVIDER_CTX);
  const user = provider.authStore;

  onMount(() => {
    checkAuth(user);
  });
</script>

<svelte:head>
  <title>Your account - Eppe Todo</title>
</svelte:head>

<Title text="Your account" />
<Card>
  <p><span class="font-semibold">Username: </span>{$user?.username ?? "N/A"}</p>
  <p>
    <span class="font-semibold">Email address: </span>{$user?.email ?? "N/A"}
  </p>
  <button
    class="py-2 px-4 mt-4 bg-red-500 hover:bg-red-400 text-white text-sm font-semibold shadow-sm rounded-md"
    on:click={user.signOut}>Sign Out</button
  >
</Card>
