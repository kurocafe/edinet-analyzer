import type { FinancialData } from "~/types/financialData"

export const useFin = () => {
  const financialData = useState<FinancialData[]>("fin:data", () => [])
  const pending = useState("fin:pending", () => false)
  const error = useState<Error | null>("fin:error", () => null)

  const fetchFinancial = async () => {
    pending.value = true
    error.value = null

    try{
      financialData.value = await $fetch("/api/financial")
    }catch(e){
      error.value = e as Error
    }finally{
      pending.value = false
    }
  }

  const getFinancialByYear = (year: number) => {
    return computed(() => financialData.value.filter(f => f.fiscalYear == year))
  }

  const getFinancialById = (id: number) => {
    return computed(() => financialData.value.filter(f => f.companyId == id))
  }

  return {
    financialData,
    pending,
    error,
    fetchFinancial,
    getFinancialById,
    getFinancialByYear
  }
}