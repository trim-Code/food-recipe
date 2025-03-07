<template>
  <div class="flex h-screen">
    <!-- Image Section -->
    <div
      class="hidden md:flex w-1/2  bg-[url('/assets/images/bg-home.jpg')] bg-cover bg-no-repeat bg-center"
    >
      <div
        class="flex flex-col items-center justify-center w-full h-full bg-black bg-opacity-50"
      >
        <h1 class="text-5xl font-bold text-white mb-4">
          Welcome to Recipe App
        </h1>
        <p class="text-lg text-white">
          Discover create and share your favorite recipes.
        </p>
      </div>
    </div>

    <!-- Authentication Form Section -->
    <div
      class="flex flex-col items-center justify-center w-full md:w-1/2 bg-white p-8"
    >
      <h2 class="text-3xl font-bold mb-6 text-[#2a3a6ddd]">Sign In</h2>
      <form class="w-full max-w-md">
        <div class="mb-4">
          <label class="block text-[#2a3a6ddd] text-sm font-bold mb-2" for="email"
            >Email</label
          >
          <input
            type="email"
            id="email"
            class="w-full px-3 py-2 border rounded-lg focus:outline-none  "
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
            class="w-full px-3 py-2 border rounded-lg focus:outline-none  relative "
            placeholder="Enter your password"
          />
<FontawesomeIcon
            icon="eye"
            class="absolute right-3 top-2 text-white cursor-pointer"
          />
        </div>
        <div class="mb-4 flex items-center justify-between">
          <label class="inline-flex items-center">
            <input type="checkbox" class="form-checkbox text-blue-500" />
            <span class="ml-2 text-gray-700">Remember me</span>
          </label>
          <a href="#" class="text-sm text-blue-500 hover:underline"
            >Forgot password?</a
          >
        </div>
        <div class="text-center">
          <button
            type="submit"
            class="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-700 focus:outline-none focus:ring focus:ring-blue-300"
          >
            Sign In
          </button>
        </div>
      </form>
      <p class="mt-6 text-gray-700">
        Don't have an account?
        <NuxtLink to="/auth/signup" class="text-blue-500 hover:underline"
          >Sign Up</NuxtLink
        >
      </p>
    </div>
  </div>
</template>

<script setup></script>

<style scoped></style>
