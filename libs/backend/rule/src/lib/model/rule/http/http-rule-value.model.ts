import { Expose } from 'class-transformer';
import { IsEnum, IsOptional } from 'class-validator';
import { RuleType } from '../../../rule-type.model';
import { RuleValue } from '../../../rule-value.model';
import { HttpMethod } from './http-method.model';
import { HttpStatus } from './http-status.model';

export class HttpRuleValue extends RuleValue {
  @Expose()
  @IsEnum(HttpStatus)
  status: HttpStatus;

  @Expose()
  @IsEnum(HttpMethod)
  method: HttpMethod;

  @Expose()
  @IsOptional()
  path: string;

  constructor(type: RuleType, status: HttpStatus, method: HttpMethod, path: string) {
    super(type);
    this.status = status;
    this.method = method;
    this.path = path;
  }
}
