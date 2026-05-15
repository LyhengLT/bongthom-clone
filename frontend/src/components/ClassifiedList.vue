// Right side Box
<template>
  <li v-for="ad in ads" :key="ad.title">
    <a href="#"><b>{{ ad.title }}</b></a>
    <br><em>{{ ad.category }}</em>
  </li>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const ads = ref([])

onMounted(async () => {
  const res = await fetch('http://localhost:3000/latest-classified')
  ads.value = await res.json()
})
</script>