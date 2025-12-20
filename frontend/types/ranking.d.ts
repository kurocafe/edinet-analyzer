export interface Rankings{
  Year: number;
  Industry: string;
  GrowthRanking: GrowthRankingItem[];
  DividendRanking: DividendRankingItem[];
}

export interface RankingItem {
  rank: number
  companyId: number
  companyName: string
  secCode: string
  currentValue: number
  previousValue: number
  growthRate: number
  currentYear: number
  previousYear: number
}

