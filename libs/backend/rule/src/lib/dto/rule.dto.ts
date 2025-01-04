import { Field, InputType, ObjectType } from '@nestjs/graphql';
import { IsNotEmpty, IsOptional, IsString, IsUUID, MaxLength, MinLength } from 'class-validator';
import { Prisma } from '@prisma/client';

@InputType('RuleInput')
@ObjectType('Rule')
export class RuleDto {
  @Field()
  @IsUUID()
  id: string;

  @Field()
  @IsNotEmpty()
  @IsString()
  @MinLength(3)
  @MaxLength(50)
  name: string;

  @Field()
  @IsNotEmpty()
  @IsString()
  @MaxLength(50)
  type: string;

  @Field()
  @IsNotEmpty()
  value: string;

  @Field({ nullable: true })
  @IsOptional()
  @IsUUID()
  tenantId?: string;

  constructor(id: string, name: string, type: string, value: string, tenantId?: string) {
    this.id = id;
    this.name = name;
    this.type = type;
    this.value = value;
    this.tenantId = tenantId;
  }
}

@InputType()
export class CreateRuleDto {
  @Field()
  @IsNotEmpty()
  @IsString()
  @MinLength(3)
  @MaxLength(50)
  name: string;

  @Field()
  @IsNotEmpty()
  @IsString()
  @MaxLength(50)
  type: string;

  @Field()
  @IsNotEmpty()
  value: string;

  @Field({ nullable: true })
  @IsOptional()
  @IsUUID()
  tenantId?: string;

  constructor(name: string, type: string, value: string, tenantId?: string) {
    this.name = name;
    this.type = type;
    this.value = value;
    this.tenantId = tenantId;
  }
}

@InputType()
export class UpdateRuleDto {
  @Field({ nullable: true })
  @IsOptional()
  @IsString()
  @MinLength(3)
  @MaxLength(50)
  name?: string;

  @Field({ nullable: true })
  @IsOptional()
  @IsString()
  @MaxLength(50)
  type?: string;

  @Field({ nullable: true })
  @IsOptional()
  value?: string;

  @Field({ nullable: true })
  @IsOptional()
  @IsUUID()
  tenantId?: string;
}
