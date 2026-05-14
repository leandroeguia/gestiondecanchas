<template>
  <div class="min-h-screen bg-slate-900 text-white flex flex-col items-center justify-center font-sans">
    <div class="p-8 border border-slate-700 rounded-2xl bg-slate-800 shadow-xl text-center">
      <h1 class="text-4xl font-black text-green-400 mb-2">Nuxt 3 + Go</h1>
      <p class="text-slate-400 mb-8">Probando la conexión del sistema de canchas</p>

      <div class="bg-slate-900 p-4 rounded-lg border border-slate-700 inline-block min-w-[300px]">
        <p class="text-sm uppercase tracking-widest text-slate-500 mb-2">Respuesta del Backend:</p>
        
        <div v-if="pending" class="animate-pulse text-yellow-500">Conectando con Go...</div>
        <div v-else-if="error" class="text-red-400">❌ Error: No se pudo contactar a Go</div>
        <div v-else class="text-xl font-mono text-blue-300">
          {{ data.mensaje }}
        </div>
      </div>

      <button 
        @click="refresh" 
        class="mt-6 block w-full py-3 px-6 bg-green-600 hover:bg-green-500 transition-colors rounded-xl font-bold"
      >
        Refrescar Conexión
      </button>
    </div>
  </div>
</template>

<script setup>
// useFetch es la forma mágica de Nuxt de pedir datos. 
// Gracias al proxy en nuxt.config.ts, /api/hola -> localhost:8080/api/hola
const { data, pending, error, refresh } = await useFetch('/api/hola')
</script>