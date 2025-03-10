<template>
  <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6 p-4">
    <div
          v-if="filteredRecipes.length"
          class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-8"
        >
          <div
            v-for="recipe in filteredRecipes"
            :key="recipe.id"
            class="relative bg-white rounded-xl shadow-md overflow-hidden"
          >
            <!-- Bookmark icon -->
            <div class="absolute top-3 right-3 z-10">
              <font-awesome-icon
                icon="bookmark"
                class="text-gray-400 hover:text-[#24515b] cursor-pointer"
                @click="toggleBookmark(recipe)"
              />
            </div>
            <!-- Image -->
            <img
              :src="recipe.image"
              alt="Recipe Image"
              class="w-full h-48 object-cover"
            />
            <!-- Content -->
            <div class="p-4">
              <h2 class="text-lg font-bold mb-2">{{ recipe.title }}</h2>
              <p class="text-sm text-gray-600">By {{ recipe.creator }}</p>
              <p class="text-md font-semibold mt-2 text-[#24515b]">
                {{ recipe.price }}
              </p>
            </div>
          </div>
        </div>

        <!-- No recipes message -->
        <div
          v-else
          class="text-center text-xl font-medium p-6 bg-white rounded-xl shadow-md w-2/3 mx-auto"
        >
          <p>No recipes found for this category.</p>
        </div>
</div>

</template>

<script setup>

import { ref, computed } from "vue";

const props = defineProps({
  activeTab: {
    type: String,
    required: true,
  },
});

const recipes = {
  meal: [
    {
      id: 1,
      title: "Spaghetti Carbonara",
      creator: "Chef Anna",
      price: "$12",
      image: "/assets/images/spaghetti.jpg",
    },
    {
      id: 2,
      title: "Beef Steak",
      creator: "Chef John",
      price: "$18",
      image: "/assets/images/steak.jpg",
    },
  ],
  cuisine: [
    {
      id: 3,
      title: "Sushi Deluxe",
      creator: "Chef Sato",
      price: "$20",
      image: "/assets/images/sushi.jpg",
    },
  ],
  diet: [
    {
      id: 4,
      title: "Vegan Buddha Bowl",
      creator: "Chef Lisa",
      price: "$15",
      image: "/assets/images/buddha-bowl.jpg",
    },
  ],
  "cooking-method": [
    {
      id: 5,
      title: "Grilled Salmon",
      creator: "Chef Mike",
      price: "$22",
      image: "/assets/images/salmon.jpg",
    },
  ],
  "course-type": [
    {
      id: 6,
      title: "Chocolate Cake",
      creator: "Chef Emma",
      price: "$8",
      image: "/assets/images/cake.jpg",
    },
  ],
  "special-occasions": [
    {
      id: 7,
      title: "Thanksgiving Turkey",
      creator: "Chef Paul",
      price: "$30",
      image: "/assets/images/turkey.jpg",
    },
  ],
};


// Computed to get recipes for current tab
const filteredRecipes = computed(() => {
  return recipes[activeTab.value] || [];
});

// Router instance
const router = useRouter();

// Method to navigate to recipe page
const goToRecipePage = (id) => {
  router.push({ name: "recipe", params: { id } });
};

// Method to toggle bookmark
const toggleBookmark = (recipe) => {
  let bookmarks = JSON.parse(localStorage.getItem("bookmarks")) || [];
  const index = bookmarks.findIndex((item) => item.id === recipe.id);
  if (index === -1) {
    bookmarks.push(recipe);
  } else {
    bookmarks.splice(index, 1);
  }
  localStorage.setItem("bookmarks", JSON.stringify(bookmarks));
};

</script>

<style lang="scss" scoped>

</style>