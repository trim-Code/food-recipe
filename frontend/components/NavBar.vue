<template>
  <nav class="bg-black w-full shadow-lg  z-50">
    <div class="container mx-auto flex items-center justify-between py-4 px-0">
      <!-- Logo -->
      <NuxtLink to="/">
        <div class="w-10 font-bold text-yellow-500">
          <img src="/logo.png" alt="" />
        </div>
      </NuxtLink>

      <!-- Desktop Navigation -->
      <div class="hidden md:flex space-x-8">
        <NuxtLink to="/"
          ><button class="font-semibold text-white">HOME</button></NuxtLink
        >
        <NuxtLink to="/how-it-works"
          ><button class="font-semibold text-white">
            HOW IT WORKS
          </button></NuxtLink
        >
        <div
          v-for="item in menuItems"
          :key="item.label"
          class="relative group"
          @mouseenter="showDropdown = item.label"
          @mouseleave="showDropdown = null"
        >
          <NuxtLink to="/my-recipes">
            <button class="font-semibold text-white">RECIPES</button>
          </NuxtLink>
          <div
            v-if="showDropdown === '/my-recipes'"
            class="absolute -left-8 mt-0 w-64 bg-black text-white shadow-lg p-6 rounded-md transition-opacity duration-300 ease-in-out"
            @mouseenter="showDropdown = item.label"
            @mouseleave="showDropdown = null"
          >
            <ul class="space-y-2">
              <li
                v-for="dropdownItem in item.dropdown"
                :key="dropdownItem"
                class="hover:text-yellow-400 cursor-pointer"
              >
                {{ dropdownItem }}
              </li>
            </ul>
          </div>
        </div>
      </div>

      <!-- Right Side Icons -->
      <div class="flex items-center md:flex space-x-8">
        <form action="" class="flex gap-3 justify-between">
          <div class="relative">
            <input
              type="text"
              name=""
              id=""
              placeholder="Search..."
              class="bg-gray-700 text-white px-5 py-2 rounded-full focus:outline-none focus:ring focus:ring-yellow-500"
            />
            <NuxtLink to="/search">
              <button class="absolute right-2 top-2 text-white">
                <font-awesome-icon icon="search" />
              </button>
            </NuxtLink>
          </div>
        </form>
        <NuxtLink to="/create-recipes">
          <button class="">
            <font-awesome-icon icon="plus" />
          </button>
        </NuxtLink>

        <button class="">
          <font-awesome-icon icon="bookmark" />
        </button>
        <NuxtLink to="/cart">
          <button class="">
            <font-awesome-icon icon="shopping-cart" />
          </button>
        </NuxtLink>

        <!-- Profile Dropdown -->
        <div class="relative group" @click="showProfile = !showProfile">
          <button class="flex items-center space-x-2 nav-link">
            <img
              src="assets/images/jj.jpg"
              alt="Profile"
              class="rounded-full w-8 h-8"
            />
            <span>John</span>
          </button>
          <div class="dropdown-menu" v-if="showProfile">
            <NuxtLink to="/my-recipes" class="dropdown-item"
              >My Recipes</NuxtLink
            >
            <NuxtLink to="/profile/settings" class="dropdown-item"
              >Settings</NuxtLink
            >
            <NuxtLink to="/auth/login" class="dropdown-item">Logout</NuxtLink>
          </div>
        </div>
      </div>

      <!-- Mobile Menu Button -->
      <button
        @click="mobileMenuOpen = !mobileMenuOpen"
        class="md:hidden text-xl"
      >
        <font-awesome-icon :icon="mobileMenuOpen ? 'times' : 'bars'" />
      </button>
    </div>

    <!-- Mobile Menu -->
    <div
      v-if="mobileMenuOpen"
      class="md:hidden bg-black text-white px-6 py-4 space-y-3 transition-transform duration-300 ease-in-out"
    >
      <NuxtLink to="/my-recipes" class="block">RECIPES</NuxtLink>
      <NuxtLink to="/popular" class="block">POPULAR</NuxtLink>
      <NuxtLink to="/holidays" class="block">HOLIDAYS</NuxtLink>
      <NuxtLink to="/healthy-diet" class="block">HEALTHY & DIET</NuxtLink>
      <NuxtLink to="/cuisine" class="block">CUISINE</NuxtLink>
      <NuxtLink to="/seasonal" class="block">SEASONAL</NuxtLink>
    </div>
  </nav>
</template>

<script setup>
import { ref } from "vue";

const showDropdown = ref(null);
const mobileMenuOpen = ref(false);
const showProfile = ref(false);

const menuItems = [
  {
    label: "RECIPES",
    dropdown: [
      "Breakfast & Brunch",
      "Lunch",
      "Dinner",
      "Desserts",
      "BBQ & Grill",
      "Baking",
      "See More",
    ],
  },
  // {
  //   label: "POPULAR",
  //   dropdown: ["Trending", "Top Rated", "New"],
  // },
  // {
  //   label: "HOLIDAYS",
  //   dropdown: ["Christmas", "Thanksgiving", "Easter"],
  // },
  // {
  //   label: "HEALTHY & DIET",
  //   dropdown: ["Low Carb", "Keto", "Vegan"],
  // },
  // {
  //   label: "CUISINE",
  //   dropdown: ["Italian", "Mexican", "Chinese"],
  // },
  // {
  //   label: "SEASONAL",
  //   dropdown: ["Spring", "Summer", "Fall", "Winter"],
  // },
];
</script>

<style scoped>
.nav-link {
  @apply text-gray-300  transition-all px-3 py-2;
}

.icon-link {
  @apply text-xl  transition-all px-3;
}

.dropdown-menu {
  @apply absolute hidden group-hover:block bg-black shadow-lg rounded-lg py-2 w-60 mt-0;
}

.dropdown-item {
  @apply block px-4 py-2;
}

.mobile-link {
  @apply block px-4 py-2 text-gray-700;
}
</style>
