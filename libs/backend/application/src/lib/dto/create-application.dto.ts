import { InputType, OmitType } from '@nestjs/graphql';
import { ApplicationDto } from './application.dto';

@InputType()
export class CreateApplicationDto extends OmitType(ApplicationDto, ['id'] as const) {}
