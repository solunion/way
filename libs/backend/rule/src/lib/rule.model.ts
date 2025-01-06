import { Expose, Type } from 'class-transformer';
import { IsNotEmpty, IsNotEmptyObject, IsUUID, MaxLength, MinLength, ValidateNested } from 'class-validator';
import { HttpRuleValue } from './model/rule/http/http-rule-value.model';
import { RuleType } from './rule-type.model';
import { RuleValue } from './rule-value.model';

export class NewRule {
  @Expose()
  @IsNotEmpty()
  @MinLength(3, { message: 'Name must be at least 3 characters long' })
  @MaxLength(50, { message: 'Name must not exceed 50 characters' })
  name: string;

  @Expose()
  @ValidateNested({ each: true })
  @IsNotEmptyObject()
  @Type(() => RuleValue, {
    discriminator: {
      property: 'type',
      subTypes: [{ value: HttpRuleValue, name: RuleType.HTTP }],
    },
    keepDiscriminatorProperty: true,
  })
  value: HttpRuleValue | RuleValue;

  @Expose()
  tenantId?: string;

  constructor(name: string, value: HttpRuleValue | RuleValue, tenantId?: string) {
    this.name = name;
    this.value = value;
    this.tenantId = tenantId;
  }
}

export class Rule extends NewRule {
  @Expose()
  @IsUUID()
  id: string;

  constructor(id: string, name: string, value: RuleValue, tenantId?: string) {
    super(name, value, tenantId);
    this.id = id;
  }
}
