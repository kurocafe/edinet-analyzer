import type {RevenueGrowthRanking} from "~/types/ranking"

export default defineEventHandler(() => {
  const mock: RevenueGrowthRanking[] = [
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

  return mock
})