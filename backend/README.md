# EDINET Analyzer - Backend API

Go + Gin + GORM + MySQLで構築したEDINET財務データ分析API

## 技術スタック

- **Go**: 1.25.4
- **Webフレームワーク**: Gin
- **ORM**: GORM
- **データベース**: MySQL 8.0
- **その他**: godotenv (環境変数管理)

## ディレクトリ構成

```
backend/
├── main.go                      # エントリーポイント、APIルーティング
├── go.mod                       # 依存関係管理
├── config/
│   └── database.go              # DB接続設定（リトライ機能付き）
├── models/
│   ├── company.go               # 企業モデル
│   └── financial_data.go        # 財務データモデル
├── services/
│   ├── edinet.go                # EDINET API連携（書類一覧、ダウンロード）
│   └── xbrl_parser.go           # XBRLパーサー（ZIP→財務データ）
├── scripts/
│   ├── seed_companies.go        # 企業データCSVインポート
│   ├── fetch_all_companies.go   # 全企業データ自動取得スクリプト
│   └── test_xbrl.go             # XBRL解析テストスクリプト
└── (以下は空ディレクトリ)
    ├── handlers/                # HTTPハンドラー（未実装）
    ├── repositories/            # データアクセス層（未実装）
    ├── middlewares/             # ミドルウェア（未実装）
    └── utils/                   # ユーティリティ（未実装）
```

## セットアップ

### 1. 環境変数設定

`backend/.env` ファイルを作成：

```env
# データベース設定
DB_HOST=mysql
DB_PORT=3306
DB_USER=analyzer_user
DB_PASSWORD=analyzer_password
DB_NAME=edinet_analyzer_db

# EDINET API設定
EDINET_API_KEY=あなたのAPIキー
```

EDINET APIキーは[こちら](https://disclosure2.edinet-fsa.go.jp/)から取得してください。

### 2. 起動（Docker Compose推奨）

```bash
# プロジェクトルートで実行
docker-compose up
```

### 3. 企業データのインポート（初回のみ）

```bash
cd backend/scripts
go run seed_companies.go
```

### 4. MySQLに直接接続（デバッグ用）

```bash
docker exec -it edinet-mysql mysql -u analyzer_user -p --default-character-set=utf8mb4
# パスワード: analyzer_password
```

## 実装済みAPI

### 基本API

| メソッド | エンドポイント | 説明 | クエリパラメータ |
|---------|---------------|------|-----------------|
| GET | `/` | ヘルスチェック | - |
| GET | `/api/v1/companies` | 企業一覧取得 | - |
| POST | `/api/v1/companies` | 企業追加 | - |
| GET | `/api/v1/financial-data` | 財務データ取得 | `fiscalYear`, `companyId` |
| POST | `/api/v1/financial-data` | 財務データ追加 | - |
| GET | `/api/v1/companies/:id/financial-data` | 特定企業の財務データ | - |

### EDINET連携API

| メソッド | エンドポイント | 説明 | パラメータ |
|---------|---------------|------|-----------|
| GET | `/api/v1/edinet/documents` | EDINET書類一覧取得 | `date` (必須), `secCode` (任意) |
| GET | `/api/v1/edinet/documents/:docID/download` | 書類ZIPダウンロード | `docID` (パス) |
| POST | `/api/v1/edinet/documents/:docID/parse` | XBRL解析→DB保存 | `docID` (パス), `companyID` (JSON) |

### APIリクエスト例

#### 基本API

```bash
# 企業一覧取得
curl http://localhost:8080/api/v1/companies

# 企業追加
curl -X POST http://localhost:8080/api/v1/companies \
  -H "Content-Type: application/json" \
  -d '{"name":"ソニーグループ","secCode":"6758"}'

# 財務データ取得（2023年度のみ）
curl http://localhost:8080/api/v1/financial-data?fiscalYear=2023

# 財務データ追加
curl -X POST http://localhost:8080/api/v1/financial-data \
  -H "Content-Type: application/json" \
  -d '{
    "companyId": 1,
    "fiscalYear": 2023,
    "revenue": 10000000000000,
    "operatingIncome": 1500000000000,
    "netIncome": 1400000000000,
    "dividend": 120
  }'
```

#### EDINET連携API

```bash
# EDINET書類一覧取得（全件）
curl "http://localhost:8080/api/v1/edinet/documents?date=2024-06-28"

# EDINET書類一覧取得（証券コード指定）
curl "http://localhost:8080/api/v1/edinet/documents?date=2024-06-28&secCode=5973"

# レスポンス例
# [{
#   "docID": "S100TUID",
#   "edinetCode": "E01441",
#   "secCode": "59730",
#   "filerName": "株式会社トーアミ",
#   "docTypeCode": "120",
#   "periodStart": "2023-04-01",
#   "periodEnd": "2024-03-31",
#   "submitDateTime": "2024-06-28 09:00",
#   "xbrlFlag": "1"
# }]

# 書類ZIPダウンロード
curl "http://localhost:8080/api/v1/edinet/documents/S100TX1S/download"

# レスポンス例
# {
#   "message": "ダウンロード完了",
#   "filename": "/app/data/edinet/S100TX1S.zip",
#   "docID": "S100TX1S"
# }

# XBRL解析してDBに保存
curl -X POST "http://localhost:8080/api/v1/edinet/documents/S100TX1S/parse" \
  -H "Content-Type: application/json" \
  -d '{"companyID": 1}'

# レスポンス例
# {
#   "message": "財務データを保存しました",
#   "financialData": {
#     "ID": 1,
#     "companyId": 1,
#     "fiscalYear": 2024,
#     "revenue": 520800000000,
#     "operatingIncome": 62500000000,
#     "netIncome": 43200000000,
#     "dividend": 125
#   }
# }
```

## スクリプト

### 全企業データ自動取得

15社の企業データを自動でダウンロード・解析・DB保存します：

```bash
docker exec -it analyzer-backend go run /app/scripts/fetch_all_companies.go
```

**処理内容:**
1. 15社の証券コードで書類を検索（複数日付を自動試行）
2. 有価証券報告書をダウンロード（ZIP形式）
3. XBRLファイルを解析して財務データ抽出
4. データベースに自動保存

### XBRL解析テスト

単一のZIPファイルをテスト解析します：

```bash
docker exec -it analyzer-backend go run /app/scripts/test_xbrl.go
```

## 未実装API（要実装）

- `GET /api/v1/companies/:id` - 企業詳細取得
- `GET /api/v1/rankings/revenue-growth` - 売上成長率ランキング
- `GET /api/v1/rankings/profit-growth` - 営業利益成長率ランキング
- `GET /api/v1/rankings/income-growth` - 純利益成長率ランキング
- `GET /api/v1/rankings/dividend` - 配当金額ランキング

詳細は `/files/API_SPEC.md` を参照

## データベース構造

### companies（企業）

| カラム | 型 | 説明 |
|--------|---|------|
| id | BIGINT | 主キー（自動採番） |
| name | VARCHAR(255) | 企業名 |
| sec_code | VARCHAR(10) | 証券コード（ユニーク） |
| edinet_code | VARCHAR(10) | EDINETコード |
| industry | VARCHAR(100) | 業種 |

### financial_data（財務データ）

| カラム | 型 | 説明 |
|--------|---|------|
| id | BIGINT | 主キー（自動採番） |
| company_id | BIGINT | 企業ID（外部キー） |
| fiscal_year | INT | 会計年度（例: 2023） |
| revenue | BIGINT | 売上高（円） |
| operating_income | BIGINT | 営業利益（円） |
| net_income | BIGINT | 純利益（円） |
| dividend | INT | 配当金（円/株） |

## 開発状況

### 実装済み
- ✅ 基本的なCRUD API（企業、財務データ）
- ✅ CORS設定（フロントエンド連携対応）
- ✅ データベース接続（リトライ機能付き）
- ✅ GORMモデル定義
- ✅ 自動マイグレーション
- ✅ **EDINET書類一覧取得API**（証券コードフィルタリング対応）
- ✅ **EDINET書類ダウンロードAPI**（ZIP形式）
- ✅ **XBRLパーサー実装**（ZIP→財務データ抽出）
- ✅ **XBRL解析API**（解析結果をDB保存）
- ✅ **全企業データ自動取得スクリプト**
- ✅ **services層の部分的実装**（edinet.go, xbrl_parser.go）

### 未実装（TODO）
- ❌ レイヤー構造（handlers/repositories）への完全分割
- ❌ ランキングAPI（成長率計算ロジック）
- ❌ 企業詳細取得API
- ❌ エラーハンドリングの強化
- ❌ バリデーション
- ❌ ログ出力の整備
- ❌ XBRL解析精度の向上（企業ごとのタグ名の違いに対応）

## 注意事項

- APIハンドラーは`main.go`に直書きされています（リファクタリング推奨）
- 設計書で定義されたレイヤー構造は部分的実装（services/のみ）
- ランキング機能はまだ実装されていません
- **XBRL解析は正規表現ベース**のため、企業によってタグ名が異なる場合は抽出できない可能性があります
- ダウンロードしたZIPファイルは `/app/data/edinet/` に保存されます
- `fetch_all_companies.go` は15社のIT・通信企業のデータを取得します

## XBRLパーサーの仕様

**抽出する財務データ:**
- 会計年度（ファイル名から推測）
- 売上高（NetSalesSummaryOfBusinessResults）
- 営業利益（OperatingIncome）
- 純利益（ProfitLossAttributableToOwnersOfParentSummaryOfBusinessResults）
- 配当金（DividendPaidPerShareSummaryOfBusinessResults）

**対応フォーマット:**
- PublicDoc配下のXBRLファイル（.xbrl）
- contextRef="CurrentYearDuration" の値を抽出

## 関連ドキュメント

- API仕様書: `/files/API_SPEC.md`
- 設計書: `/files/design.md`
- 要件定義: `/files/requirements.md`
