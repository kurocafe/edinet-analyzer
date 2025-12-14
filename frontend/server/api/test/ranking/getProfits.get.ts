import type { ProfitGrowthRanking } from "~/types/ranking"

export default defineEventHandler(() => {
  const mock: ProfitGrowthRanking[] = [
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

  return mock
})