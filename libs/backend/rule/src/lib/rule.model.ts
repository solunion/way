import { Field, ObjectType } from '@nestjs/graphql';
import { IsNotEmpty, IsUUID, MaxLength, MinLength } from 'class-validator';

@ObjectType()
export class NewRule {
  @Field()
  @IsNotEmpty()
  @MinLength(3, { message: 'Name must be at least 3 characters long' })
  @MaxLength(50, { message: 'Name must not exceed 50 characters' })
  name: string;

  @Field()
  @IsNotEmpty()
  @MaxLength(50, { message: 'Type must not exceed 50 characters' })
  type: string;

  @Field()
  @IsNotEmpty()
  value: string;

  @Field({ nullable: true })
  tenantId?: string;

  constructor(name: string, type: string, value: string, tenantId?: string) {
    this.name = name;
    this.type = type;
    this.value = value;
    this.tenantId = tenantId;
  }
}

@ObjectType()
export class Rule extends NewRule {
  @Field()
  @IsUUID()
  id: string;

  constructor(id: string, name: string, type: string, value: string, tenantId?: string) {
    super(name, type, value, tenantId);
    this.id = id;
  }
}
