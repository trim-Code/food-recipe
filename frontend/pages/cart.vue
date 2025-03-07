<template>
  <div class="container mx-auto px-4 py-12">
    <h1 class="text-3xl font-bold m-8 text-center">Your Recipe Cart</h1>
    
    <div class="bg-black shadow-lg rounded-lg p-6">
      <div v-for="(recipe, index) in recipes" :key="index" class="flex items-center justify-between border-b pb-4 mb-4">
        <div class="flex items-center space-x-4">
          <img :src="recipe.image" :alt="recipe.name" class="w-20 h-20 object-cover rounded-lg" />
          <div>
            <h2 class="text-lg font-semibold">{{ recipe.name }}</h2>
            <p class="text-sm text-gray-600">${{ recipe.price }}</p>
          </div>
        </div>
        <div class="flex items-center space-x-2">
          <button @click="decreaseQuantity(index)" class="px-2 py-1 bg-gray-200 rounded">-</button>
          <span class="text-lg">{{ recipe.quantity }}</span>
          <button @click="increaseQuantity(index)" class="px-2 py-1 bg-gray-200 rounded">+</button>
        </div>
      </div>
      
      <div class="text-right text-xl font-bold mt-4">Total: ${{ totalPrice }}</div>
      <button class="mt-4 w-full bg-orange-500 text-white py-3 rounded-lg">Checkout</button>
    </div>
  </div>
  


</template>

<script setup>
import { ref, computed } from 'vue';

const recipes = ref([
  { name: "Spaghetti Bolognese", image: "assets/images/jj.jpg", price: 12.99, quantity: 1 },
  { name: "Grilled Salmon", image: "assets/images/jj.jpg", price: 15.99, quantity: 1 },
  { name: "Caesar Salad", image: "assets/images/leaf.png", price: 9.99, quantity: 1 },
]);

const totalPrice = computed(() => {
  return recipes.value.reduce((sum, recipe) => sum + recipe.price * recipe.quantity, 0).toFixed(2);
});

const increaseQuantity = (index) => {
  recipes.value[index].quantity++;
};

const decreaseQuantity = (index) => {
  if (recipes.value[index].quantity > 1) {
    recipes.value[index].quantity--;
  }
};
</script>

<style scoped>
button {
  transition: background 0.3s;
}
button:hover {
  background: #ff8c42;
}
</style>
