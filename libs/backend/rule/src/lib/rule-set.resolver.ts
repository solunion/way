import { Args, Mutation, Query, Resolver } from '@nestjs/graphql';
import { plainToInstance } from 'class-transformer';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { CreateRuleSetDto, RuleSetDto, UpdateRuleSetDto } from './dto/rule-set.dto';
import { RuleSet } from './rule-set.model';
import { RuleSetService } from './rule-set.service';

@Resolver(() => RuleSetDto)
export class RuleSetResolver {
  constructor(private readonly ruleSetService: RuleSetService) {}

  @Mutation(() => RuleSetDto)
  createRuleSet(@Args('createRuleSetDto') createRuleSetDto: CreateRuleSetDto): Observable<RuleSetDto> {
    const ruleSet = plainToInstance(RuleSet, createRuleSetDto);
    return this.ruleSetService.create$(ruleSet).pipe(
      map(createdRuleSet => plainToInstance(RuleSetDto, createdRuleSet)),
    );
  }

  @Query(() => RuleSetDto, { nullable: true })
  ruleSet(@Args('id') id: string): Observable<RuleSetDto | null> {
    return this.ruleSetService.findById$(id).pipe(
      map(ruleSet => ruleSet ? plainToInstance(RuleSetDto, ruleSet) : null),
    );
  }

  @Mutation(() => RuleSetDto)
  updateRuleSet(@Args('id') id: string, @Args('updateRuleSetDto') updateRuleSetDto: UpdateRuleSetDto): Observable<RuleSetDto> {
    const ruleSet = plainToInstance(RuleSet, updateRuleSetDto);
    return this.ruleSetService.update$(id, ruleSet).pipe(
      map(updatedRuleSet => plainToInstance(RuleSetDto, updatedRuleSet)),
    );
  }

  @Mutation(() => RuleSetDto)
  removeRuleSet(@Args('id') id: string): Observable<RuleSetDto> {
    return this.ruleSetService.delete$(id).pipe(
      map(deletedRuleSet => plainToInstance(RuleSetDto, deletedRuleSet)),
    );
  }

  @Query(() => [RuleSetDto])
  ruleSets(): Observable<RuleSetDto[]> {
    return this.ruleSetService.findAll$().pipe(
      map(ruleSets => ruleSets.map(ruleSet => plainToInstance(RuleSetDto, ruleSet))),
    );
  }

  @Query(() => [RuleSetDto])
  ruleSetsByTenantId(@Args('tenantId') tenantId: string): Observable<RuleSetDto[]> {
    return this.ruleSetService.findByTenantId$(tenantId).pipe(
      map(ruleSets => ruleSets.map(ruleSet => plainToInstance(RuleSetDto, ruleSet))),
    );
  }
} 