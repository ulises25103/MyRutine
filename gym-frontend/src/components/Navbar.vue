<template>
  <nav class="bg-gray-800 text-white px-6 py-4 flex justify-between items-center shadow-md">
    <span class="text-xl font-bold text-primary">MiGym</span>
    <button
      @click="logout"
      class="text-red-400 hover:text-red-300 font-semibold cursor-pointer"
    >
      Cerrar sesión
    </button>
  </nav>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()

const logout = () => {
  localStorage.removeItem('token')
  router.push('/login')
}
</script>

<style scoped>
.text-primary {
  @apply text-green-400;
}
</style>
