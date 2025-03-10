<template>
  <div class="bg-gray-900 min-h-screen text-white p-4 sm:p-6">
    <!-- Top Bar -->
    <div class="flex justify-between items-center border-b border-gray-700 pb-4">
      <button class="text-gray-400 hover:text-white">&larr; New recipe</button>
      <button class="bg-red-500 text-white px-4 py-2 rounded-lg">Save recipe</button>
    </div>

    <!-- Recipe Form -->
    <div class="max-w-4xl mx-auto bg-gray-800 p-6 rounded-lg mt-6 shadow-lg">
      <!-- Recipe Image & Name -->
      <div class="flex flex-col sm:flex-row items-center sm:items-start space-y-4 sm:space-y-0 sm:space-x-4">
        <div class="w-24 h-24 bg-gray-700 rounded-lg flex items-center justify-center">
          <span class="text-gray-400 text-2xl">📷</span>
        </div>
        <div class="flex-1 w-full">
          <input
            type="text"
            placeholder="Recipe name"
            class="w-full px-3 py-2 bg-gray-700 rounded text-white focus:outline-none"
          />
          <input
            type="text"
            placeholder="Original URL"
            class="w-full px-3 py-2 bg-gray-700 rounded text-white mt-2 focus:outline-none"
          />
        </div>
      </div>

      <!-- Servings & Time -->
      <div class="flex flex-wrap gap-4 mt-4">
        <div class="flex items-center space-x-2">
          <span>Servings</span>
          <button class="bg-gray-700 px-2 py-1 rounded">-</button>
          <span class="px-2">1</span>
          <button class="bg-gray-700 px-2 py-1 rounded">+</button>
        </div>
        <div class="flex space-x-2 w-full sm:w-auto">
          <select class="bg-gray-700 px-3 py-2 rounded text-white focus:outline-none w-full">
            <option>--</option>
          </select>
          <select class="bg-gray-700 px-3 py-2 rounded text-white focus:outline-none w-full">
            <option>--</option>
          </select>
        </div>
      </div>

      <!-- Ingredients & Directions (Responsive Grid) -->
      <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mt-6">
        <!-- Ingredients -->
        <div>
          <h3 class="font-bold">Ingredients</h3>
          <input
            type="text"
            placeholder="What's the first ingredient?"
            class="w-full px-3 py-2 bg-gray-700 rounded text-white mt-2 focus:outline-none"
          />
        </div>

        <!-- Directions -->
        <div>
          <h3 class="font-bold">Directions</h3>
          <div class="flex items-center space-x-2 mt-2">
            <span class="text-red-500 font-bold">1</span>
            <input
              type="text"
              placeholder="What's the first step?"
              class="w-full px-3 py-2 bg-gray-700 rounded text-white focus:outline-none"
            />
          </div>
        </div>
      </div>

      <!-- Notes -->
      <div class="mt-6">
        <h3 class="font-bold">Notes</h3>
        <textarea
          placeholder="Add a note..."
          class="w-full px-3 py-2 bg-gray-700 rounded text-white mt-2 h-24 focus:outline-none"
        ></textarea>
      </div>
    </div>
  </div>
</template>
