<script>
  import Navbar from "$lib/components/Navbar.svelte";
  import { onMount } from "svelte";
  import "../app.css";
  import { createContext } from "$lib/provider";
  import { page } from "$app/stores";
  import { get } from "svelte/store";
  import { goto } from "$app/navigation";

  const provider = createContext();

  // TODO: Implement login + use local storage + refresh token

  onMount(() => {
    return provider.authStore.subscribe((user) => {
      const url = get(page).url.pathname;

      console.log(url);

      if (url === "/todos" && user === null) {
        goto("/");
      } else if (
        (url === "/" || url == "/auth/register" || url === "/auth/login") &&
        user !== null
      ) {
        goto("/todos");
      }
    });
  });
</script>

<div class="px-10 md:px-20 lg:px-40 min-h-screen bg-white font-lexend">
  <Navbar />
  <slot />
</div>
