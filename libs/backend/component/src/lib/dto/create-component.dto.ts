import { InputType, OmitType } from '@nestjs/graphql';
import { ComponentDto } from './component.dto';

@InputType()
export class CreateComponentDto extends OmitType(ComponentDto, ['id'] as const) {} 