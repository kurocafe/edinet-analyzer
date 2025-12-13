import type { Company } from "~/types/company"

export const useCompany = () => {
  const companies = useState<Company[]>("companies", () => [])
  
}