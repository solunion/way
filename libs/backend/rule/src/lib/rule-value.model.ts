import { Expose } from 'class-transformer';
import { IsEnum } from 'class-validator';
import { RuleType } from './rule-type.model';

export abstract class RuleValue {
  @Expose()
  @IsEnum(RuleType)
  type: RuleType;

  constructor(type: RuleType) {
    this.type = type;
  }
}
