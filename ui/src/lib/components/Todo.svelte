<script lang="ts">
  import { createEventDispatcher } from "svelte";

  export let id: number;
  export let title: string;
  export let done: boolean;

  const dispatcher = createEventDispatcher();

  $: dispatcher("checked", {
    id,
    title,
    done,
  });
</script>

<li class="flex gap-4 py-2">
  <input type="checkbox" bind:checked={done} class="w-4" />
  <span class="text-lg" class:line-through={done}>{title}</span>
</li>
