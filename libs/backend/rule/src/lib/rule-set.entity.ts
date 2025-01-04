import { RuleSet } from '@prisma/client';

export class RuleSetEntity implements RuleSet {
  id!: string;
  name!: string;
  tenantId!: string | null;
  parentId!: string | null;
  createdAt!: Date;
  updatedAt!: Date;
  deletedAt!: Date | null;
}
