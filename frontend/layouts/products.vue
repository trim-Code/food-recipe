<template>
  <div class="flex flex-col h-screen">
    <main class="grow-[2]">
      <NuxtPage />
    </main>
    <footer class="p-5 bg-[#221e1e]">
      <div class='flex items-center gap-6'>
        <div class="text-[#12b488]"><NuxtLink to="/products">Logo</NuxtLink></div>
        <nav>
          <ul
            class="list-none flex flex-col justify-center items-start gap-5 text-white no-underline"
          >
            <li><NuxtLink to="/">Home</NuxtLink></li>
            <li><NuxtLink to="/about">About</NuxtLink></li>
            <li ><NuxtLink to="/products">Products</NuxtLink></li>
          </ul>
        </nav>
      </div>
    </footer>
  </div>
</template>
<script setup></script>
