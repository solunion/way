import { Tenant } from '@prisma/client';

export class TenantEntity implements Tenant {
  name: string;
  id: string;
  description: string | null;
  createdAt: Date;
  updatedAt: Date;
  deletedAt: Date | null;

  constructor(name: string, id: string, description: string | null, createdAt: Date, updatedAt: Date, deletedAt: Date | null) {
    this.name = name;
    this.id = id;
    this.description = description;
    this.createdAt = createdAt;
    this.updatedAt = updatedAt;
    this.deletedAt = deletedAt;
  }
}
