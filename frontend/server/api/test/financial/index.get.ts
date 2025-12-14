import type {FinancialData} from "~/types/financialData"

export default defineEventHandler(() => {
  const datas: FinancialData[] = [
    {
    "ID": 1,
    "companyId": 1,
    "fiscalYear": 2023,
    "revenue": 10000000000000,
    "operatingIncome": 1500000000000,
    "netIncome": 1400000000000,
    "dividend": 120,
    }
  ]

  return datas
})