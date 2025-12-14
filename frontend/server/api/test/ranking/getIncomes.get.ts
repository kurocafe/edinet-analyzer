import type { IncomeGrowthRanking } from "~/types/ranking"

export default defineEventHandler(() => {
  const mock: IncomeGrowthRanking[] = [
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
})