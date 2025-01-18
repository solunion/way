import { Field, InputType, ObjectType } from '@nestjs/graphql';
import { Expose } from 'class-transformer';
import { IsNotEmpty, IsOptional, IsString, IsUUID, MaxLength, MinLength } from 'class-validator';
import { GraphQLJSON } from 'graphql-type-json';

@InputType()
@ObjectType('RuleOutput')
export class RuleDto {
  @Field()
  @IsUUID()
  @Expose()
  id: string;

  @Field()
  @Expose()
  @IsNotEmpty()
  @IsString()
  @MinLength(3)
  @MaxLength(50)
  name: string;

  @Field(() => GraphQLJSON)
  @Expose()
  value: JSON;

  @IsUUID()
  @Expose()
  @IsOptional()
  @Field({ nullable: true })
  tenantId?: string;

  constructor(id: string, name: string, value: JSON, tenantId?: string) {
    this.id = id;
    this.name = name;
    this.value = value;
    this.tenantId = tenantId;
  }
}
