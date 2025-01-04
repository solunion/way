import { Prisma, Rule } from '@prisma/client';

export class RuleEntity implements Rule {
  id!: string;
  name!: string;
  type!: string;
  value!: Prisma.JsonValue;
  tenantId!: string | null;
  createdAt!: Date;
  updatedAt!: Date;
  deletedAt!: Date | null;
}
