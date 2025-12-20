<template>
  <div>
    all graphs should be here!!
    <h1>Revenues: </h1>
    <Table :rankingdata="revenues" />
    <pre>
      {{ JSON.stringify(revenues, null, 3) }}
    </pre>

    <h1>Profits: </h1>
    <Table :rankingdata="profits" />
    <pre>
      {{ JSON.stringify(profits, null, 3) }}
    </pre>

    <h1>Incomes:</h1>
    <Table :rankingdata="incomes" />
    <pre>
      {{ JSON.stringify(incomes, null, 3) }}
    </pre>
  </div>
</template>

<script setup>
import Table from '~/components/Ranking/Table.vue';

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