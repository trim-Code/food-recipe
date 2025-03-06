<template>
  <div>
    <h1>About</h1>
    <p>This is a about page</p>
    <div class="text-black">{{ data }}</div>
  </div>
</template>

<script setup>
//ommit the second arg to make it get request
// const { data } = await useFetch("./api/john?name=yohannes");
// const { data } = await useFetch("./api/john?name=yohannes",{
//   method: "post",
//   body:{age: 20}
// });
// const { data } = await useFetch("./api/john");
const { data } = await useFetch("./api/currency/GBP");

useHead({
  title: "About - Recipe App",
  meta: [{ name: "description", content: "About page description" }],
});
</script>
