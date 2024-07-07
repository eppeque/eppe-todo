<script lang="ts">
  import Form from "$lib/components/Form.svelte";
  import Title from "$lib/components/Title.svelte";
  import type { InputData } from "$lib/InputData";
  import { getContext, onMount } from "svelte";
  import { StoreProvider, PROVIDER_CTX } from "$lib/provider";
  import { checkNotAuth } from "$lib/redirect";
  import { goto } from "$app/navigation";

  const inputs: InputData[] = [
    {
      label: "Email",
      type: "email",
      name: "email",
      required: true,
    },
    {
      label: "Password",
      type: "password",
      name: "password",
      required: true,
    },
  ];

  const provider = getContext<StoreProvider>(PROVIDER_CTX);

  let errorMessage = "";

  onMount(() => {
    checkNotAuth(provider.authStore);
  });

  async function handleSubmit(e: SubmitEvent) {
    e.preventDefault();

    const entries = new FormData(e.target as HTMLFormElement);
    const data = Object.fromEntries(entries);

    const email = data.email as string;
    const password = data.password as string;

    try {
      await provider.authStore.login(email, password);
      await goto("/todos");
    } catch (e: any) {
      errorMessage = e;
    }
  }
</script>

<svelte:head>
  <title>Sign In - Eppe Todo</title>
</svelte:head>

<div class="h-[80vh] flex justify-center items-center">
  <div
    class="shadow-xl rounded-lg p-10 w-full md:w-[75%] lg:w-[60%] text-center"
  >
    <Title text="Sign In" />
    {#if errorMessage}
      <p class="py-2 text-red-600">{errorMessage}</p>
    {/if}
    <Form {inputs} submitText="Login" on:submit={handleSubmit} />
  </div>
</div>
