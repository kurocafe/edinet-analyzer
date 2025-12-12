export interface Company {
  Name: string;
  SecCode: string;
  EDINETCode: string;
  Industry: string;
  FinancialData: FinancialData[];
  Documents: Document[];
}