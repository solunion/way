import type { Component } from '@prisma/client';

export class ComponentEntity implements Component {
  id: string;
  name: string;
  description: string | null;
  applicationId: string;
  tenantId: string | null;
  createdAt: Date;
  updatedAt: Date;
  deletedAt: Date | null;

  constructor(
    id: string,
    name: string,
    description: string | null,
    applicationId: string,
    tenantId: string | null,
    createdAt: Date,
    updatedAt: Date,
    deletedAt: Date | null
  ) {
    this.id = id;
    this.name = name;
    this.description = description;
    this.applicationId = applicationId;
    this.tenantId = tenantId;
    this.createdAt = createdAt;
    this.updatedAt = updatedAt;
    this.deletedAt = deletedAt;
  }
} 