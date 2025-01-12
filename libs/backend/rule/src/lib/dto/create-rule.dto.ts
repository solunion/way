import { InputType, OmitType } from '@nestjs/graphql';
import { RuleDto } from './rule.dto';

@InputType()
export class CreateRuleDto extends OmitType(RuleDto, ['id'] as const) {}
