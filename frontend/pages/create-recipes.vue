<template>
  <div class="container mx-auto px-4 py-12">
    <h1 class="text-4xl font-bold text-center mb-8">Create a New Recipe</h1>
    <form class="max-w-2xl mx-auto bg-white p-8 rounded-lg shadow-lg">
      <div class="mb-4">
        <label class="block text-gray-700 text-sm font-bold mb-2" for="name"
          >Recipe Name</label
        >
        <input
          type="text"
          id="name"
          class="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring focus:ring-blue-300"
          placeholder="Enter recipe name"
        />
      </div>
      <div class="mb-4">
        <label class="block text-gray-700 text-sm font-bold mb-2" for="image"
          >Recipe Image</label
        >
        <input
          type="file"
          id="image"
          class="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring focus:ring-blue-300"
        />
      </div>
      <div class="mb-4">
        <label
          class="block text-gray-700 text-sm font-bold mb-2"
          for="description"
          >Description</label
        >
        <textarea
          id="description"
          class="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring focus:ring-blue-300"
          placeholder="Enter recipe description"
        ></textarea>
      </div>
      <div class="mb-4">
        <label
          class="block text-gray-700 text-sm font-bold mb-2"
          for="ingredients"
          >Ingredients</label
        >
        <textarea
          id="ingredients"
          class="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring focus:ring-blue-300"
          placeholder="Enter ingredients, separated by commas"
        ></textarea>
      </div>
      <div class="mb-4">
        <label class="block text-gray-700 text-sm font-bold mb-2" for="steps"
          >Steps</label
        >
        <textarea
          id="steps"
          class="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring focus:ring-blue-300"
          placeholder="Enter steps, separated by commas"
        ></textarea>
      </div>
      <div class="mb-4">
        <label class="block text-gray-700 text-sm font-bold mb-2" for="other"
          >Other Details</label
        >
        <textarea
          id="other"
          class="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring focus:ring-blue-300"
          placeholder="Enter any other details"
        ></textarea>
      </div>
      <div class="text-center">
        <button
          type="submit"
          class="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-700 focus:outline-none focus:ring focus:ring-blue-300"
        >
          Add Recipe
        </button>
      </div>
    </form>
  </div>
</template>

<script setup></script>

<style scoped>
.container {
  background: linear-gradient(to right, #f8fafc, #e2e8f0);
}
</style>
