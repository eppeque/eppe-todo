<script lang="ts">
  import Navbar from "$lib/components/Navbar.svelte";
  import { onMount } from "svelte";
  import "../app.css";
  import { createContext } from "$lib/provider";
  import { get } from "svelte/store";
  import { page } from "$app/stores";
  import { goto } from "$app/navigation";

  const provider = createContext();

  onMount(async () => {
    await provider.authStore.init();

    provider.authStore.subscribe((user) => {
      if (user) {
        provider.todoStore.init();
      }

      const url = get(page).url.pathname;

      if (
        (url === "/" || url === "/auth/login" || url === "/auth/register") &&
        user
      ) {
        goto("/todos");
        return;
      }

      if ((url === "/todos" || url === "/auth/account") && user) {
        goto("/");
      }
    });
  });
</script>

<div class="px-10 md:px-20 lg:px-40 min-h-screen bg-white font-lexend">
  <Navbar />
  <slot />
</div>
