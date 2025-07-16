<template>
  <Navbar />
  <div class="min-h-screen bg-gray-900 text-white p-6">
    <h1 class="text-3xl font-bold mb-6">Tus Rutinas</h1>
    <ul>
      <li
        v-for="routine in routines"
        :key="routine.ID"
        class="mb-4 p-4 rounded-lg bg-gray-800 shadow"
      >
        <h2 class="text-xl font-semibold">{{ routine.Name }}</h2>
        <p class="text-sm text-gray-400">ID: {{ routine.ID }}</p>
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import Navbar from '../components/Navbar.vue'
import { ref, onMounted } from 'vue'
import api from '../services/api.ts'

interface Routine {
  ID: number
  Name: string
  // Agrega otros campos si tu backend los retorna
}

const routines = ref<Routine[]>([])

onMounted(async () => {
  try {
    const res = await api.get('/routines')
    routines.value = res.data
  } catch (error) {
    console.error('Error al obtener las rutinas:', error)
    alert('No se pudieron cargar las rutinas')
  }
})
</script>
