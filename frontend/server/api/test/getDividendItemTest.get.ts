export default defineEventHandler(() => {
  //dividend GET api test
  //need to verify if type is correct from backend API with zod
  return {
    Rank: 1,
    Company: {
      Name: "CuriousForge",
      SecCode: "shit",
      EDINETCode: "dogwater",
      Industry: "CuriousForge",
      FinancialData: [],
      Documents: []
    },
    DividendYield: 2.3,
    StockPrice: 1000
  }
})