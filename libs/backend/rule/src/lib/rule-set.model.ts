import { Field, ObjectType } from '@nestjs/graphql';
import { IsNotEmpty, IsUUID, MaxLength, MinLength } from 'class-validator';

@ObjectType()
export class NewRuleSet {
  @Field()
  @IsNotEmpty()
  @MinLength(3, { message: 'Name must be at least 3 characters long' })
  @MaxLength(50, { message: 'Name must not exceed 50 characters' })
  name: string;

  @Field({ nullable: true })
  tenantId?: string;

  constructor(name: string, tenantId?: string) {
    this.name = name;
    this.tenantId = tenantId;
  }
}

@ObjectType()
export class RuleSet extends NewRuleSet {
  @Field()
  @IsUUID()
  id: string;

  constructor(id: string, name: string, tenantId?: string) {
    super(name, tenantId);
    this.id = id;
  }
} 