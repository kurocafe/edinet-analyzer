import type { Company } from "~/types/company"

export const useCompany = () => {
  const companies = useState<Company[]>("companies", () => [])
  const pending = useState("company_pending", () => false)
  const error = useState<Error | null>("company_error", () => null)
  
  const fetchCompany = async () => {
    error.value = null
    pending.value = true

    try{
      companies.value = await $fetch("/api/test/getCompanies")
    }catch(e){
      error.value = e as Error
    }finally{
      pending.value = false
    }
  }

  const getCompanyById = (id: number) => {
    
    return computed(() =>
      companies.value.find(c => c.ID === id)
    )
    
  }

  return {
    companies,
    pending, 
    error,
    fetchCompany,
    getCompanyById
  }
}