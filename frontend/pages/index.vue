<template>
  <div class="w-screen min-h-screen flex flex-col">
    <Navbar />

    <NuxtPage />

    <!-- Hero Section -->
    <div class="relative h-screen w-screen flex-grow">
      <div
        class="absolute inset-0 bg-[url('/assets/images/bg-home.jpg')] bg-cover bg-no-repeat bg-center"
      ></div>
      <div class="absolute inset-0 bg-black bg-opacity-40"></div>
      <div class="absolute inset-0 bg-black bg-opacity-50"></div>
      <div
        class="relative flex flex-col items-center justify-center h-full w-3/4 m-auto"
      >
        <h1 class="text-4xl w-2/3 text-center font-bold text-[#faebefe5] mb-4">
          Master new flavors and turn every
          <span class="bg-[#24515b] px-3 inline-block transform -skew-x-12"
            >meal</span
          >
          into a masterpiece!
        </h1>
        <p class="mt-3 pb-6 text-lg text-center w-2/3 text-[#faebefc9]">
          Unlock the secrets of delicious cooking. Explore Cook and Savor Every
          Bite! Shop cook and savor – your favorite recipes anytime anywhere!
        </p>

        <button class="btn relative -bottom-40 z-60 group">
          <span>Explore Now</span>
          <font-awesome-icon
            icon="arrow-down"
            class="w-4 text-[#302a2b] group-hover:size-5"
          />
        </button>
      </div>
    </div>

    <!-- Main Section with Tabs -->
    <main class="w-full pt-8 bg-[#fdf2e9] min-h-screen">
      <div class="mx-auto container px-4">
        <!-- Tab Buttons -->
        <ul class="flex justify-center gap-4 items-center w-screen mb-8">
          <li v-for="tab in tabs" :key="tab.key">
            <button
              @click="selectTab(tab.key)"
              class="btn-catagories px-8 uppercase bg-[#e67e22]"
              :class="{ 'bg-[#e67e22] text-white ': tab.key === activeTab } "
            >
              {{ tab.label }}
            </button>
          </li>
        </ul>

        <!-- Recipe Cards -->
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
              :src="getImageUrl(recipe.image)"
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
    </main>
  </div>
</template>

<script setup>
import { ref, computed } from "vue";
import { useRouter } from "vue-router";

// Tabs array
const tabs = [
  { key: "meal", label: "Meal" },
  { key: "cuisine", label: "Cuisine" },
  { key: "diet", label: "Diet" },
  { key: "cooking-method", label: "Cooking Method" },
  { key: "course-type", label: "Course Type" },
  { key: "special-occasions", label: "Special Occasions" },
];

// Recipe data mapped by tab
const recipes = {
  meal: [
    {
      id: 1,
      title: "Spaghetti Carbonara",
      creator: "Chef Anna",
      price: "$12",
      image: "dd.jpeg",
    },
    {
      id: 2,
      title: "Beef Steak",
      creator: "Chef John",
      price: "$18",
      image: "ss.jpg",
    },
  ],
  cuisine: [
    {
      id: 3,
      title: "Sushi Deluxe",
      creator: "Chef Sato",
      price: "$20",
      image: "uu.jpg",
    },
  ],
  diet: [
    {
      id: 4,
      title: "Vegan Buddha Bowl",
      creator: "Chef Lisa",
      price: "$15",
      image: "bg.jpg",
    },
  ],
  "cooking-method": [
    {
      id: 5,
      title: "Grilled Salmon",
      creator: "Chef Mike",
      price: "$22",
      image: "tt.jpg",
    },
  ],
  "course-type": [
    {
      id: 6,
      title: "Chocolate Cake",
      creator: "Chef Emma",
      price: "$8",
      image: "dd.jpeg",
    },
  ],
  "special-occasions": [
    {
      id: 7,
      title: "Thanksgiving Turkey",
      creator: "Chef Paul",
      price: "$30",
      image: "bg-home.jpg",
    },
  ],
};

// Reactive active tab
const activeTab = ref(tabs[0].key); // Default active tab


const selectTab = (key) => {
  activeTab.value = key;
};


const filteredRecipes = computed(() => {
  return recipes[activeTab.value] || [];
});


const getImageUrl = (imageName) => {
  return new URL(`/assets/images/${imageName}`, import.meta.url).href;
};


const router = useRouter();


const goToRecipePage = (id) => {
  router.push({... recipes[activeTab.value] , params: { id } });
};


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


useHead({
  title: "Kimem-Recipe",
  meta: [
    {
      name: "description",
      content:
        "A comprehensive, user-friendly app for finding and sharing delicious recipes.",
    },
  ],
});
</script>

<style scoped>

.btn-catagories {
  @apply py-2 rounded-full  shadow hover:bg-[#eea564] transition-all;
}

.card:hover {
  @apply shadow-lg transform scale-105 transition-transform;
}
</style>