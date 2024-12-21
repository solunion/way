import { Module } from '@nestjs/common';
import { DatabaseModule } from '@way/backend-database';
import { TenantRestController } from './tenant.rest.controller';
import { TenantService } from './tenant.service';

@Module({
  imports: [DatabaseModule],
  controllers: [TenantRestController],
  providers: [TenantService],
  exports: [TenantService],
})
export class TenantModule {}
