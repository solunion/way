import { InputType, OmitType } from '@nestjs/graphql';
import { ResourceDto } from './resource.dto';

@InputType('CreateResourceInput')
export class CreateResourceDto extends OmitType(ResourceDto, ['id'] as const) {} 