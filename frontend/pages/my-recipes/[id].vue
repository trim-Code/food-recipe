<template>
  <div >
    <Head>
      <Title>Product | {{ product.title }}</Title>
      <Meta name="description" :content="product.description" />
    </Head>
    <ProductDetails :product="product" />

  </div>
</template>

<script setup>
const { id } = useRoute().params;
definePageMeta({
  layout: "products",
});
const uri="https://fakestoreapi.com/products/"+id;
const {data:product}= await useFetch(uri,{key:id});

if (!product.value) {
  
  
  throw createError({
    statusCode: 404,
    message: "Product not found",
    fatal: true
  });

  //falta:true forces the error to be displayed even for bworser errors
} 

</script>
