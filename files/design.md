# 設計書：EDINET財務ランキングシステム

## 1. システム構成

### 1.1 全体アーキテクチャ

```
┌─────────────────────────────────────────────┐
│            ユーザー（ブラウザ）               │
└──────────────────┬──────────────────────────┘
                   │ HTTP/HTTPS
                   ▼
┌─────────────────────────────────────────────┐
│         Nuxt.js (SSR)                       │
│         ポート: 3000                         │
│  ┌─────────────────────────────────────┐   │
│  │ Pages (Vue Components)              │   │
│  │  - index.vue                        │   │
│  │  - ranking/growth.vue               │   │
│  │  - ranking/dividend.vue             │   │
│  └─────────────────────────────────────┘   │
│  ┌─────────────────────────────────────┐   │
│  │ Components                          │   │
│  │  - RankingTable.vue                 │   │
│  │  - RankingChart.vue                 │   │
│  └─────────────────────────────────────┘   │
└──────────────────┬──────────────────────────┘
                   │ REST API (JSON)
                   ▼
┌─────────────────────────────────────────────┐
│         Go (Gin) API Server                 │
│         ポート: 8080                         │
│  ┌─────────────────────────────────────┐   │
│  │ Handlers (HTTP)                     │   │
│  │  - GetGrowthRanking                 │   │
│  │  - GetDividendRanking               │   │
│  └────────────┬────────────────────────┘   │
│               ▼                             │
│  ┌─────────────────────────────────────┐   │
│  │ Services (Business Logic)           │   │
│  │  - CalculateGrowthRate              │   │
│  │  - GenerateRanking                  │   │
│  └────────────┬────────────────────────┘   │
│               ▼                             │
│  ┌─────────────────────────────────────┐   │
│  │ Repositories (Data Access)          │   │
│  │  - GORM                             │   │
│  └────────────┬────────────────────────┘   │
└───────────────┼─────────────────────────────┘
                │ SQL
                ▼
┌─────────────────────────────────────────────┐
│         MySQL 8.0                           │
│         ポート: 3306                         │
│  ┌─────────────────────────────────────┐   │
│  │ Tables                              │   │
│  │  - companies                        │   │
│  │  - financial_data                   │   │
│  │  - documents                        │   │
│  └─────────────────────────────────────┘   │
└─────────────────────────────────────────────┘

External API:
┌─────────────────────────────────────────────┐
│         EDINET API                          │
│         (金融庁)                             │
└─────────────────────────────────────────────┘
```

### 1.2 データフロー

```
【データ取得フロー】
1. EDINET API → Go Backend → MySQL
   (書類一覧・XBRLダウンロード・財務データ抽出)

【ランキング表示フロー】
2. ブラウザ → Nuxt.js → Go API → MySQL → レスポンス
   (ランキングデータ取得・表示)
```

## 2. ディレクトリ構成

```
edinet-ranking/
├── backend/                           # Go バックエンド
│   ├── main.go                       # エントリーポイント
│   ├── go.mod
│   ├── go.sum
│   ├── .env                          # 環境変数
│   ├── Dockerfile
│   │
│   ├── config/                       # 設定
│   │   ├── config.go                # アプリケーション設定
│   │   └── database.go              # DB接続設定
│   │
│   ├── models/                       # GORMモデル
│   │   ├── company.go               # 企業モデル
│   │   ├── financial_data.go        # 財務データモデル
│   │   ├── document.go              # 書類モデル
│   │   └── ranking.go               # ランキングモデル
│   │
│   ├── handlers/                     # HTTPハンドラー
│   │   ├── ranking_handler.go       # ランキングAPI
│   │   ├── company_handler.go       # 企業API
│   │   └── admin_handler.go         # 管理API
│   │
│   ├── services/                     # ビジネスロジック
│   │   ├── edinet_service.go        # EDINET API連携
│   │   ├── ranking_service.go       # ランキング生成
│   │   └── analysis_service.go      # 財務分析
│   │
│   ├── repositories/                 # データアクセス層
│   │   ├── company_repository.go
│   │   ├── financial_repository.go
│   │   └── document_repository.go
│   │
│   ├── middlewares/                  # ミドルウェア
│   │   ├── cors.go                  # CORS設定
│   │   └── logger.go                # ログ
│   │
│   ├── parser/                       # データ解析
│   │   └── xbrl_parser.go           # XBRL解析
│   │
│   └── utils/                        # ユーティリティ
│       ├── date.go
│       └── validator.go
│
├── frontend/                          # Nuxt.js フロントエンド
│   ├── package.json
│   ├── nuxt.config.ts                # Nuxt設定
│   ├── tsconfig.json                 # TypeScript設定
│   ├── tailwind.config.js            # Tailwind設定
│   ├── Dockerfile
│   │
│   ├── pages/                        # ページ（ファイルベースルーティング）
│   │   ├── index.vue                # トップページ
│   │   └── ranking/
│   │       ├── growth.vue           # 成長率ランキング
│   │       └── dividend.vue         # 配当利回りランキング
│   │
│   ├── components/                   # コンポーネント
│   │   ├── Ranking/
│   │   │   ├── Table.vue           # ランキング表
│   │   │   ├── Chart.vue           # グラフ
│   │   │   └── Card.vue            # カード表示
│   │   ├── Layout/
│   │   │   ├── Header.vue
│   │   │   └── Footer.vue
│   │   └── Common/
│   │       ├── Loading.vue
│   │       └── ErrorMessage.vue
│   │
│   ├── composables/                  # Composition API
│   │   ├── useRanking.ts           # ランキングロジック
│   │   ├── useCompany.ts           # 企業データロジック
│   │   └── useApi.ts               # API通信
│   │
│   ├── types/                        # TypeScript型定義
│   │   ├── ranking.ts
│   │   └── company.ts
│   │
│   ├── assets/                       # 静的アセット
│   │   └── css/
│   │       └── tailwind.css
│   │
│   └── public/                       # 公開ファイル
│       └── favicon.ico
│
├── docker-compose.yml                # Docker Compose設定
├── .gitignore
└── README.md                         # プロジェクト説明
```

## 3. データベース設計

### 3.1 テーブル設計

#### **companies（企業）**
```sql
CREATE TABLE companies (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    sec_code VARCHAR(10) UNIQUE NOT NULL,
    edinet_code VARCHAR(10),
    industry VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_sec_code (sec_code),
    INDEX idx_edinet_code (edinet_code)
);
```

#### **financial_data（財務データ）**
```sql
CREATE TABLE financial_data (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    company_id BIGINT UNSIGNED NOT NULL,
    fiscal_year INT NOT NULL,
    revenue BIGINT,
    operating_income BIGINT,
    net_income BIGINT,
    dividend INT,
    stock_price INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (company_id) REFERENCES companies(id),
    INDEX idx_company_year (company_id, fiscal_year),
    UNIQUE KEY unique_company_year (company_id, fiscal_year)
);
```

#### **documents（書類情報）**
```sql
CREATE TABLE documents (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    company_id BIGINT UNSIGNED NOT NULL,
    doc_id VARCHAR(50) UNIQUE NOT NULL,
    doc_type_code VARCHAR(10),
    fiscal_year INT,
    submit_date_time DATETIME,
    period_start DATE,
    period_end DATE,
    file_path VARCHAR(500),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (company_id) REFERENCES companies(id),
    INDEX idx_doc_id (doc_id),
    INDEX idx_fiscal_year (fiscal_year)
);
```

### 3.2 GORMモデル定義

#### **models/company.go**
```go
package models

import "gorm.io/gorm"

type Company struct {
    gorm.Model
    Name        string         `gorm:"not null" json:"name"`
    SecCode     string         `gorm:"uniqueIndex;not null" json:"secCode"`
    EDINETCode  string         `json:"edinetCode"`
    Industry    string         `json:"industry"`
    FinancialData []FinancialData `gorm:"foreignKey:CompanyID" json:"financialData,omitempty"`
    Documents   []Document     `gorm:"foreignKey:CompanyID" json:"documents,omitempty"`
}
```

#### **models/financial_data.go**
```go
package models

import "gorm.io/gorm"

type FinancialData struct {
    gorm.Model
    CompanyID       uint    `gorm:"not null;index:idx_company_year" json:"companyId"`
    Company         Company `gorm:"foreignKey:CompanyID" json:"company,omitempty"`
    FiscalYear      int     `gorm:"not null;index:idx_company_year" json:"fiscalYear"`
    Revenue         int64   `json:"revenue"`
    OperatingIncome int64   `json:"operatingIncome"`
    NetIncome       int64   `json:"netIncome"`
    Dividend        int     `json:"dividend"`
    StockPrice      int     `json:"stockPrice"`
}
```

#### **models/document.go**
```go
package models

import (
    "time"
    "gorm.io/gorm"
)

type Document struct {
    gorm.Model
    CompanyID      uint      `gorm:"not null" json:"companyId"`
    Company        Company   `gorm:"foreignKey:CompanyID" json:"company,omitempty"`
    DocID          string    `gorm:"uniqueIndex;not null" json:"docId"`
    DocTypeCode    string    `json:"docTypeCode"`
    FiscalYear     int       `json:"fiscalYear"`
    SubmitDateTime time.Time `json:"submitDateTime"`
    PeriodStart    time.Time `json:"periodStart"`
    PeriodEnd      time.Time `json:"periodEnd"`
    FilePath       string    `json:"filePath"`
}
```

#### **models/ranking.go**
```go
package models

// GrowthRankingItem 利益成長率ランキング項目
type GrowthRankingItem struct {
    Rank          int     `json:"rank"`
    Company       Company `json:"company"`
    GrowthRate    float64 `json:"growthRate"`
    PrevIncome    int64   `json:"prevIncome"`
    CurrentIncome int64   `json:"currentIncome"`
}

// DividendRankingItem 配当利回りランキング項目
type DividendRankingItem struct {
    Rank          int     `json:"rank"`
    Company       Company `json:"company"`
    DividendYield float64 `json:"dividendYield"`
    Dividend      int     `json:"dividend"`
    StockPrice    int     `json:"stockPrice"`
}

// Rankings ランキング全体
type Rankings struct {
    Year            int                     `json:"year"`
    Industry        string                  `json:"industry"`
    GrowthRanking   []GrowthRankingItem     `json:"growthRanking,omitempty"`
    DividendRanking []DividendRankingItem   `json:"dividendRanking,omitempty"`
}
```

## 4. 主要な関数設計

### 4.1 Phase 1: データ収集（完了✅）

```go
// api/client.go
func FetchDocumentList(date, apiKey, endpoint, tp string) (*DocumentListResponse, error)
```

### 4.2 Phase 2: データ保存

```go
// storage/csv.go
func SaveToCSV(data interface{}, filename string) error
func LoadFromCSV(filename string) ([]FinancialData, error)

// api/client.go（拡張）
func FetchMultipleDates(dateRange []string, apiKey string) ([]DocumentListResponse, error)
```

### 4.3 Phase 3: 書類ダウンロード

```go
// api/document.go
func DownloadDocument(docID, apiKey string, fileType int) ([]byte, error)
func SaveDocument(docID string, data []byte, outputDir string) error
```

### 4.4 Phase 4: データ抽出

```go
// parser/xbrl.go
func ParseXBRL(data []byte) (*FinancialData, error)
func ExtractNetIncome(xbrl *XBRL) (int64, error)
func ExtractDividend(xbrl *XBRL) (int, error)
```

### 4.5 Phase 5: ランキング生成

```go
// analyzer/growth.go
func CalculateGrowthRate(prev, current int64) float64

// analyzer/ranking.go
func GenerateGrowthRanking(histories []FinancialHistory) []GrowthRankingItem
func GenerateDividendRanking(histories []FinancialHistory) []DividendRankingItem
func SortByGrowthRate(items []GrowthRankingItem) []GrowthRankingItem
```

### 4.6 Phase 6: Web表示

```go
// web/server.go
func StartServer(port int) error

// web/handlers.go
func HandleIndex(w http.ResponseWriter, r *http.Request)
func HandleRanking(w http.ResponseWriter, r *http.Request)

// web/render.go
func RenderRankingPage(rankings Rankings) (string, error)
```

## 5. 処理フロー

### 5.1 全体フロー

```
1. 企業リストの読み込み
   ↓
2. 各企業の書類一覧を取得（過去3年分）
   ↓
3. 有価証券報告書をダウンロード
   ↓
4. XBRLファイルから財務データを抽出
   ↓
5. データをCSVに保存
   ↓
6. ランキングを生成
   ↓
7. WebページにHTMLとして表示
```

### 5.2 Phase 2の詳細フロー

```
1. 日付範囲を生成（過去3年の決算月）
2. 並行処理で書類一覧を取得
   - goroutineで各日付を並行処理
   - channelで結果を収集
3. 有価証券報告書のみフィルタリング
4. 企業情報をCSVに保存
```

### 5.3 Phase 4の詳細フロー

```
1. XBRLファイルを読み込み
2. XMLをパース
3. 必要なタグを検索
   - 純利益: jpcrp_cor:NetIncome
   - 配当金: jpcrp_cor:DividendPaidPerShare
4. データを構造体に格納
5. CSVに保存
```

### 5.4 Phase 5の詳細フロー

```
【利益成長率ランキング】
1. 各企業の2022年、2023年の純利益を取得
2. 成長率を計算: (2023 - 2022) / 2022 × 100
3. 成長率でソート（降順）
4. 上位10社を抽出

【配当利回りランキング】
1. 各企業の2023年配当金と株価を取得
2. 利回りを計算: 配当金 / 株価 × 100
3. 利回りでソート（降順）
4. 上位10社を抽出
```

## 6. API連携仕様

### 6.1 書類一覧API

**リクエスト例:**
```
GET https://api.edinet-fsa.go.jp/api/v2/documents.json?date=2023-06-30&type=2&Subscription-Key=xxx
```

**レスポンス（抜粋）:**
```json
{
  "metadata": { "status": "200", "resultset": { "count": 182 } },
  "results": [
    {
      "docID": "S100XXXX",
      "filerName": "ソニーグループ株式会社",
      "secCode": "6758",
      "docTypeCode": "120",
      "submitDateTime": "2023-06-30 15:00"
    }
  ]
}
```

### 6.2 書類取得API

**リクエスト例:**
```
GET https://api.edinet-fsa.go.jp/api/v2/documents/S100XXXX?type=1&Subscription-Key=xxx
```

**レスポンス:**
- ZIPファイル（XBRL含む）

## 7. XBRLデータ構造（例）

```xml
<jpcrp_cor:NetIncome contextRef="CurrentYearInstant" unitRef="JPY">
    4900000000000
</jpcrp_cor:NetIncome>

<jpcrp_cor:DividendPaidPerShare contextRef="CurrentYearInstant" unitRef="JPYPerShare">
    300
</jpcrp_cor:DividendPaidPerShare>
```

## 8. フロントエンド実装例（Nuxt.js）

### 8.1 ページコンポーネント

#### **pages/index.vue**
```vue
<template>
  <div class="container mx-auto px-4 py-8">
    <h1 class="text-4xl font-bold text-center mb-8">
      IT・通信業界 財務ランキング {{ currentYear }}
    </h1>

    <!-- 利益成長率ランキング -->
    <section class="mb-12">
      <h2 class="text-2xl font-bold mb-4">
        💹 利益成長率ランキング
      </h2>
      <RankingChart v-if="growthData" :data="growthData.rankings" type="growth" />
      <RankingTable v-if="growthData" :rankings="growthData.rankings" type="growth" />
    </section>

    <!-- 配当利回りランキング -->
    <section class="mb-12">
      <h2 class="text-2xl font-bold mb-4">
        💰 配当利回りランキング
      </h2>
      <RankingChart v-if="dividendData" :data="dividendData.rankings" type="dividend" />
      <RankingTable v-if="dividendData" :rankings="dividendData.rankings" type="dividend" />
    </section>
  </div>
</template>

<script setup lang="ts">
const config = useRuntimeConfig()
const currentYear = new Date().getFullYear()

// APIから利益成長率ランキングを取得
const { data: growthData, pending: growthPending, error: growthError } = 
  await useFetch(`${config.public.apiBase}/api/v1/rankings/growth`)

// APIから配当利回りランキングを取得
const { data: dividendData, pending: dividendPending, error: dividendError } = 
  await useFetch(`${config.public.apiBase}/api/v1/rankings/dividend`)
</script>
```

### 8.2 コンポーネント

#### **components/Ranking/Table.vue**
```vue
<template>
  <div class="bg-white rounded-lg shadow-lg overflow-hidden">
    <table class="min-w-full">
      <thead class="bg-gradient-to-r from-blue-500 to-blue-600 text-white">
        <tr>
          <th class="px-6 py-3 text-left text-sm font-semibold">順位</th>
          <th class="px-6 py-3 text-left text-sm font-semibold">企業名</th>
          <th class="px-6 py-3 text-left text-sm font-semibold">証券コード</th>
          <th class="px-6 py-3 text-right text-sm font-semibold">
            {{ type === 'growth' ? '成長率' : '配当利回り' }}
          </th>
        </tr>
      </thead>
      <tbody class="divide-y divide-gray-200">
        <tr v-for="item in rankings" 
            :key="item.rank"
            class="hover:bg-gray-50 transition-colors">
          <td class="px-6 py-4">
            <span class="flex items-center">
              <span v-if="item.rank <= 3" class="mr-2 text-2xl">
                {{ getMedal(item.rank) }}
              </span>
              <span class="font-semibold">{{ item.rank }}</span>
            </span>
          </td>
          <td class="px-6 py-4 font-medium text-gray-900">
            {{ item.company.name }}
          </td>
          <td class="px-6 py-4 text-gray-600">
            {{ item.company.secCode }}
          </td>
          <td class="px-6 py-4 text-right">
            <span :class="getRateClass(item)" class="font-bold text-lg">
              {{ formatRate(item) }}
            </span>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
import type { GrowthRankingItem, DividendRankingItem } from '~/types/ranking'

const props = defineProps<{
  rankings: (GrowthRankingItem | DividendRankingItem)[]
  type: 'growth' | 'dividend'
}>()

const getMedal = (rank: number): string => {
  const medals = { 1: '🥇', 2: '🥈', 3: '🥉' }
  return medals[rank as keyof typeof medals] || ''
}

const formatRate = (item: any): string => {
  const value = item.growthRate || item.dividendYield
  return `${value >= 0 ? '+' : ''}${value.toFixed(1)}%`
}

const getRateClass = (item: any): string => {
  const value = item.growthRate || item.dividendYield
  if (value >= 50) return 'text-green-600'
  if (value >= 30) return 'text-green-500'
  if (value >= 10) return 'text-blue-500'
  return 'text-gray-700'
}
</script>
```

#### **components/Ranking/Chart.vue**
```vue
<template>
  <div class="bg-white rounded-lg shadow-lg p-6 mb-6">
    <Bar :data="chartData" :options="chartOptions" class="h-96" />
  </div>
</template>

<script setup lang="ts">
import { Bar } from 'vue-chartjs'
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  BarElement,
  CategoryScale,
  LinearScale,
  type ChartOptions
} from 'chart.js'

ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale)

const props = defineProps<{
  data: any[]
  type: 'growth' | 'dividend'
}>()

const chartData = computed(() => ({
  labels: props.data.map(item => item.company.name),
  datasets: [{
    label: props.type === 'growth' ? '成長率 (%)' : '配当利回り (%)',
    data: props.data.map(item => item.growthRate || item.dividendYield),
    backgroundColor: props.data.map((_, index) => {
      if (index === 0) return 'rgba(34, 197, 94, 0.8)'  // 1位: 緑
      if (index === 1) return 'rgba(59, 130, 246, 0.8)' // 2位: 青
      if (index === 2) return 'rgba(251, 146, 60, 0.8)' // 3位: オレンジ
      return 'rgba(156, 163, 175, 0.6)'                  // その他: グレー
    }),
    borderColor: props.data.map((_, index) => {
      if (index === 0) return 'rgb(34, 197, 94)'
      if (index === 1) return 'rgb(59, 130, 246)'
      if (index === 2) return 'rgb(251, 146, 60)'
      return 'rgb(156, 163, 175)'
    }),
    borderWidth: 2
  }]
}))

const chartOptions: ChartOptions<'bar'> = {
  responsive: true,
  maintainAspectRatio: false,
  plugins: {
    legend: {
      display: false
    },
    tooltip: {
      callbacks: {
        label: (context) => {
          return `${context.parsed.y.toFixed(1)}%`
        }
      }
    }
  },
  scales: {
    y: {
      beginAtZero: true,
      ticks: {
        callback: (value) => `${value}%`
      }
    }
  }
}
</script>
```

### 8.3 Composables（ロジック）

#### **composables/useApi.ts**
```typescript
export const useApi = () => {
  const config = useRuntimeConfig()
  const baseURL = config.public.apiBase

  const fetchGrowthRanking = async () => {
    return await $fetch(`${baseURL}/api/v1/rankings/growth`)
  }

  const fetchDividendRanking = async () => {
    return await $fetch(`${baseURL}/api/v1/rankings/dividend`)
  }

  const fetchCompanies = async () => {
    return await $fetch(`${baseURL}/api/v1/companies`)
  }

  return {
    fetchGrowthRanking,
    fetchDividendRanking,
    fetchCompanies
  }
}
```

### 8.4 型定義

#### **types/ranking.ts**
```typescript
export interface Company {
  id: number
  name: string
  secCode: string
  edinetCode?: string
  industry?: string
}

export interface GrowthRankingItem {
  rank: number
  company: Company
  growthRate: number
  prevIncome: number
  currentIncome: number
}

export interface DividendRankingItem {
  rank: number
  company: Company
  dividendYield: number
  dividend: number
  stockPrice: number
}

export interface Rankings {
  year: number
  industry: string
  growthRanking?: GrowthRankingItem[]
  dividendRanking?: DividendRankingItem[]
}
```

### 8.5 Nuxt設定

#### **nuxt.config.ts**
```typescript
export default defineNuxtConfig({
  devtools: { enabled: true },
  
  modules: [
    '@nuxtjs/tailwindcss'
  ],

  runtimeConfig: {
    public: {
      apiBase: process.env.API_BASE_URL || 'http://localhost:8080'
    }
  },

  app: {
    head: {
      title: 'EDINET 財務ランキング',
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' }
      ]
    }
  }
})
```

### 8.1 トップページ（index.html）

```html
<!DOCTYPE html>
<html>
<head>
    <title>IT・通信業界 財務ランキング</title>
    <link rel="stylesheet" href="/static/css/style.css">
    <script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
</head>
<body>
    <header>
        <h1>IT・通信業界 財務ランキング 2023</h1>
    </header>
    
    <main>
        <section class="ranking">
            <h2>利益成長率ランキング</h2>
            <canvas id="growthChart"></canvas>
            <table>
                <!-- ランキングテーブル -->
            </table>
        </section>
        
        <section class="ranking">
            <h2>配当利回りランキング</h2>
            <canvas id="dividendChart"></canvas>
            <table>
                <!-- ランキングテーブル -->
            </table>
        </section>
    </main>
    
    <script src="/static/js/chart-config.js"></script>
</body>
</html>
```

## 9. エラーハンドリング方針

### 9.1 API通信エラー

| エラー | 対応 |
|--------|------|
| ネットワークエラー | リトライ（最大3回） |
| 401 Unauthorized | APIキーの確認を促す |
| 404 Not Found | 書類が見つからない旨を記録 |
| 429 Too Many Requests | 待機してリトライ |

### 9.2 データ欠損

| ケース | 対応 |
|--------|------|
| 純利益データなし | ランキングから除外 |
| 配当金データなし | 配当利回りランキングから除外 |
| 株価データなし | 手動入力または除外 |

## 10. パフォーマンス最適化

### 10.1 並行処理

```go
// 複数企業のデータを並行取得
func FetchAllCompanies(companies []Company) []FinancialHistory {
    results := make(chan FinancialHistory, len(companies))
    
    for _, company := range companies {
        go func(c Company) {
            data := fetchCompanyData(c)
            results <- data
        }(company)
    }
    
    // 結果を収集
    var histories []FinancialHistory
    for range companies {
        histories = append(histories, <-results)
    }
    return histories
}
```

### 10.2 キャッシュ

```go
// 一度取得したデータはキャッシュ
type Cache struct {
    data map[string]FinancialData
    mu   sync.RWMutex
}
```

## 11. テスト方針

### 11.1 単体テスト

```go
// analyzer/growth_test.go
func TestCalculateGrowthRate(t *testing.T) {
    tests := []struct {
        prev, current int64
        expected      float64
    }{
        {100, 150, 50.0},
        {100, 75, -25.0},
    }
    // ...
}
```

### 11.2 統合テスト

- APIとの実際の通信テスト
- エンドツーエンドのデータフロー検証

## 12. セキュリティ

- APIキーは環境変数で管理
- `.env`ファイルは`.gitignore`に追加
- HTMLエスケープ処理

## 13. 今後の拡張ポイント

- データベース導入（PostgreSQL/SQLite）
- 認証機能（プライベート運用の場合）
- API化（JSON形式でデータ提供）
- グラフの種類追加（円グラフ、レーダーチャートなど）
- 業界比較機能
