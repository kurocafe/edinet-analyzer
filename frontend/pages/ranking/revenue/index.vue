<script setup>
  import {ref} from "vue"
  const {data, pending} = await useFetch("/api/test/getGrowthItemTest")
  const test = ref(data.value)

  console.log(data.value)

  function changeTest(input) {
    test.value = input
  }
</script>

<template>
  <div v-if="pending">
    loading...
  </div>
  <pre v-else>
    {{ JSON.stringify(test, null, 3) }}
  </pre>
  <input v-model="test" class="border-2 border-red-200 rounded-2xl text-center"/>
</template>