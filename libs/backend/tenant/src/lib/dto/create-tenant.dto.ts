import { InputType, OmitType } from '@nestjs/graphql';
import { TenantDto } from './tenant.dto';

@InputType('CreateTenantInput')
export class CreateTenantDto extends OmitType(TenantDto, ['id'] as const) {}
