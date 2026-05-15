// Other Ann
<template>
  <template v-for="link in items" :key="link.label">
    <tr class="leftinfo">
      <td>»<a href="#">{{ link.label }} ({{ link.count }})</a></td>
    </tr>
  </template>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const items = ref([])

onMounted(async () => {
  const res = await fetch('http://localhost:3000/other-announcement')
  items.value = await res.json()
})
</script>