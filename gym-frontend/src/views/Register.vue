<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-gray-900 to-gray-800">
    <div class="bg-white dark:bg-gray-900 rounded-2xl shadow-xl p-8 w-full max-w-md border border-gray-700">
      <h2 class="text-3xl font-bold text-center text-primary mb-6 text-indigo-400">Registrarse</h2>

      <form @submit.prevent="register" class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-300 mb-1">Email</label>
          <input
            type="email"
            v-model="email"
            required
            class="w-full px-4 py-2 rounded-lg border border-gray-600 bg-gray-800 text-white focus:outline-none focus:ring-2 focus:ring-primary"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-300 mb-1">Contraseña</label>
          <input
            type="password"
            v-model="password"
            required
            class="w-full px-4 py-2 rounded-lg border border-gray-600 bg-gray-800 text-white focus:outline-none focus:ring-2 focus:ring-primary"
          />
        </div>

        <button
          type="submit"
          class="w-full py-2 bg-indigo-600 hover:bg-indigo-700 transition-colors rounded font-semibold text-white cursor-pointer"
        >
          Registrarse
        </button>
        <p v-if="error" class="text-red-400 text-center text-sm">{{ error }}</p>
      </form>

      <p class="mt-4 text-sm text-center text-gray-400">
        ¿Ya tenés cuenta?
        <router-link to="/login" class="text-primary hover:underline">Iniciar sesión</router-link>
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import api from '../services/api'

const email = ref('')
const password = ref('')
const error = ref('')
const router = useRouter()

const register = async () => {
  try {
    const response = await api.post('/register', {
      email: email.value,
      password: password.value,
    })
    console.log(response.data)
    router.push('/login')
  } catch (err: any) {
    error.value = err.response?.data?.message || 'Error al registrar'
  }
}
</script>
