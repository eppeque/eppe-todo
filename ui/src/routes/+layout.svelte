<script lang="ts">
  import Navbar from "$lib/components/Navbar.svelte";
  import { onMount } from "svelte";
  import "../app.css";
  import { createContext } from "$lib/provider";

  const provider = createContext();

  onMount(provider.authStore.init);
</script>

<div class="px-10 md:px-20 lg:px-40 min-h-screen bg-white font-lexend">
  <Navbar />
  <slot />
</div>
