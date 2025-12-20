import type { RankingItem } from "~/types/ranking"

interface searchQuery {
  limit?: number;
  baseYear?: number;
  targetYear?: number;
}

export const useRanking = () => {
  const revenues = useState<RankingItem[]>("rank:revenues", () => [])
  const profits = useState<RankingItem[]>("rank:profits", () => [])
  const incomes = useState<RankingItem[]>("rank:incomes", () => [])
  
  const pending = useState("rank:pending", () => false)
  const error = useState<Error | null>("rank:error", () => null)

  const fetchRevenues = async () => {
    pending.value = true
    error.value = null
    let url = "/api/rankings/revenue-growth"

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
      profits.value = await $fetch("/api/rankings/profit-growth")
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
      incomes.value = await $fetch("/api/rankings/income-growth")
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