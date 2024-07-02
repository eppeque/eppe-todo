<svelte:head>
  <title>Introducing Eppe Todo</title>
</svelte:head>

<div class="mx-auto max-w-2xl pt-16">
  <div class="hidden sm:mb-8 sm:flex sm:justify-center">
    <div
      class="relative rounded-full px-3 py-1 text-sm leading-6 text-gray-600 ring-1 ring-gray-900/10 hover:ring-gray-900/20"
    >
      Register today and start managing your day. <a
        href="/auth/register"
        class="font-semibold text-teal-600"
        ><span class="absolute inset-0" aria-hidden="true"></span>Register
        <span aria-hidden="true">&rarr;</span></a
      >
    </div>
  </div>
  <div class="text-center">
    <h1 class="text-4xl font-bold tracking-tight text-gray-900 sm:text-6xl">
      Get Organized.<br />Get things done.
    </h1>
    <p class="mt-6 text-lg leading-8 text-gray-600">
      Create and manage tasks easily and securely in a simple and user-friendly
      application.
    </p>
    <div class="mt-10 flex items-center justify-center gap-x-6">
      <a
        href="/auth/register"
        class="rounded-md bg-teal-500 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-teal-400 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
        >Get started</a
      >
      <a
        href="/auth/login"
        class="text-sm font-semibold leading-6 text-gray-900"
        >I have an account <span aria-hidden="true">→</span></a
      >
    </div>
  </div>
</div>
