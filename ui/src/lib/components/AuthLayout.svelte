<script lang="ts">
  import { goto } from "$app/navigation";
  import { PROVIDER_CTX, StoreProvider } from "$lib/provider";
  import { getContext, onMount } from "svelte";

  const provider = getContext<StoreProvider>(PROVIDER_CTX);
  const authStore = provider.authStore;

  let ok = false;

  onMount(() => {
    return authStore.subscribe((state) => {
      if (!state.initializing) {
        ok = state.user !== null;

        if (!ok) {
          goto("/auth/login");
        } else {
          provider.todoStore.init();
        }
      }
    });
  });
</script>

{#if ok}
  <slot />
{:else}
  <div></div>
{/if}
