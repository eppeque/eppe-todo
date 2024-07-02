<script>
  import { page } from "$app/stores";
</script>

<svelte:head>
  <title>Error - Eppe Todo</title>
</svelte:head>

<main class="grid min-h-full place-items-center px-6 py-24 sm:py-32 lg:px-8">
  <div class="text-center">
    <p class="font-semibold text-teal-600">{$page.status}</p>
    <h1
      class="mt-4 text-3xl font-bold tracking-tight text-gray-900 sm:text-5xl"
    >
      {$page.error?.message}
    </h1>
    <p class="mt-6 text-base leading-7 text-gray-600">
      Go back to safety by clicking the button below.
    </p>
    <div class="mt-10 flex items-center justify-center gap-x-6">
      <a
        href="/"
        class="rounded-md bg-teal-500 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-teal-400 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
        >Go back home</a
      >
    </div>
  </div>
</main>
