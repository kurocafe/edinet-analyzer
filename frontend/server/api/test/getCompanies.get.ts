import type { Company } from "~/types/company"

export default defineEventHandler(() => {
  const mock: Company[] = [
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

  return mock
})