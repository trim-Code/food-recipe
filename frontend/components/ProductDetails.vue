<template>
  <div class="card">
    <div class="grid grid-cols-2 gap-10">
      <div class="p-7">
        <img
          :src="product.image"
          alt="product.name"
          class="mx-auto my-7 w-full h-96 object-cover"
        />
      </div>
      <div class="p-7">
        <h2 class="text-4xl my-7 font-bold">{{ product.title }}</h2>
        <p class="text-xl font-bold my-7">Price: ${{ product.price }}</p>
        <h3 class="text-lg mb-4 pb-2 font-bold border-b-2">
          Product Description
        </h3>
        <p class="text-lg mb-7">{{ product.description }}</p>
        <button class="btn flex">
          <i class="material-icons mr-2">add_shopping_cart</i>
          <span>Add to cart</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
const { product } = defineProps(["product"]);
</script>

<style scoped>
img {
  max-width: 400px;
}
</style>
