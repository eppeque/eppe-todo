<script lang="ts">
  import Title from "$lib/components/Title.svelte";
  import { PROVIDER_CTX, StoreProvider } from "$lib/provider";
  import { checkAuth } from "$lib/redirect";
  import { getContext, onMount } from "svelte";

  const provider = getContext<StoreProvider>(PROVIDER_CTX);
  const authStore = provider.authStore;

  onMount(() => {
    checkAuth(authStore);
  });
</script>

<Title text="Your todos" />
