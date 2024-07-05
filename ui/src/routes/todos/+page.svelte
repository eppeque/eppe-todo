<script lang="ts">
  import Title from "$lib/components/Title.svelte";
  import { PROVIDER_CTX, StoreProvider } from "$lib/provider";
  import { checkAuth } from "$lib/redirect";
  import { getContext, onMount } from "svelte";

  const provider = getContext<StoreProvider>(PROVIDER_CTX);
  const todos = provider.todoStore;

  onMount(async () => {
    checkAuth(provider.authStore);

    const token = localStorage.getItem("auth");

    if (token !== null) {
      await todos.init(token);
    }
  });
</script>

<Title text="Your todos" />
<ul>
  {#if $todos}
    {#each $todos as todo}
      <li>{todo.title}</li>
    {/each}
  {/if}
</ul>
