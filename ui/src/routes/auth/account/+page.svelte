<script lang="ts">
  import { goto } from "$app/navigation";
  import Card from "$lib/components/Card.svelte";
  import Title from "$lib/components/Title.svelte";
  import { PROVIDER_CTX, type StoreProvider } from "$lib/provider";
  import { getContext } from "svelte";

  const provider = getContext<StoreProvider>(PROVIDER_CTX);
  const authState = provider.authStore;

  function signOut() {
    authState.signOut();
    goto("/");
  }
</script>

<svelte:head>
  <title>Your account - Eppe Todo</title>
</svelte:head>

<Title text="Your account" />
<Card>
  <p>
    <span class="font-semibold">Username: </span>{$authState.user?.username ??
      "N/A"}
  </p>
  <p>
    <span class="font-semibold">Email address: </span>{$authState.user?.email ??
      "N/A"}
  </p>
  <button
    class="py-2 px-4 mt-4 bg-red-500 hover:bg-red-400 text-white text-sm font-semibold shadow-sm rounded-md"
    on:click={signOut}>Sign Out</button
  >
</Card>
