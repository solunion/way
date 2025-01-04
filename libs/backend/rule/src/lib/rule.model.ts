import { IsNotEmpty, IsUUID, MaxLength, MinLength } from 'class-validator';
import { Prisma } from '@prisma/client';

export class NewRule {
  @IsNotEmpty()
  @MinLength(3, { message: 'Name must be at least 3 characters long' })
  @MaxLength(50, { message: 'Name must not exceed 50 characters' })
  name: string;

  @IsNotEmpty()
  @MaxLength(50, { message: 'Type must not exceed 50 characters' })
  type: string;

  @IsNotEmpty()
  value: Prisma.JsonValue;

  tenantId?: string;

  constructor(name: string, type: string, value: Prisma.JsonValue, tenantId?: string) {
    this.name = name;
    this.type = type;
    this.value = value;
    this.tenantId = tenantId;
  }
}

export class Rule extends NewRule {
  @IsUUID()
  id: string;

  constructor(id: string, name: string, type: string, value: Prisma.JsonValue, tenantId?: string) {
    super(name, type, value, tenantId);
    this.id = id;
  }
} 