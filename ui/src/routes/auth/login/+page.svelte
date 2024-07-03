<script lang="ts">
  import Form from "$lib/components/Form.svelte";
  import Title from "$lib/components/Title.svelte";
  import type { InputData } from "$lib/InputData";

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
</script>

<svelte:head>
  <title>Sign In - Eppe Todo</title>
</svelte:head>

<div class="h-[80vh] flex justify-center items-center">
  <div class="shadow-xl rounded-lg p-20 text-center">
    <Title text="Sign In" />
    <Form {inputs} submitText="Login" />
  </div>
</div>
