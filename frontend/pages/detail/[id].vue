<template>
  <div>
    <div v-if="pending || finPending" class="loading">
      loading...
    </div>

    <div v-if="error || finError" class="err">
      {{ error }}
    </div>

    <div v-if="company">
      <h2>会社名:</h2>
      {{ company.name }}
    </div>

    <div class="text-9xl">
      ここにグラフを作成
    </div>

    <pre v-if="financialData">
      {{ JSON.stringify(data, null, 3) }}
    </pre>
  </div>
</template>

<script setup>
  const route = useRoute()
  const id = Number(route.params.id)

  const { pending, error, getCompanyById } = useCompany()
  const company = getCompanyById(id)
  
  const {financialData, pending: finPending, error: finError, fetchFinancial, getFinancialById, getFinancialByYear} = useFin()
  await fetchFinancial()

  const data = getFinancialById(id)
  
  // console.log(financialData.value)
</script>