import { Module } from '@nestjs/common';
import { DatabaseModule } from '@way/backend-database';
import { RuleSetController } from './rule-set.controller';
import { RuleSetRepository } from './rule-set.repository';
import { RuleSetResolver } from './rule-set.resolver';
import { RuleSetService } from './rule-set.service';
import { RuleController } from './rule.controller';
import { RuleRepository } from './rule.repository';
import { RuleResolver } from './rule.resolver';
import { RuleService } from './rule.service';

@Module({
  imports: [DatabaseModule],
  controllers: [RuleController, RuleSetController],
  providers: [RuleService, RuleSetService, RuleRepository, RuleSetRepository, RuleResolver, RuleSetResolver],
  exports: [RuleService, RuleSetService],
})
export class RuleModule {}
