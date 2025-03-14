<template>
  <div class="bg-gray-900 text-white p-6 grow">
    <!-- Header -->
    <header class="flex items-center mb-6">
      <font-awesome-icon
        icon="arrow-left"
        class="cursor-pointer text-[#bdb6a7] size-8"
        @click="$router.go(-1)"
      ></font-awesome-icon>
      <div class="flex-grow text-center">
        <h1
          class="flex gap-4 items-center justify-center text-4xl text-[#bdb6a7] font-bold mb-6"
        >
          <img src="assets/images/left-shap.png" alt="" class="w-28" /> Kimem
          Recipe
          <img src="assets/images/right-shap.png" alt="" class="w-28" />
        </h1>
      </div>
    </header>

    <!-- Cookbook Section -->
    <section
      class="mt-10 max-w-3xl mx-auto bg-gray-800 p-6 rounded-lg shadow-lg"
    >
      <h2 class="text-2xl font-bold">Your Recipes</h2>
      <div class="flex space-x-4 mt-3">
        <button class="text-gray-400 hover:text-white">All recipes</button>
        <button class="text-white font-bold underline">Categories</button>
      </div>

      <!-- Search Bar -->
      <div class="relative mt-4">
        <input
          type="text"
          placeholder="Search..."
          class="w-full px-4 py-2 rounded bg-gray-700 text-white border border-gray-600 focus:outline-none"
        />
        <span class="absolute right-3 top-2.5 text-gray-400">🔍</span>
      </div>

      <!-- Empty State -->
      <div class="flex flex-col items-center mt-6">
        <img src="assets/images/2.png" alt=" " class="w-24 mb-3" />
        <p class="text-gray-400">Categories you create will appear here.</p>
      </div>
    </section>

    <!-- Recipe Form Section -->
    <section
      class="mt-10 max-w-3xl mx-auto bg-gray-800 p-6 rounded-lg shadow-lg"
    >
      <h2 class="text-2xl font-bold mb-6">Create a New Recipe</h2>
      <form class="space-y-4" @submit.prevent="handleSubmit">
        <div>
          <label class="block text-gray-400 mb-2" for="name">Recipe Name</label>
          <input
            type="text"
            id="name"
            class="w-full px-4 py-2 rounded bg-gray-700 text-white border border-gray-600 focus:outline-none"
            placeholder="Enter recipe name"
          />
        </div>
        <div>
          <label class="block text-gray-400 mb-2" for="image"
            >Recipe Image</label
          >
          <input
            type="file"
            id="image"
            @change="previewImage"
            class="w-full px-4 py-2 rounded bg-gray-700 text-white border border-gray-600 focus:outline-none"
          />
          <div v-if="imageUrl" class="mt-4">
            <img
              :src="imageUrl"
              alt="Image Preview"
              class="w-28 object-cover rounded-xl"
            />
          </div>
        </div>
        <div>
          <label class="block text-gray-400 mb-2" for="description"
            >Description</label
          >
          <textarea
            id="description"
            class="w-full px-4 py-2 rounded bg-gray-700 text-white border border-gray-600 focus:outline-none"
            placeholder="Enter recipe description"
          ></textarea>
        </div>
        <div>
          <label class="block text-gray-400 mb-2" for="ingredients"
            >Ingredients</label
          >
          <textarea
            id="ingredients"
            class="w-full px-4 py-2 rounded bg-gray-700 text-white border border-gray-600 focus:outline-none"
            placeholder="Enter ingredients, separated by commas"
          ></textarea>
        </div>
        <div>
          <label class="block text-gray-400 mb-2" for="steps">Steps</label>
          <textarea
            id="steps"
            class="w-full px-4 py-2 rounded bg-gray-700 text-white border border-gray-600 focus:outline-none"
            placeholder="Enter steps, separated by commas"
          ></textarea>
        </div>
        <div>
          <label class="block text-gray-400 mb-2" for="price">Price</label>
          <input
            type="number"
            id="price"
            class="w-full px-4 py-2 rounded bg-gray-700 text-white border border-gray-600 focus:outline-none"
            placeholder="Enter price"
          />
        </div>
        <div class="text-center">
          <button
            type="submit"
            class="bg-red-500 text-white px-4 py-2 rounded-lg hover:bg-red-700 focus:outline-none cursor-pointer"
          >
            Add Recipe
          </button>
        </div>
      </form>
    </section>

    <!-- Floating Action Button -->
    <div class="fixed bottom-6 right-6">
      <button
        @click="toggleMenu"
        class="bg-red-500 w-12 h-12 rounded-full flex items-center justify-center shadow-lg cursor-pointer"
      >
        ➕
      </button>
      <div v-if="menuOpen" class="bg-gray-800 rounded-lg shadow-lg p-3 mt-2">
        <NuxtLink to="/create-recipes/create-new-recipe">
          <button
            class="block text-white py-1 hover:bg-gray-700 px-2 rounded cursor-pointer"
          >
            ✏️ New recipe from scratch
          </button>
        </NuxtLink>
        <button
          class="block text-white py-1 hover:bg-gray-700 px-2 rounded cursor-pointer"
        >
          📁 New Categories
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
definePageMeta({
  layout: "create-recipes",
});
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { ref } from "vue";

const menuOpen = ref(false);
const imageUrl = ref(null);

const toggleMenu = () => {
  menuOpen.value = !menuOpen.value;
};

const previewImage = (event) => {
  const file = event.target.files[0];
  if (file) {
    imageUrl.value = URL.createObjectURL(file);
  } else {
    imageUrl.value = null;
  }
};

const handleSubmit = (event) => {
  event.preventDefault(); 
};
</script>
