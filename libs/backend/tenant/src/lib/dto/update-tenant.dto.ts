import { InputType, OmitType, PartialType } from '@nestjs/graphql';
import { TenantDto } from './tenant.dto';

@InputType()
export class UpdateTenantDto extends PartialType(OmitType(TenantDto, ['id'] as const)) {}
