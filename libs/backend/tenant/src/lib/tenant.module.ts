import { Module } from '@nestjs/common';
import { DatabaseModule } from '@way/backend-database';
import { TenantRepository } from './tenant.repository';
import { TenantResolver } from './tenant.resolver';
import { TenantController } from './tenant.controller';
import { TenantService } from './tenant.service';

@Module({
  imports: [DatabaseModule],
  controllers: [TenantController],
  providers: [TenantResolver, TenantService, TenantRepository],
  exports: [TenantService],
})
export class TenantModule {}
