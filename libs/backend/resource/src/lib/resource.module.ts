import { Module } from '@nestjs/common';
import { DatabaseModule } from '@way/backend-database';
import { ResourceRepository } from './resource.repository';
import { ResourceResolver } from './resource.resolver';
import { ResourceService } from './resource.service';

@Module({
  imports: [DatabaseModule],
  providers: [ResourceService, ResourceRepository, ResourceResolver],
  exports: [ResourceService],
})
export class ResourceModule {} 