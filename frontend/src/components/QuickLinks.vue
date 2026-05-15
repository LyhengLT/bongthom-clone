<template>
  <template v-for="link in links" :key="link.label">
    <tr v-if="link.img" class="leftinfo" id="facebookicon">
      <td>
        <a :href="link.url" target="_blank">
          <img :src="link.img" alt="Facebook">
          {{ link.label }}
        </a>
      </td>
    </tr>
    <tr v-else class="leftinfo" style="font-weight: bold;">
      <td>
        »<a :href="link.url">{{ link.label }}</a>
        <span v-if="link.extra" v-html="link.extra"></span>
      </td>
    </tr>
  </template>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const links = ref([])

onMounted(async () => {
  const res = await fetch('http://localhost:3000/quick-link')
  links.value = await res.json()
})
</script>