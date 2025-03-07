<template>
  <div class="flex h-screen">
    <!-- Image Section -->
    <div
      class="hidden md:flex w-1/2 bg-[url('/assets/images/jj.jpg')] bg-cover bg-no-repeat bg-center"
    >
      <div
        class="flex flex-col items-center justify-center w-full h-full bg-black bg-opacity-60"
      >
        <h1 class="text-5xl font-bold text[#d2c8aa] mb-4">
          Welcome to Recipe App
        </h1>
        <p class="text-lg text[#d2c8aa]">
          Discover create and share your favorite recipes.
        </p>
      </div>
    </div>

    <!-- Signup Form Section -->
    <div
      class="flex flex-col items-center justify-center w-full md:w-1/2 bg-[white] p-8"
    >
      <h2 class="text-3xl font-bold mb-6 text-[#2a3a6ddd]">Sign Up</h2>
      <form class="w-full max-w-md">
        <div class="mb-4">
          <label class="block text-[#2a3a6ddd] text-sm font-bold mb-2" for="name"
            >Name</label
          >
          <input
            type="text"
            id="name"
            class="w-full px-3 py-2 border rounded-lg focus:outline-none"
            placeholder="Enter your name"
          />
        </div>
        <div class="mb-4">
          <label class="block text-[#2a3a6ddd] text-sm font-bold mb-2" for="email"
            >Email</label
          >
          <input
            type="email"
            id="email"
            class="w-full px-3 py-2 border rounded-lg focus:outline-none"
            placeholder="Enter your email"
          />
        </div>
        <div class="mb-4">
          <label
            class="block text-[#2a3a6ddd] text-sm font-bold mb-2"
            for="password"
            >Password</label
          >
          <input
            type="password"
            id="password"
            class="w-full px-3 py-2 border rounded-lg focus:outline-none"
            placeholder="Enter your password"
          />
        </div>
        <div class="mb-4">
          <label
            class="block text-[#2a3a6ddd] text-sm font-bold mb-2"
            for="confirm-password"
            >Confirm Password</label
          >
          <input
            type="password"
            id="confirm-password"
            class="w-full px-3 py-2 border rounded-lg focus:outline-none"
            placeholder="Confirm your password"
          />
        </div>
        <div class="text-center">
          <button
            type="submit"
            class="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-[#2a3a6ddd] focus:outline-none"
          >
            Sign Up
          </button>
        </div>
      </form>
      <p class="mt-6 text-gray-700">
        Already have an account?
        <NuxtLink to="/auth/login" class="text-blue-500 hover:underline"
          >Sign In</NuxtLink
        >
      </p>
    </div>
  </div>
</template>

<script setup></script>

<style scoped></style>
