<template>
  <div class="flex">
    <!-- Sidebar Filter -->
    <aside v-if="openFilter" class="w-1/4 h-screen bg-[#0a0a0a] pt-6 relative">
      <div class="flex flex-col h-full px-6">
        <div class="flex justify-between">
          <p class="text-xl font-[500] text-white">Filter</p>
          <button
            class="bg-[#a3b2be8f] rounded-md px-2 flex items-center justify-center"
            @click="resetFilters"
          >
            <p class="font-[500] text-white">Reset</p>
          </button>
        </div>
        <ul class="mt-4">
          <li
            v-for="category in categories"
            :key="category.id"
            @click="filterByCategory(category)"
            class="text-gray-400 hover:text-white cursor-pointer"
          >
            {{ category.name }}
          </li>
        </ul>
        <div class="mt-6">
          <label class="text-white">Preparation Time:</label>
          <div class="flex flex-col">
            <div
              v-for="(time, index) in prepTimes"
              :key="index"
              class="flex items-center"
            >
              <input
                type="checkbox"
                v-model="selectedPrepTimes"
                :value="time"
                class="mr-2"
              />
              <label class="text-gray-400">{{ time.label }}</label>
            </div>
          </div>
        </div>
        <div class="mt-4">
          <label class="text-white">Ingredients:</label>
          <input
            type="text"
            v-model="ingredientQuery"
            placeholder="Search ingredient"
            class="w-full px-2 py-1 rounded-md"
          />
        </div>
      </div>
      <button
        class="absolute top-6 -right-8 text-gray-400 cursor-pointer hover:text-white"
        @click="openFilter = false"
      >
        <font-awesome-icon icon="xmark" class="text-3xl" />
      </button>
    </aside>

    <!-- Main Content -->
    <main class="flex-1 pt-2 flex flex-col">
      <div class="flex flex-col gap-1 items-center">
        <h1 class="text-4xl font-bold text-[#bdb6a7]">Recipes</h1>
        <img src="assets/images/heading-shap.png" alt="" class="w-56" />
      </div>
      <div
        class="filter-sort flex justify-between items-center w-full mt-4 px-20 relative"
      >
        <button
          v-if="!openFilter"
          class="text-black px-5 py-2 rounded-md bg-white flex gap-3 cursor-pointer items-center"
          @click="openFilter = !openFilter"
        >
          <font-awesome-icon icon="filter" class="text-black" /> Filter
        </button>
        <div class="flex flex-col items-center gap-2 relative">
          <button
            class="text-black px-5 py-2 rounded-md bg-white flex gap-3 cursor-pointer items-center"
            @click="openSort = !openSort"
            :disabled="openFilter"
          >
            <font-awesome-icon icon="chevron-down" /> {{ currentSort }}
          </button>
          <div v-if="openSort" class="absolute top-full left-0 w-full">
            <ul class="bg-white shadow-lg rounded-md mt-2 text-black w-full">
              <li
                v-for="(sort, index) in sorts"
                :key="index"
                @click="setSort(sort)"
                class="px-4 py-2 hover:bg-gray-200 cursor-pointer"
              >
                {{ sort }}
              </li>
            </ul>
          </div>
        </div>
      </div>

      <!-- Recipe List -->
      <div class="grid grid-cols-3 gap-6 px-20 mt-6">
        <div
          v-for="recipe in filteredRecipes"
          :key="recipe.id"
          class="bg-white p-4 rounded-md shadow-lg"
        >
          <img
            :src="recipe.featuredImage"
            alt="Recipe Image"
            class="w-full h-40 object-cover rounded-md"
            v-if="recipe.featuredImage"
          />
          <h3 class="text-xl font-semibold mt-2">{{ recipe.title }}</h3>
          <p class="text-gray-600 text-sm">By {{ recipe.creator }}</p>
          <p class="text-gray-500 text-sm">{{ recipe.prepTime }} min</p>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, computed } from "vue";
const openFilter = ref(false);
const openSort = ref(false);
const currentSort = ref("All");
const ingredientQuery = ref("");
const prepTimeChecked = ref(false);
const categories = ref([
  { id: 1, name: "Breakfast" },
  { id: 2, name: "Lunch" },
  { id: 3, name: "Dinner" },
]);
const sorts = [
  "All",
  "Price:Low-to-High",
  "Price:High-to-Low",
  "Sort by Name",
  "Sort by Date",
  "Sort by Rating",
];
const recipes = ref([
  {
    id: 1,
    title: "Pancakes",
    creator: "Alice",
    prepTime: 20,
    featuredImage: "/assets/images/jj.jpg",
    category: "Breakfast",
  },
  {
    id: 2,
    title: "Spaghetti",
    creator: "Bob",
    prepTime: 30,
    featuredImage: "/assets/images/dd.jpeg",
    category: "Dinner",
  },
  {
    id: 3,
    title: " ",
    creator: "John",
    prepTime: 50,
    featuredImage: "/assets/images/bg-home.jpg",
    category: "Dinner",
  },
]);
const selectedPrepTimes = ref([]);
const prepTimes = [
  { label: "0-15 min", min: 0, max: 15 },
  { label: "16-30 min", min: 16, max: 30 },
  { label: "31-45 min", min: 31, max: 45 },
  { label: "46-60 min", min: 46, max: 60 },
  { label: "61-120 min", min: 61, max: 120 },
];
const filteredRecipes = computed(() => {
  return recipes.value
    .filter((recipe) => {
      return (
        selectedPrepTimes.value.length === 0 ||
        selectedPrepTimes.value.some(
          (time) => recipe.prepTime >= time.min && recipe.prepTime <= time.max
        )
      );
    })
    .filter(
      (recipe) =>
        ingredientQuery.value === "" ||
        recipe.ingredients?.includes(ingredientQuery.value)
    );
});
function filterByCategory(category) {
  recipes.value = recipes.value.filter(
    (recipe) => recipe.category === category.name
  );
}
function resetFilters() {
  preparationTime.value = 0;
  ingredientQuery.value = "";
}
function setSort(sort) {
  currentSort.value = sort;
  openSort.value = false;
}
</script>
