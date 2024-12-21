import { Module } from '@nestjs/common';
import { TenantModule } from '@way/backend-tenant';
import { AppController } from './app.controller';
import { AppService } from './app.service';

@Module({
  imports: [TenantModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
