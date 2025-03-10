<template>
  <div class="min-h-screen bg-gray-100 p-6">
    <div class="max-w-5xl mx-auto bg-white shadow-lg rounded-lg overflow-hidden">
      <!-- Header -->
      <div class="bg-[#24515b] p-6">
        <h2 class="text-2xl font-bold text-white">Account Settings</h2>
        <p class="text-gray-200">Manage your personal information and preferences</p>
      </div>

      <!-- Content -->
      <div class="p-6 grid grid-cols-1 md:grid-cols-2 gap-6">
        <!-- Profile Info -->
        <div>
          <h3 class="text-lg font-semibold text-gray-700 mb-4">Profile Information</h3>
          <form class="space-y-4">
            <div>
              <label class="block text-gray-600">Full Name</label>
              <input type="text" placeholder="Your name" class="w-full mt-1 px-4 py-2 rounded-md border border-gray-300" />
            </div>
            <div>
              <label class="block text-gray-600">Email Address</label>
              <input type="email" placeholder="you@example.com" class="w-full mt-1 px-4 py-2 rounded-md border border-gray-300" />
            </div>
            <button type="button" class="bg-[#24515b] text-white px-4 py-2 rounded-md">Save Changes</button>
          </form>
        </div>

        <!-- Change Password -->
        <div>
          <h3 class="text-lg font-semibold text-gray-700 mb-4">Change Password</h3>
          <form class="space-y-4">
            <div>
              <label class="block text-gray-600">Current Password</label>
              <input type="password" placeholder="Current password" class="w-full mt-1 px-4 py-2 rounded-md border border-gray-300" />
            </div>
            <div>
              <label class="block text-gray-600">New Password</label>
              <input type="password" placeholder="New password" class="w-full mt-1 px-4 py-2 rounded-md border border-gray-300" />
            </div>
            <div>
              <label class="block text-gray-600">Confirm Password</label>
              <input type="password" placeholder="Confirm new password" class="w-full mt-1 px-4 py-2 rounded-md border border-gray-300" />
            </div>
            <button type="button" class="bg-[#24515b] text-white px-4 py-2 rounded-md">Update Password</button>
          </form>
        </div>

        <!-- Preferences -->
        <div>
          <h3 class="text-lg font-semibold text-gray-700 mb-4">Preferences</h3>
          <form class="space-y-4">
            <div>
              <label class="block text-gray-600">Preferred Cuisine</label>
              <select class="w-full mt-1 px-4 py-2 rounded-md border border-gray-300">
                <option>Italian</option>
                <option>Mexican</option>
                <option>Asian</option>
                <option>Mediterranean</option>
                <option>American</option>
              </select>
            </div>
            <div>
              <label class="block text-gray-600">Diet Type</label>
              <select class="w-full mt-1 px-4 py-2 rounded-md border border-gray-300">
                <option>Vegetarian</option>
                <option>Vegan</option>
                <option>Keto</option>
                <option>Gluten-Free</option>
                <option>High-Protein</option>
              </select>
            </div>
            <button type="button" class="bg-[#24515b] text-white px-4 py-2 rounded-md">Save Preferences</button>
          </form>
        </div>

        <!-- Delete Account -->
        <div>
          <h3 class="text-lg font-semibold text-red-600 mb-4">Delete Account</h3>
          <p class="text-gray-600 mb-4">Permanently delete your account and all data. This action cannot be undone.</p>
          <button type="button" class="bg-red-600 text-white px-4 py-2 rounded-md">Delete Account</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
useHead({
  title: 'Settings | Kimem-Recipe',
  meta: [
    { name: 'description', content: 'Manage your profile, preferences, and account settings in Kimem-Recipe' }
  ]
});
</script>
