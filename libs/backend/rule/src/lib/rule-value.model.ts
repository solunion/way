import { Expose, Transform } from 'class-transformer';
import { IsDefined, IsEnum } from 'class-validator';
import { RuleType } from './rule-type.model';

export abstract class RuleValue {
  @Expose()
  @IsEnum(RuleType)
  @Transform(({ value }) => {
    const result = Object.keys(RuleType)[Object.values(RuleType).indexOf(value?.toUpperCase())];
    if (!result) { throw new Error("Invalid rule type."); }
    return result;
  })
  readonly type: RuleType;

  constructor(type: RuleType) {
    this.type = type;
  }
}
