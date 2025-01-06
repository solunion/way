import { Prisma, Rule } from '@prisma/client';
import { Exclude, Expose } from 'class-transformer';

export class RuleEntity implements Rule {

  @Expose()
  id: string;
  @Expose()
  name: string;
  @Expose()
  type: string;
  @Expose()
  value: Prisma.JsonValue;
  @Expose()
  tenantId: string | null;
  @Exclude()
  createdAt: Date;
  @Exclude()
  updatedAt: Date;
  @Exclude()
  deletedAt: Date | null;


  constructor(id: string, name: string, type: string, value: Prisma.JsonValue, tenantId: string | null, createdAt: Date, updatedAt: Date, deletedAt: Date | null) {
    this.id = id;
    this.name = name;
    this.type = type;
    this.value = value;
    this.tenantId = tenantId;
    this.createdAt = createdAt;
    this.updatedAt = updatedAt;
    this.deletedAt = deletedAt;
  }
}
