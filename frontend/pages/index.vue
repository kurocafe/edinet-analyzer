<template>
  <div class=" max-w-[1000px] my-auto p-8">
    <h1>EDINET Analyzer</h1>
    <p>財務データランキングシステム</p>

    <div class="bg-white p-8 rounded-lg shadow-sm">
      <h2>企業一覧</h2>
      
      <!-- ローディング中 -->
      <div v-if="pending" class="loading">
        読み込み中...
      </div>
      
      <!-- エラー -->
      <div v-if="error" class="err">
        エラー: {{ error }}
      </div>
      
      <!-- データ表示 -->
      <div v-else>
        <p v-if="companies.length === 0">
          企業データがありません
        </p>
        
        <table v-else class="w-full border-collapse mb-4">
          <thead>
            <tr>
              <th class="company-table-th">ID</th>
              <th class="company-table-th">企業名</th>
              <th class="company-table-th">証券コード</th>
              <th class="company-table-th">業界</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="company in companies" :key="company.ID" class="company-table-tr">
              <td class="company-table-td">{{ company.ID }}</td>
              <td class="company-table-td">
                <NuxtLink :to="`/detail/${company.ID}`" class=" hover:text-blue-500">{{ company.name }}</NuxtLink>
              </td>
              <td class="company-table-td">{{ company.secCode }}</td>
              <td class="company-table-td">{{ company.industry }}</td>
            </tr>
          </tbody>
        </table>
      </div>
      
      <button @click="fetchCompany()" class=" bg-blue-500 text-white py-2 px-4 rounded-sm cursor-pointer hover:bg-[#1d4ed8]">
        再読み込み
      </button>
    </div>
  </div>
</template>

<script setup>
// nuxt.config.tsで設定した環境変数を取得
  const config = useRuntimeConfig()

  const {companies, pending, error, fetchCompany} = useCompany()
  await fetchCompany()

  
</script>