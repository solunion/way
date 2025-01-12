import { Module } from '@nestjs/common';
import { DatabaseModule } from '@way/backend-database';
import { ApplicationService } from './application.service';
import { ApplicationRepository } from './application.repository';
import { ApplicationController } from './application.controller';
import { ApplicationResolver } from './application.resolver';

@Module({
  imports: [DatabaseModule],
  controllers: [ApplicationController],
  providers: [ApplicationService, ApplicationRepository, ApplicationResolver],
  exports: [ApplicationService, ApplicationRepository],
})
export class ApplicationModule {}
