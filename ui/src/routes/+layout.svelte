<script>
  import Navbar from "$lib/components/Navbar.svelte";
  import "../app.css";
</script>

<div class="px-10 md:px-20 lg:px-40 min-h-screen bg-white font-lexend">
  <Navbar />
  <slot />
</div>
