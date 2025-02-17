import { Args, Mutation, Query, Resolver } from '@nestjs/graphql';
import { plainToInstance } from 'class-transformer';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { CreateRuleDto } from './dto/create-rule.dto';
import { RuleDto } from './dto/rule.dto';
import { UpdateRuleDto } from './dto/update-rule.dto';
import { Rule } from './rule.model';
import { RuleService } from './rule.service';

@Resolver(() => RuleDto)
export class RuleResolver {
  constructor(private readonly ruleService: RuleService) {}

  @Mutation(() => RuleDto)
  createRule(@Args('createRuleDto') createRuleDto: CreateRuleDto): Observable<RuleDto> {
    const rule = plainToInstance(Rule, createRuleDto);
    return this.ruleService.create$(rule).pipe(map((createdRule) => plainToInstance(RuleDto, createdRule)));
  }

  @Query(() => RuleDto, { nullable: true })
  getRuleById(@Args('id') id: string): Observable<RuleDto | null> {
    return this.ruleService.findById$(id).pipe(map((rule) => (rule ? plainToInstance(RuleDto, rule) : null)));
  }

  @Mutation(() => RuleDto)
  updateRule(@Args('id') id: string, @Args('updateRuleDto') updateRuleDto: UpdateRuleDto): Observable<RuleDto> {
    const rule = plainToInstance(Rule, updateRuleDto);
    return this.ruleService.update$(id, rule).pipe(map((updatedRule) => plainToInstance(RuleDto, updatedRule)));
  }

  @Mutation(() => RuleDto)
  deleteRule(@Args('id') id: string): Observable<RuleDto> {
    return this.ruleService.delete$(id).pipe(map((deletedRule) => plainToInstance(RuleDto, deletedRule)));
  }

  @Query(() => [RuleDto])
  getRules(): Observable<RuleDto[]> {
    return this.ruleService.findAll$().pipe(map((rules) => rules.map((rule) => plainToInstance(RuleDto, rule))));
  }

  @Query(() => [RuleDto])
  getRulesByTenantId(@Args('tenantId') tenantId: string): Observable<RuleDto[]> {
    return this.ruleService.findByTenantId$(tenantId).pipe(map((rules) => rules.map((rule) => plainToInstance(RuleDto, rule))));
  }
}
