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
├── main.go                 # エントリーポイント、APIルーティング
├── go.mod                  # 依存関係管理
├── config/
│   └── database.go         # DB接続設定（リトライ機能付き）
├── models/
│   ├── company.go          # 企業モデル
│   └── financial_data.go   # 財務データモデル
├── scripts/
│   └── seed_companies.go   # 企業データCSVインポート
└── (以下は空ディレクトリ)
    ├── handlers/           # HTTPハンドラー（未実装）
    ├── repositories/       # データアクセス層（未実装）
    ├── services/           # ビジネスロジック（未実装）
    ├── middlewares/        # ミドルウェア（未実装）
    ├── parser/             # XBRLパーサー（未実装）
    └── utils/              # ユーティリティ（未実装）
```

## セットアップ

### 1. 環境変数設定

`backend/.env` ファイルを作成：

```env
DB_HOST=mysql
DB_PORT=3306
DB_USER=analyzer_user
DB_PASSWORD=analyzer_password
DB_NAME=edinet_analyzer_db
```

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

| メソッド | エンドポイント | 説明 | クエリパラメータ |
|---------|---------------|------|-----------------|
| GET | `/` | ヘルスチェック | - |
| GET | `/api/v1/companies` | 企業一覧取得 | - |
| POST | `/api/v1/companies` | 企業追加 | - |
| GET | `/api/v1/financial-data` | 財務データ取得 | `fiscalYear`, `companyId` |
| POST | `/api/v1/financial-data` | 財務データ追加 | - |
| GET | `/api/v1/companies/:id/financial-data` | 特定企業の財務データ | - |

### APIリクエスト例

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

### 未実装（TODO）
- ❌ レイヤー構造（handlers/services/repositories）への分割
- ❌ ランキングAPI（成長率計算ロジック）
- ❌ EDINET API連携
- ❌ XBRLパーサー
- ❌ エラーハンドリングの強化
- ❌ バリデーション
- ❌ ログ出力の整備

## 注意事項

- 現在のコードはすべて`main.go`に直書きされています
- 設計書で定義されたレイヤー構造（handlers/services/repositories）は未実装です
- ランキング機能はまだ実装されていません
- EDINET APIとの連携機能はありません

## 関連ドキュメント

- API仕様書: `/files/API_SPEC.md`
- 設計書: `/files/design.md`
- 要件定義: `/files/requirements.md`
