export interface Company {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  name: string;
  secCode: string;
  edinetCode: string;
  financialData?: FinancialData[];
}