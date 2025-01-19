import { Field } from '@nestjs/graphql';
import { IsNotEmpty, IsUUID, MaxLength, MinLength, IsEnum } from 'class-validator';

export enum ResourceType {
  HTTP = 'HTTP',
  ROUTE = 'ROUTE',
  FORM = 'FORM',
  MODEL = 'MODEL'
}

export class NewResource {
  @IsNotEmpty()
  @MinLength(3, { message: 'Name must be at least 3 characters long' })
  @MaxLength(50, { message: 'Name must not exceed 50 characters' })
  name: string;

  @IsNotEmpty()
  @IsEnum(ResourceType)
  type: ResourceType;

  @IsNotEmpty()
  @IsUUID()
  componentId: string;

  @Field({ nullable: true })
  @IsUUID()
  tenantId?: string;

  constructor(name: string, type: ResourceType, componentId: string, tenantId?: string) {
    this.name = name;
    this.type = type;
    this.componentId = componentId;
    this.tenantId = tenantId;
  }
}

export class Resource extends NewResource {
  @IsUUID()
  id: string;

  constructor(id: string, name: string, type: ResourceType, componentId: string, tenantId?: string) {
    super(name, type, componentId, tenantId);
    this.id = id;
  }
} 