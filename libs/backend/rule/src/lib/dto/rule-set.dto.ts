import { Field, InputType, ObjectType } from '@nestjs/graphql';
import { IsNotEmpty, IsOptional, IsString, IsUUID, MaxLength, MinLength } from 'class-validator';

@InputType('RuleSetInput')
@ObjectType('RuleSet')
export class RuleSetDto {
  @Field()
  @IsUUID()
  id: string;

  @Field()
  @IsNotEmpty()
  @IsString()
  @MinLength(3)
  @MaxLength(50)
  name: string;

  @Field({ nullable: true })
  @IsOptional()
  @IsUUID()
  tenantId?: string;

  constructor(id: string, name: string, tenantId?: string) {
    this.id = id;
    this.name = name;
    this.tenantId = tenantId;
  }
}

@InputType()
export class CreateRuleSetDto {
  @Field()
  @IsNotEmpty()
  @IsString()
  @MinLength(3)
  @MaxLength(50)
  name: string;

  @Field({ nullable: true })
  @IsOptional()
  @IsUUID()
  tenantId?: string;

  constructor(name: string, tenantId?: string) {
    this.name = name;
    this.tenantId = tenantId;
  }
}

@InputType()
export class UpdateRuleSetDto {
  @Field({ nullable: true })
  @IsOptional()
  @IsString()
  @MinLength(3)
  @MaxLength(50)
  name?: string;

  @Field({ nullable: true })
  @IsOptional()
  @IsUUID()
  tenantId?: string;
} 