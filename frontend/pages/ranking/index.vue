<template>
  <div>
    all graphs should be here!!
  </div>
</template>

<script setup>
  const {
    revenues, 
    profits, 
    incomes, 
    pending, 
    error, 
    fetchIncomes, 
    fetchProfits, 
    fetchRevenues
  } = useRanking()
  
  await fetchIncomes()
  await fetchRevenues()
  await fetchProfits()
</script>