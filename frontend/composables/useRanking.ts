import type { RevenueGrowthRanking, ProfitGrowthRanking, IncomeGrowthRanking } from "~/types/ranking"

interface searchQuery {
  limit?: number;
  baseYear?: number;
  targetYear?: number;
}

export const useRanking = () => {
  const revenues = useState<RevenueGrowthRanking[]>("rank:revenues", () => [])
  const profits = useState<ProfitGrowthRanking[]>("rank:profits", () => [])
  const incomes = useState<IncomeGrowthRanking[]>("rank:incomes", () => [])
  
  const pending = useState("rank:pending", () => false)
  const error = useState<Error | null>("rank:error", () => null)

  const fetchRevenues = async ({limit, baseYear, targetYear}: searchQuery = {}) => {
    pending.value = true
    error.value = null
    let url = "/api/test/ranking/getRevenues?"
    if(limit) url += `limit=${limit}`
    if(baseYear) url += `baseYear=${baseYear}`
    if(targetYear) url += `targetYear=${targetYear}`

    try{
      revenues.value = await $fetch(url)
    }catch(e){
      error.value = e as Error
    }finally{
      pending.value = false
    }
  }

  const fetchProfits = async () => {
    pending.value = true
    error.value = null

    try{
      profits.value = await $fetch("/api/test/ranking/getProfits")
    }catch(e){
      error.value = e as Error
    }finally{
      pending.value = false
    }
  }

  const fetchIncomes = async() => {
    pending.value = true
    error.value = null

    try{
      incomes.value = await $fetch("/api/test/ranking/getIncomes")
    }catch(e){
      error.value = e as Error
    }finally{
      pending.value = false
    }
  }


  return {
    revenues,
    profits,
    incomes,
    pending,
    error,
    fetchRevenues,
    fetchProfits,
    fetchIncomes
  }
}