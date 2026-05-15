// SpecialSchedule
<template>
  <template v-for="link in items" :key="link.label">
    <tr class="leftinfo">
      <td>»<a :href="link.url">{{ link.label }} ({{ link.count }})</a></td>
    </tr>
  </template>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const items = ref([])

onMounted(async () => {
  const res = await fetch('http://localhost:3000/special-schedule')
  items.value = await res.json()
})
</script>