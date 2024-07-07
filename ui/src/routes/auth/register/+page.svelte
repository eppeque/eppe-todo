<script lang="ts">
  import { goto } from "$app/navigation";
  import Form from "$lib/components/Form.svelte";
  import Title from "$lib/components/Title.svelte";
  import type { InputData } from "$lib/InputData";
  import { PROVIDER_CTX, StoreProvider } from "$lib/provider";
  import { checkNotAuth } from "$lib/redirect";
  import { getContext, onMount } from "svelte";

  const inputs: InputData[] = [
    {
      label: "Username",
      type: "text",
      name: "username",
      required: true,
    },
    {
      label: "Email address",
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
    {
      label: "Password confirmation",
      type: "password",
      name: "passwordConfirm",
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
    errorMessage = "";

    const entries = new FormData(e.target as HTMLFormElement);
    const data = Object.fromEntries(entries);

    const username = data.username as string;
    const email = data.email as string;
    const password = data.password as string;
    const passwordConfirm = data.passwordConfirm as string;

    if (password !== passwordConfirm) {
      errorMessage = "The password is different than its confirmation";
      return;
    }

    try {
      await provider.authStore.register(username, email, password);
      await goto("/todos");
    } catch (e: any) {
      errorMessage = e;
    }
  }
</script>

<svelte:head>
  <title>Sign Up - Eppe Todo</title>
</svelte:head>

<Title text="Sign Up" />
{#if errorMessage}
  <p class="py-2 text-red-600 text-lg">{errorMessage}</p>
{/if}
<Form {inputs} submitText="Register" on:submit={handleSubmit} />
