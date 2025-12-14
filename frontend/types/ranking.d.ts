export interface Rankings{
  Year: number;
  Industry: string;
  GrowthRanking: GrowthRankingItem[];
  DividendRanking: DividendRankingItem[];
}

export interface RevenueGrowthRanking {
  rank: number
  company: Company
  baseYear: number
  targetYear: number
  baseRevenue: number
  targetRevenue: number
  growthRate: number       // パーセント
}

export interface IncomeGrowthRanking {
  rank: number
  company: Company
  baseYear: number
  targetYear: number
  baseIncome: number
  targetIncome: number
  growthRate: number 
}

export interface DividendRanking {
  rank: number
  company: Company
  fiscalYear: number
  dividend: number
}