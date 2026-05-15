// JobList
<template>
  <li v-for="job in jobs" :key="job.title">
    <a href="#">
      <img :src="job.isNew ? '/images/new.gif' : '/images/static.png'" alt="">
      <b>{{ job.title }}</b>
    </a>
    <br>
    <em>{{ job.company }}
      <img src="/images/companydetails.png" alt="companydetail">
    </em>
  </li>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const jobs = ref([])

onMounted(async () => {
  const res = await fetch('http://localhost:3000/latest-job')
  jobs.value = await res.json()
})
</script>