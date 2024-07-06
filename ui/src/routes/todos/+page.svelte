<script lang="ts">
  import Title from "$lib/components/Title.svelte";
  import Todo from "$lib/components/Todo.svelte";
  import { PROVIDER_CTX, StoreProvider } from "$lib/provider";
  import { checkAuth } from "$lib/redirect";
  import { getContext, onMount } from "svelte";

  const provider = getContext<StoreProvider>(PROVIDER_CTX);
  const todos = provider.todoStore;

  let title = "";

  onMount(() => checkAuth(provider.authStore));

  async function onKeyDown(e: KeyboardEvent) {
    if (e.key !== "Enter") return;

    await todos.add(title);

    title = "";
  }

  function onChecked(e: CustomEvent) {
    const id = e.detail.id;
    const title = e.detail.title;
    const done = e.detail.done;

    todos.update(id, title, done);
  }
</script>

<svelte:head>
  <title>Your todos - Eppe Todo</title>
</svelte:head>

<Title text="Your todos" />
<input
  type="text"
  bind:value={title}
  on:keydown={onKeyDown}
  placeholder="Todo title"
  class="block w-[75%] mx-auto my-4 p-3 outline-none border-2 border-gray-200 rounded-md focus:border-teal-500"
/>
<ul class="my-4 w-[75%] mx-auto bg-gray-100 p-10 rounded-lg">
  {#each $todos as todo (todo.id)}
    <Todo {...todo} on:checked={onChecked} />
  {/each}
</ul>
