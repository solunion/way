import { InputType, PartialType } from '@nestjs/graphql';
import { ComponentDto } from './component.dto';

@InputType('UpdateComponentInput')
export class UpdateComponentDto extends PartialType(ComponentDto) {} 