<!-- <template>
  <div>
    <p v-if="loading">Loading...</p>
    <p v-if="error">Error: {{ error.message }}</p>
    <p v-if="data">There are {{ data.recipes.length }} recipes.</p>
    <p v-for="recipe in data?.recipes" :key="recipe.id">{{ recipe.title }}</p>
    <form @submit.prevent="createRecipe">
      <input v-model="newRecipeTitle" placeholder="Recipe Title" />
      <input v-model="newRecipeDescription" placeholder="Recipe Description" />
      <button type="submit">Create Recipe</button>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useQuery, useMutation } from '@vue/apollo-composable'
import { GET_RECIPES, CREATE_RECIPE } from '~/src/queries'

const { result: data, loading, error } = useQuery(GET_RECIPES)

const newRecipeTitle = ref('')
const newRecipeDescription = ref('')

const { mutate: createRecipeMutation } = useMutation(CREATE_RECIPE, {
  update: (cache, { data: { insert_recipes } }) => {
    const existingRecipes = cache.readQuery({ query: GET_RECIPES })
    cache.writeQuery({
      query: GET_RECIPES,
      data: {
        recipes: [...existingRecipes.recipes, ...insert_recipes.returning],
      },
    })
  },
})

const createRecipe = async () => {
  try {
    console.log('Creating recipe with title:', newRecipeTitle.value)
    await createRecipeMutation({
      title: newRecipeTitle.value,
      description: newRecipeDescription.value,
    })
    newRecipeTitle.value = ''
    newRecipeDescription.value = ''
  } catch (err) {
    console.error('Error creating recipe:', err)
  }
}

if (error.value) {
  console.error('Error fetching recipes:', error.value)
}
</script>
 -->
<!-- filepath: /Users/mac/Desktop/food-recipe-app/frontend/pages/about.vue -->
<template>
  <div>
    <h1>About Page</h1>

  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useQuery } from '@vue/apollo-composable';

import { GET_RECIPES } from '../src/queries';

const { data, loading, error } = useQuery(GET_RECIPES);
</script>