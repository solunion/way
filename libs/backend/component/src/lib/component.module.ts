import { Module } from '@nestjs/common';
import { DatabaseModule } from '@way/backend-database';
import { ComponentController } from './component.controller';
import { ComponentResolver } from './component.resolver';
import { ComponentService } from './component.service';
import { ComponentRepository } from './component.repository';

@Module({
  imports: [DatabaseModule],
  controllers: [ComponentController],
  providers: [ComponentResolver, ComponentService, ComponentRepository],
  exports: [ComponentService],
})
export class ComponentModule {}
