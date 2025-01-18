import { Args, Mutation, Query, Resolver } from '@nestjs/graphql';
import { plainToInstance } from 'class-transformer';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { CreateRuleSetDto, RuleSetDto, UpdateRuleSetDto } from './dto/rule-set.dto';
import { RuleSet } from './rule-set.model';
import { RuleSetService } from './rule-set.service';

@Resolver(() => RuleSetDto)
export class RuleSetResolver {
  constructor(private readonly service: RuleSetService) {}

  @Mutation(() => RuleSetDto)
  createRuleSet(@Args('createRuleSetDto') createRuleSetDto: CreateRuleSetDto): Observable<RuleSetDto> {
    const ruleSet = plainToInstance(RuleSet, createRuleSetDto);
    return this.service.create$(ruleSet).pipe(map((createdRuleSet) => plainToInstance(RuleSetDto, createdRuleSet)));
  }

  @Query(() => RuleSetDto, { nullable: true })
  getRuleSetById(@Args('id') id: string): Observable<RuleSetDto | null> {
    return this.service.findById$(id).pipe(map((ruleSet) => (ruleSet ? plainToInstance(RuleSetDto, ruleSet) : null)));
  }

  @Mutation(() => RuleSetDto)
  updateRuleSet(@Args('id') id: string, @Args('updateRuleSetDto') updateRuleSetDto: UpdateRuleSetDto): Observable<RuleSetDto> {
    const ruleSet = plainToInstance(RuleSet, updateRuleSetDto);
    return this.service.update$(id, ruleSet).pipe(map((updatedRuleSet) => plainToInstance(RuleSetDto, updatedRuleSet)));
  }

  @Mutation(() => RuleSetDto)
  deleteRuleSet(@Args('id') id: string): Observable<RuleSetDto> {
    return this.service.delete$(id).pipe(map((deletedRuleSet) => plainToInstance(RuleSetDto, deletedRuleSet)));
  }

  @Query(() => [RuleSetDto])
  getRuleSets(): Observable<RuleSetDto[]> {
    return this.service.findAll$().pipe(map((ruleSets) => ruleSets.map((ruleSet) => plainToInstance(RuleSetDto, ruleSet))));
  }

  @Query(() => [RuleSetDto])
  getRuleSetsByTenantId(@Args('tenantId') tenantId: string): Observable<RuleSetDto[]> {
    return this.service.findByTenantId$(tenantId).pipe(map((ruleSets) => ruleSets.map((ruleSet) => plainToInstance(RuleSetDto, ruleSet))));
  }
}
