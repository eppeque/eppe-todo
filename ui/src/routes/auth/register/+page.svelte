<script lang="ts">
  import Form from "$lib/components/Form.svelte";
  import Title from "$lib/components/Title.svelte";
  import type { InputData } from "$lib/InputData";

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

  let errorMessage = "";

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

    // Requête HTTP à envoyer...
    const res = await fetch("http://localhost:8080/api/register", {
      method: "POST",
      body: JSON.stringify({
        username,
        email,
        password,
      }),
    });

    const json = await res.json();

    if (res.ok) {
      console.log(json);
    } else {
      errorMessage = json.message;
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
