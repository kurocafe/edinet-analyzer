<template>
  <div class="container">
    <h1>EDINET Analyzer</h1>
    <p>財務データランキングシステム</p>

    <div class="section">
      <h2>企業一覧</h2>
      
      <!-- ローディング中 -->
      <div v-if="pending" class="loading">
        読み込み中...
      </div>
      
      <!-- エラー -->
      <div v-else-if="error" class="error">
        エラー: {{ error.message }}
      </div>
      
      <!-- データ表示 -->
      <div v-else>
        <p v-if="companies.length === 0">
          企業データがありません
        </p>
        
        <table v-else class="company-table">
          <thead>
            <tr>
              <th>ID</th>
              <th>企業名</th>
              <th>証券コード</th>
              <th>業界</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="company in companies" :key="company.ID">
              <td>{{ company.ID }}</td>
              <td>{{ company.name }}</td>
              <td>{{ company.secCode }}</td>
              <td>{{ company.industry }}</td>
            </tr>
          </tbody>
        </table>
      </div>
      
      <button @click="refresh()" class="refresh-btn">
        再読み込み
      </button>
    </div>
  </div>
</template>

<script setup>
// nuxt.config.tsで設定した環境変数を取得
const config = useRuntimeConfig()

// APIから企業一覧を取得
// 特に何も書かなければGETになる（デフォルト）
const { data: companies, pending, error, refresh } = await useFetch(
  `${config.public.apiBase}/api/v1/companies`,
  {
    default: () => []
  }
)
</script>

<style scoped>
.container {
  max-width: 1000px;
  margin: 0 auto;
  padding: 2rem;
}

h1 {
  color: #2563eb;
  font-size: 2.5rem;
  margin-bottom: 0.5rem;
}

p {
  color: #6b7280;
  margin-bottom: 2rem;
}

.section {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

h2 {
  font-size: 1.5rem;
  margin-bottom: 1rem;
  color: #1f2937;
}

.loading {
  text-align: center;
  padding: 2rem;
  color: #6b7280;
  font-style: italic;
}

.error {
  padding: 1rem;
  background: #fee;
  border: 1px solid #fcc;
  border-radius: 4px;
  color: #c00;
}

.company-table {
  width: 100%;
  border-collapse: collapse;
  margin-bottom: 1rem;
}

.company-table th {
  background: #f3f4f6;
  padding: 0.75rem;
  text-align: left;
  font-weight: 600;
  border-bottom: 2px solid #e5e7eb;
}

.company-table td {
  padding: 0.75rem;
  border-bottom: 1px solid #e5e7eb;
}

.company-table tr:hover {
  background: #f9fafb;
}

.refresh-btn {
  background: #2563eb;
  color: white;
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1rem;
}

.refresh-btn:hover {
  background: #1d4ed8;
}
</style>