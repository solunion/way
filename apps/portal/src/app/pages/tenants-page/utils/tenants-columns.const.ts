import { TableColumns, TableColumnType } from '../../../shared/components/table/table.models';

export const TENANTS_COLUMS: TableColumns[] = [
  {
    id: 'name',
    name: 'Name',
    type: TableColumnType.STRING,
  },
  {
    id: 'description',
    name: 'Description',
    type: TableColumnType.STRING,
  }
]
