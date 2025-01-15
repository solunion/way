import { Application } from '@prisma/client';

export class ApplicationEntity implements Application {
  id: string;
  name: string;
  description: string | null;
  tenantId: string | null;
  createdAt: Date;
  updatedAt: Date;
  deletedAt: Date | null;


  constructor(id: string, name: string, description: string | null, tenantId: string | null, createdAt: Date, updatedAt: Date, deletedAt: Date | null) {
    this.id = id;
    this.name = name;
    this.description = description;
    this.tenantId = tenantId;
    this.createdAt = createdAt;
    this.updatedAt = updatedAt;
    this.deletedAt = deletedAt;
  }
}
