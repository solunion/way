import { Field, InputType, ObjectType } from '@nestjs/graphql';
import { Expose, Transform, Type } from 'class-transformer';
import { IsEnum, IsNotEmpty, IsNotEmptyObject, IsOptional, IsString, IsUUID, MaxLength, MinLength, ValidateNested } from 'class-validator';
import { GraphQLJSON } from 'graphql-type-json';
import { HttpRuleValue } from '../model/rule/http/http-rule-value.model';
import { RuleType } from '../rule-type.model';
import { RuleValue } from '../rule-value.model';

@InputType()
@ObjectType('Rule')
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

  @Field()
  @Expose()
  @IsEnum(RuleType)
  type: string;

  @Field(() => GraphQLJSON)
  @Expose()
  // @IsNotEmptyObject()
  @Transform(({value, obj}) => {
    // value = obj.value;
    // console.log("Response: ", value);
    value.type = obj.type;
    return value;
  }, {toPlainOnly: true})
  @Type(() => RuleValue, {
    discriminator: {
      property: 'type',
      subTypes: [{ value: HttpRuleValue, name: RuleType.HTTP }],
    },
    keepDiscriminatorProperty: true,
  })
  value: HttpRuleValue | RuleValue;

  @IsUUID()
  @Expose()
  @IsOptional()
  @Field({ nullable: true })
  tenantId?: string;

  constructor(id: string, name: string, type: string, value: HttpRuleValue | RuleValue, tenantId?: string) {
    this.id = id;
    this.name = name;
    this.type = type;
    this.value = value;
    this.tenantId = tenantId;
  }
}
