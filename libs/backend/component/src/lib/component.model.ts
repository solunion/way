import { Field } from '@nestjs/graphql';
import { IsNotEmpty, IsUUID, MaxLength, MinLength } from 'class-validator';

export class NewComponent {
  @IsNotEmpty()
  @MinLength(3, { message: 'Name must be at least 3 characters long' })
  @MaxLength(50, { message: 'Name must not exceed 50 characters' })
  name: string;

  @Field({ nullable: true })
  @MaxLength(200, { message: 'Description must not exceed 200 characters' })
  description?: string;

  @IsNotEmpty()
  @IsUUID()
  applicationId: string;

  @Field({ nullable: true })
  @IsUUID()
  tenantId?: string;

  constructor(
    name: string,
    applicationId: string,
    description?: string,
    tenantId?: string
  ) {
    this.name = name;
    this.applicationId = applicationId;
    this.description = description;
    this.tenantId = tenantId;
  }
}

export class Component extends NewComponent {
  @IsUUID()
  id: string;

  constructor(
    id: string,
    name: string,
    applicationId: string,
    description?: string,
    tenantId?: string
  ) {
    super(name, applicationId, description, tenantId);
    this.id = id;
  }
} 