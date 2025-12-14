export interface FinancialData {
  ID: number
  companyId: number
  fiscalYear: number
  revenue: number          // 売上高（円）
  operatingIncome: number  // 営業利益（円）
  netIncome: number        // 純利益（円）
  dividend: number         // 配当（円/株）
  company?: Company
}