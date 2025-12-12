<script setup>
  import { ref } from "vue"

  const {data, pending} = await useFetch("/api/test/getDividendItemTest")
  const dividData = ref(data.value)

  // const {data, pending} = await useFetch("/api/test/getGrowthItemTest")
  // const dividData = ref(data.value)

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
    {{ JSON.stringify(dividData, null, 3) }}
  </pre>
</template>