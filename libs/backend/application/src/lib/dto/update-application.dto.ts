import { InputType, PartialType } from '@nestjs/graphql';
import { ApplicationDto } from './application.dto';

@InputType('UpdateApplicationInput')
export class UpdateApplicationDto extends PartialType(ApplicationDto) {}
