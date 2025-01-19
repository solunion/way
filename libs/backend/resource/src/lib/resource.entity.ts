import type { Resource, ResourceType } from '@prisma/client';

export class ResourceEntity implements Resource {
  id: string;
  name: string;
  type: ResourceType;
  componentId: string;
  tenantId: string | null;
  createdAt: Date;
  updatedAt: Date;
  deletedAt: Date | null;

  constructor(
    id: string,
    name: string,
    type: ResourceType,
    componentId: string,
    tenantId: string | null,
    createdAt: Date,
    updatedAt: Date,
    deletedAt: Date | null
  ) {
    this.id = id;
    this.name = name;
    this.type = type;
    this.componentId = componentId;
    this.tenantId = tenantId;
    this.createdAt = createdAt;
    this.updatedAt = updatedAt;
    this.deletedAt = deletedAt;
  }
} 