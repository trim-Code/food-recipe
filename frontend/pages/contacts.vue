<template>
  <div
    class="  w-full mx-auto py-12 px-4 sm:px-6 lg:px-8  bg-[#f7f6f6] flex flex-col items-center  h-screen "
  >

  
  <div class=" flex flex-col items-center justify-center mb-8  "> 
    <h1 class="text-2xl font-bold text-[#153342e1]">Get in Touch</h1>
    <img src="/assets/images/heading-shap.png" alt="heading-shape" class="w-3/4" />
  </div>
  <div
  class=" grid grid-cols-1 md:flex gap-8 shadow-md rounded-3xl py-12 pl-8 bg-white"
  >

      <div 
        class="bg-[#ea5440] text-white rounded-2xl p-8 space-y-6 shadow-lg w-1/3 relative"
      >
      <!-- <img src="assets/images/3.png" alt="" class="absolute -top-32 -right-16 h-1/2  transform rotate-12" /> -->
        <h3 class="text-sm uppercase">Connect With Us</h3>
        <h2 class="text-3xl font-bold">Join the Kimem Recipe Community</h2>
        <div class="space-y-6">
          <div class="flex items-center space-x-3">
            <div class="flex items-center justify-center rounded-full p-3 border border-gray-300">
              <font-awesome-icon icon="location-dot" class="text-xl text-gray-200  " />
            </div>
            <div>
              <p class="font-semibold">Headquarter</p>
              <p class="text-sm">AddisAbaba, Ethiopia</p>
            </div>
          </div>
          <div class="flex items-center space-x-3">
            <div class="flex items-center justify-center rounded-full p-3 border border-gray-300">
              <font-awesome-icon icon="envelope" class="text-xl text-gray-200   " />
            </div>
            <div>
              <p class="font-semibold">Contact Us</p>
              <p class="text-sm">info@kimemrecipe.com</p>
            </div>
          </div>
          <div class="flex items-center space-x-3">
            <div class="flex items-center justify-center rounded-full p-3 border-gray-300 border">
              <font-awesome-icon icon="phone" class="text-xl text-gray-200  " />
            </div>
            <div>
              <p class="font-semibold">Call Us</p>
              <p class="text-sm">(+251) 900-0000</p>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-2xl p-8 space-y-6">
        <form class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div>
            <label for="first-name" class="block text-sm font-medium text-black"
              >First Name</label
            >
            <input type="text" id="first-name" class="input-field  " />
          </div>
          <div>
            <label for="last-name" class="block text-sm font-medium text-black"
              >Last Name</label
            >
            <input type="text" id="last-name" class="input-field" />
          </div>
          <div>
            <label for="email" class="block text-sm font-medium text-black"
              >Email Address</label
            >
            <input type="email" id="email" class="input-field" />
          </div>
          <div>
            <label for="phone" class="block text-sm font-medium text-black"
              >Phone Number</label
            >
            <input type="tel" id="phone" class="input-field" />
          </div>

          <div class="col-span-2">
            <label for="message" class="block text-sm font-medium text-black"
              >Message</label
            >
            <textarea
              id="message"
              rows="3"
              class="input-field w-full"
            ></textarea>
          </div>

          <button
            class="bg-[#ea5440] text-white px-5 py-3 rounded-lg hover:bg-orange-600 transition-all mt-8"
          >
            Send Your Feedback
          </button>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup></script>

<style scoped>
.input-field {
  @apply border-b border-gray-400 rounded-md p-1 w-full bg-white text-gray-700;
}
.input-field:focus {
  @apply border-orange-500 outline-none;
}
</style>
