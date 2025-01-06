import { Expose, Type } from 'class-transformer';
import { IsEnum, IsNotEmpty, IsUUID, MaxLength, MinLength, ValidateNested } from 'class-validator';
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
  @IsEnum(RuleType)
  type: RuleType;

  @Expose()
  @ValidateNested({each: true})
  value: HttpRuleValue | RuleValue;

  @Expose()
  tenantId?: string;

  constructor(name: string, type: RuleType, value: RuleValue, tenantId?: string) {
    this.name = name;
    this.type = type;
    this.value = value;
    this.tenantId = tenantId;
  }
}

export class Rule extends NewRule {
  @Expose()
  @IsUUID()
  id: string;

  constructor(id: string, name: string, type: RuleType, value: RuleValue, tenantId?: string) {
    super(name, type, value, tenantId);
    this.id = id;
  }
}
