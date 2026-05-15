// RSS File
<template>
  <template v-for="item in items" :key="item.link">
    <tr class="rss-info">
      <td>
        <a href="#">
          <img src="/images/feed-icon.gif" alt="rss" class="rss-icon">
          {{ item.link }}
        </a>
      </td>
    </tr>
  </template>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const items = ref([])

onMounted(async () => {
  const res = await fetch('http://localhost:3000/rss')
  items.value = await res.json()
})
</script>