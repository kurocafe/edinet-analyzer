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

    <pre v-if="financialData">
      {{ JSON.stringify(data, null, 3) }}
    </pre>
    <BarChart :data="revenues" :name="'収益'"/>
    <BarChart :data="operatingIncomes" :name="'営業利益'"/>
    <BarChart :data="netIncomes" :name="'純利益'"/>
    <BarChart :data="dividends" :name="'配当金'"/>
  </div>
</template>

<script setup>
import BarChart from '~/components/Ranking/BarChart.vue'

  const route = useRoute()
  const id = Number(route.params.id)

  const { pending, error, getCompanyById } = useCompany()
  const company = getCompanyById(id)
  
  const {financialData, pending: finPending, error: finError, fetchFinancial, getFinancialById, getFinancialByYear} = useFin()
  await fetchFinancial()

  const data = getFinancialById(id)
  const revenues = data.value.map(f => f.revenue)
  const operatingIncomes = data.value.map(f => f.operatingIncome)
  const netIncomes = data.value.map(f => f.netIncome)
  const dividends = data.value.map(f => f.dividend)
  
  // console.log(financialData.value)
</script>