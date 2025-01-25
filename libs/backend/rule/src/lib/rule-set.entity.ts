import type { RuleSet } from '@prisma/client';

export class RuleSetEntity implements RuleSet {
  id: string;
  name: string;
  tenantId: string | null;
  parentId: string | null;
  createdAt: Date;
  updatedAt: Date;
  deletedAt: Date | null;

  constructor(id: string, name: string, tenantId: string | null, parentId: string | null, createdAt: Date, updatedAt: Date, deletedAt: Date | null) {
    this.id = id;
    this.name = name;
    this.tenantId = tenantId;
    this.parentId = parentId;
    this.createdAt = createdAt;
    this.updatedAt = updatedAt;
    this.deletedAt = deletedAt;
  }
}
