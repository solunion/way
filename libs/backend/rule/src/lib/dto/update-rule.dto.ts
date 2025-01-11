import { InputType, OmitType, PartialType } from '@nestjs/graphql';
import { RuleDto } from './rule.dto';

@InputType()
export class UpdateRuleDto extends PartialType(OmitType(RuleDto, ['id'] as const)) {}
