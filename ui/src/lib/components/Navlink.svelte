<script lang="ts">
  import { page } from "$app/stores";

  export let text: string;
  export let url: string;
</script>

<li>
  <a
    href={url}
    class:text-teal-600={$page.url.pathname === url}
    class="hover:underline hover:text-teal-600">{text}</a
  >
</li>
