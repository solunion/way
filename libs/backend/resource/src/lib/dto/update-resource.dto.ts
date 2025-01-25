import { InputType, OmitType, PartialType } from '@nestjs/graphql';
import { ResourceDto } from './resource.dto';

@InputType('UpdateResourceInput')
export class UpdateResourceDto extends PartialType(OmitType(ResourceDto, ['id'] as const)) {} 