<script lang="ts">
  import { createEventDispatcher } from "svelte";

  export let id: number;
  export let title: string;
  export let done: boolean;
  export let showDivider: boolean;

  const dispatcher = createEventDispatcher();

  $: dispatcher("checked", {
    id,
    title,
    done,
  });

  function remove() {
    dispatcher("delete", { id });
  }
</script>

<li class="flex items-center gap-4 py-2" class:border-b-2={showDivider}>
  <input type="checkbox" bind:checked={done} class="w-4" />
  <span class="text-lg" class:line-through={done}>{title}</span>

  <button
    class="material-symbols-outlined text-red-600 hover:bg-gray-300 p-2 rounded-full"
    on:click={remove}>delete</button
  >
</li>
