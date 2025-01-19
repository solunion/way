import { Module } from '@nestjs/common';
import { DatabaseModule } from '@way/backend-database';
import { ResourceController } from './resource.controller';
import { ResourceRepository } from './resource.repository';
import { ResourceResolver } from './resource.resolver';
import { ResourceService } from './resource.service';

@Module({
  imports: [DatabaseModule],
  controllers: [ResourceController],
  providers: [ResourceService, ResourceRepository, ResourceResolver],
  exports: [ResourceService],
})
export class ResourceModule {} 