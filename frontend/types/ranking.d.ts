export interface Rankings{
  Year: number;
  Industry: string;
  GrowthRanking: GrowthRankingItem[];
  DividendRanking: DividendRankingItem[];
}

export interface GrowthRankingItem{  
  Rank: number;
  Company: Company;
  GrowthRate: number;
  PrevIncome: number;
  CurrentIncome: number;
}

export interface DividendRankingItem {
  Rank: number;
  Company: Company;
  DividendYield: number;
  Dividend: number;
  StockPrice: number;
}