<template>
  <div class="bookmark-page">
    <h1>Bookmarked Recipes</h1>
    <div v-if="bookmarkedRecipes.length" class="recipe-list">
      <div
        v-for="recipe in bookmarkedRecipes"
        :key="recipe.id"
        class="recipe-card"
      >
        <img :src="recipe.image" alt="Recipe Image" class="recipe-image" />
        <h2 class="recipe-title">{{ recipe.title }}</h2>
        <p class="recipe-description">{{ recipe.description }}</p>
        <button @click="removeBookmark(recipe.id)" class="remove-button">
          Remove Bookmark
        </button>
      </div>
    </div>
    <div v-else class="no-bookmarks">
      <p>No bookmarks yet. Start adding your favorite recipes!</p>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";

const bookmarkedRecipes = ref([]);

onMounted(() => {
  loadBookmarks();
});

function loadBookmarks() {
  if (typeof window !== "undefined" && typeof localStorage !== "undefined") {
    const bookmarks = JSON.parse(localStorage.getItem("bookmarks")) || [];
    bookmarkedRecipes.value = bookmarks;
  }
}

function removeBookmark(id) {
  bookmarkedRecipes.value = bookmarkedRecipes.value.filter(
    (recipe) => recipe.id !== id
  );
  if (typeof window !== "undefined" && typeof localStorage !== "undefined") {
    localStorage.setItem("bookmarks", JSON.stringify(bookmarkedRecipes.value));
  }
}
</script>

<style scoped>
.bookmark-page {
  padding: 20px;
  font-family: Arial, sans-serif;
}

h1 {
  text-align: center;
  margin-bottom: 20px;
}

.recipe-list {
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
  justify-content: center;
}

.recipe-card {
  background: #fff;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  overflow: hidden;
  width: 300px;
  text-align: center;
  padding: 20px;
}

.recipe-image {
  width: 100%;
  height: auto;
  border-bottom: 1px solid #ddd;
}

.recipe-title {
  font-size: 1.5em;
  margin: 10px 0;
}

.recipe-description {
  color: #666;
  margin-bottom: 20px;
}

.remove-button {
  background: #ff4d4d;
  color: #fff;
  border: none;
  padding: 10px 20px;
  border-radius: 5px;
  cursor: pointer;
}

.remove-button:hover {
  background: #ff1a1a;
}

.no-bookmarks {
  text-align: center;
  color: #666;
}
</style>
