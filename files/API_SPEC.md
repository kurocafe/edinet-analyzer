# EDINET Analyzer API仕様書

## ベースURL

```
http://localhost:8080
```

---

## 1. 企業関連API

### GET /api/v1/companies

**説明:** 企業一覧を取得

**リクエスト:**
```http
GET /api/v1/companies
```

**レスポンス:**
```json
[
  {
    "ID": 1,
    "CreatedAt": "2024-12-13T10:00:00Z",
    "UpdatedAt": "2024-12-13T10:00:00Z",
    "DeletedAt": null,
    "name": "ソニーグループ",
    "secCode": "6758",
    "edinetCode": ""
  },
  {
    "ID": 2,
    "CreatedAt": "2024-12-13T10:00:00Z",
    "UpdatedAt": "2024-12-13T10:00:00Z",
    "DeletedAt": null,
    "name": "任天堂",
    "secCode": "7974",
    "edinetCode": ""
  }
]
```

---

### GET /api/v1/companies/:id

**説明:** 特定企業の詳細を取得

**リクエスト:**
```http
GET /api/v1/companies/1
```

**レスポンス:**
```json
{
  "ID": 1,
  "name": "ソニーグループ",
  "secCode": "6758",
  "edinetCode": "",
  "financialData": [
    {
      "ID": 1,
      "companyId": 1,
      "fiscalYear": 2023,
      "revenue": 10000000000000,
      "operatingIncome": 1500000000000,
      "netIncome": 1400000000000,
      "dividend": 120
    }
  ]
}
```

---

### POST /api/v1/companies

**説明:** 企業を追加（テスト用）

**リクエスト:**
```json
{
  "name": "ソニーグループ",
  "secCode": "6758"
}
```

**レスポンス:**
```json
{
  "ID": 1,
  "name": "ソニーグループ",
  "secCode": "6758",
  "edinetCode": ""
}
```

---

## 2. 財務データ関連API

### GET /api/v1/financial-data

**説明:** 全財務データを取得

**リクエスト:**
```http
GET /api/v1/financial-data
```

**クエリパラメータ（任意）:**
- `fiscalYear` - 年度で絞り込み（例: 2023）
- `companyId` - 企業IDで絞り込み（例: 1）

**例:**
```http
GET /api/v1/financial-data?fiscalYear=2023
GET /api/v1/financial-data?companyId=1
```

**レスポンス:**
```json
[
  {
    "ID": 1,
    "companyId": 1,
    "fiscalYear": 2023,
    "revenue": 10000000000000,
    "operatingIncome": 1500000000000,
    "netIncome": 1400000000000,
    "dividend": 120,
    "company": {
      "ID": 1,
      "name": "ソニーグループ",
      "secCode": "6758"
    }
  }
]
```

---

### GET /api/v1/companies/:id/financial-data

**説明:** 特定企業の財務データを取得

**リクエスト:**
```http
GET /api/v1/companies/1/financial-data
```

**レスポンス:**
```json
[
  {
    "ID": 1,
    "companyId": 1,
    "fiscalYear": 2023,
    "revenue": 10000000000000,
    "operatingIncome": 1500000000000,
    "netIncome": 1400000000000,
    "dividend": 120
  },
  {
    "ID": 2,
    "companyId": 1,
    "fiscalYear": 2022,
    "revenue": 9000000000000,
    "operatingIncome": 1200000000000,
    "netIncome": 1100000000000,
    "dividend": 110
  }
]
```

---

### POST /api/v1/financial-data

**説明:** 財務データを追加（テスト用）

**リクエスト:**
```json
{
  "companyId": 1,
  "fiscalYear": 2023,
  "revenue": 10000000000000,
  "operatingIncome": 1500000000000,
  "netIncome": 1400000000000,
  "dividend": 120
}
```

**レスポンス:**
```json
{
  "ID": 1,
  "companyId": 1,
  "fiscalYear": 2023,
  "revenue": 10000000000000,
  "operatingIncome": 1500000000000,
  "netIncome": 1400000000000,
  "dividend": 120
}
```

---

## 3. ランキング関連API

> **注意:** すべてのランキングはEDINETから取得可能なデータのみを使用します。株価などの外部データは使用しません。

### GET /api/v1/rankings/revenue-growth

**説明:** 売上高成長率ランキングを取得

**リクエスト:**
```http
GET /api/v1/rankings/revenue-growth?limit=10
```

**クエリパラメータ:**
- `limit` - 取得件数（デフォルト: 10）
- `baseYear` - 基準年度（デフォルト: 2022）
- `targetYear` - 比較年度（デフォルト: 2023）

**例:**
```http
GET /api/v1/rankings/revenue-growth?limit=10&baseYear=2022&targetYear=2023
```

**レスポンス:**
```json
[
  {
    "rank": 1,
    "company": {
      "ID": 1,
      "name": "ソニーグループ",
      "secCode": "6758"
    },
    "baseYear": 2022,
    "targetYear": 2023,
    "baseRevenue": 9000000000000,
    "targetRevenue": 10000000000000,
    "growthRate": 11.11
  },
  {
    "rank": 2,
    "company": {
      "ID": 2,
      "name": "任天堂",
      "secCode": "7974"
    },
    "baseYear": 2022,
    "targetYear": 2023,
    "baseRevenue": 1000000000000,
    "targetRevenue": 1200000000000,
    "growthRate": 20.0
  }
]
```

---

### GET /api/v1/rankings/profit-growth

**説明:** 営業利益成長率ランキングを取得

**リクエスト:**
```http
GET /api/v1/rankings/profit-growth?limit=10
```

**クエリパラメータ:**
- `limit` - 取得件数（デフォルト: 10）
- `baseYear` - 基準年度（デフォルト: 2022）
- `targetYear` - 比較年度（デフォルト: 2023）

**例:**
```http
GET /api/v1/rankings/profit-growth?limit=10&baseYear=2022&targetYear=2023
```

**レスポンス:**
```json
[
  {
    "rank": 1,
    "company": {
      "ID": 1,
      "name": "ソニーグループ",
      "secCode": "6758"
    },
    "baseYear": 2022,
    "targetYear": 2023,
    "baseProfit": 1200000000000,
    "targetProfit": 1500000000000,
    "growthRate": 25.0
  }
]
```

---

### GET /api/v1/rankings/income-growth

**説明:** 純利益成長率ランキングを取得

**リクエスト:**
```http
GET /api/v1/rankings/income-growth?limit=10
```

**クエリパラメータ:**
- `limit` - 取得件数（デフォルト: 10）
- `baseYear` - 基準年度（デフォルト: 2022）
- `targetYear` - 比較年度（デフォルト: 2023）

**例:**
```http
GET /api/v1/rankings/income-growth?limit=10&baseYear=2022&targetYear=2023
```

**レスポンス:**
```json
[
  {
    "rank": 1,
    "company": {
      "ID": 1,
      "name": "ソニーグループ",
      "secCode": "6758"
    },
    "baseYear": 2022,
    "targetYear": 2023,
    "baseIncome": 1100000000000,
    "targetIncome": 1400000000000,
    "growthRate": 27.27
  }
]
```

---

### GET /api/v1/rankings/dividend

**説明:** 配当金額ランキングを取得

> **注意:** 配当金額（円/株）のランキングです。配当利回り（%）ではありません。配当利回りの計算には株価が必要で、EDINETでは取得できません。

**リクエスト:**
```http
GET /api/v1/rankings/dividend?fiscalYear=2023&limit=10
```

**クエリパラメータ:**
- `fiscalYear` - 対象年度（デフォルト: 2023）
- `limit` - 取得件数（デフォルト: 10）

**レスポンス:**
```json
[
  {
    "rank": 1,
    "company": {
      "ID": 1,
      "name": "ソニーグループ",
      "secCode": "6758"
    },
    "fiscalYear": 2023,
    "dividend": 120
  },
  {
    "rank": 2,
    "company": {
      "ID": 2,
      "name": "任天堂",
      "secCode": "7974"
    },
    "fiscalYear": 2023,
    "dividend": 110
  }
]
```

---

## 4. EDINET関連API（Nuxtサーバー経由）

### GET /api/edinet/documents

**説明:** EDINET書類一覧を取得（Nuxtサーバー経由）

**リクエスト:**
```http
GET /api/edinet/documents?date=2024-06-28&secCode=6758
```

**クエリパラメータ:**
- `date` - 提出日（YYYY-MM-DD形式）
- `secCode` - 証券コード（任意、フィルタリング用）

**レスポンス:**
```json
[
  {
    "docID": "S100XXXX",
    "edinetCode": "E01234",
    "secCode": "67580",
    "filerName": "ソニーグループ株式会社",
    "docTypeCode": "120",
    "periodStart": "2023-04-01",
    "periodEnd": "2024-03-31",
    "submitDateTime": "2024-06-28 12:00:00"
  }
]
```

---

## データ型定義（TypeScript）

### Company（企業）

```typescript
interface Company {
  ID: number
  CreatedAt: string
  UpdatedAt: string
  DeletedAt: string | null
  name: string
  secCode: string
  edinetCode: string
  financialData?: FinancialData[]
}
```

### FinancialData（財務データ）

```typescript
interface FinancialData {
  ID: number
  companyId: number
  fiscalYear: number
  revenue: number          // 売上高（円）
  operatingIncome: number  // 営業利益（円）
  netIncome: number        // 純利益（円）
  dividend: number         // 配当（円/株）
  company?: Company
}
```

### RevenueGrowthRanking（売上成長率ランキング）

```typescript
interface RevenueGrowthRanking {
  rank: number
  company: Company
  baseYear: number
  targetYear: number
  baseRevenue: number
  targetRevenue: number
  growthRate: number       // パーセント
}
```

### ProfitGrowthRanking（営業利益成長率ランキング）

```typescript
interface ProfitGrowthRanking {
  rank: number
  company: Company
  baseYear: number
  targetYear: number
  baseProfit: number
  targetProfit: number
  growthRate: number       // パーセント
}
```

### IncomeGrowthRanking（純利益成長率ランキング）

```typescript
interface IncomeGrowthRanking {
  rank: number
  company: Company
  baseYear: number
  targetYear: number
  baseIncome: number
  targetIncome: number
  growthRate: number       // パーセント
}
```

### DividendRanking（配当金額ランキング）

```typescript
interface DividendRanking {
  rank: number
  company: Company
  fiscalYear: number
  dividend: number         // 配当金額（円/株）
}
```

---

## 実装状況

- ✅ `GET /api/v1/companies` - 完成
- ✅ `POST /api/v1/companies` - 完成
- ⏸️ `GET /api/v1/companies/:id` - 未実装
- ⏸️ `GET /api/v1/financial-data` - 未実装
- ⏸️ `POST /api/v1/financial-data` - 未実装
- ⏸️ `GET /api/v1/companies/:id/financial-data` - 未実装
- ⏸️ `GET /api/v1/rankings/revenue-growth` - 未実装
- ⏸️ `GET /api/v1/rankings/profit-growth` - 未実装
- ⏸️ `GET /api/v1/rankings/income-growth` - 未実装
- ⏸️ `GET /api/v1/rankings/dividend` - 未実装
- ⏸️ `GET /api/edinet/documents` - 未実装

---

## EDINETで取得可能なデータ

本プロジェクトでは、EDINET（金融庁 電子開示システム）から以下のデータを取得します：

- ✅ 売上高（Revenue）
- ✅ 営業利益（Operating Income）
- ✅ 純利益（Net Income）
- ✅ 配当金（Dividend per Share）

以下のデータは**外部API（株価情報API）が必要**なため、本プロジェクトでは実装しません：

- ❌ 株価（Stock Price）
- ❌ 時価総額（Market Cap）
- ❌ 配当利回り（Dividend Yield）
- ❌ PER（Price Earnings Ratio）
- ❌ PBR（Price Book-value Ratio）

---

## フロントエンド使用例

### 企業一覧取得

```vue
<script setup>
const { data: companies } = await useFetch('http://localhost:8080/api/v1/companies')
</script>

<template>
  <div v-for="company in companies" :key="company.ID">
    {{ company.name }} ({{ company.secCode }})
  </div>
</template>
```

### 売上成長率ランキング取得

```vue
<script setup>
const { data: rankings } = await useFetch(
  'http://localhost:8080/api/v1/rankings/revenue-growth',
  {
    params: {
      limit: 10,
      baseYear: 2022,
      targetYear: 2023
    }
  }
)
</script>

<template>
  <div v-for="item in rankings" :key="item.rank">
    {{ item.rank }}位: {{ item.company.name }} ({{ item.growthRate.toFixed(2) }}%)
  </div>
</template>
```

### 営業利益成長率ランキング取得

```vue
<script setup>
const { data: rankings } = await useFetch(
  'http://localhost:8080/api/v1/rankings/profit-growth',
  {
    params: {
      limit: 10,
      baseYear: 2022,
      targetYear: 2023
    }
  }
)
</script>

<template>
  <div v-for="item in rankings" :key="item.rank">
    {{ item.rank }}位: {{ item.company.name }} ({{ item.growthRate.toFixed(2) }}%)
  </div>
</template>
```

### 配当金額ランキング取得

```vue
<script setup>
const { data: rankings } = await useFetch(
  'http://localhost:8080/api/v1/rankings/dividend',
  {
    params: {
      fiscalYear: 2023,
      limit: 10
    }
  }
)
</script>

<template>
  <div v-for="item in rankings" :key="item.rank">
    {{ item.rank }}位: {{ item.company.name }} ({{ item.dividend }}円/株)
  </div>
</template>
```

### EDINET書類取得（Nuxtサーバー経由）

```vue
<script setup>
const { data: documents } = await useFetch('/api/edinet/documents', {
  params: {
    date: '2024-06-28'
  }
})
</script>

<template>
  <div v-for="doc in documents" :key="doc.docID">
    {{ doc.filerName }} - {{ doc.docID }}
  </div>
</template>
```

---

## エラーレスポンス

### 400 Bad Request

```json
{
  "error": "Invalid parameters"
}
```

### 404 Not Found

```json
{
  "error": "Resource not found"
}
```

### 500 Internal Server Error

```json
{
  "error": "Internal server error"
}
```

---

## 更新履歴

- 2024-12-13: 初版作成