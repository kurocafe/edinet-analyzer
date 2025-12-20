<template>
  <table class="w-full border-collapse mb-4">
    <thead>
      <tr>
        <th class="company-table-th">ランキング</th>
        <th class="company-table-th">企業名</th>
        <th class="company-table-th">証券コード</th>
        <th class="company-table-th">成長率</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="company in props.rankingdata" :key="company.secCode" class="company-table-tr">
        <td class="company-table-td">{{ company.rank }}</td>
        <td class="company-table-td">
          <NuxtLink :to="`/detail/${company.companyId}`" class=" hover:text-blue-500">{{ company.companyName }}</NuxtLink>
        </td>
        <td class="company-table-td">{{ company.secCode }}</td>
        <td class="company-table-td">{{ company.growthRate }}</td>
      </tr>
    </tbody>
  </table>

</template>

<script setup lang="ts">
import type { RankingItem } from '~/types/ranking';

  const props = defineProps<{
    rankingdata: RankingItem[]
  }>()

</script>