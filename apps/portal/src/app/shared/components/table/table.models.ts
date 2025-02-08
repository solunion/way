export interface TableColumns {
  id: string;
  name: string;
  type: TableColumnType;
  default?: string;
}

export enum TableColumnType {
  STRING = 'string',
  NUMBER = 'number',
  DATE = 'date',
  DATETIME = 'datetime',
  BOOLEAN = 'boolean',
  OBJECT = 'object',
  ARRAY = 'array',
  CUSTOM = 'custom'
}
