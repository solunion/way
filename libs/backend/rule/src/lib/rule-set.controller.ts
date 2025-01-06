import { Body, Controller, Delete, Get, Param, Post, Put } from '@nestjs/common';
import { plainToInstance } from 'class-transformer';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { CreateRuleSetDto, RuleSetDto, UpdateRuleSetDto } from './dto/rule-set.dto';
import { RuleSet } from './rule-set.model';
import { RuleSetService } from './rule-set.service';

@Controller('rule-sets')
export class RuleSetController {
  constructor(private readonly service: RuleSetService) {}

  @Post()
  create(@Body() createRuleSetDto: CreateRuleSetDto): Observable<RuleSetDto> {
    const ruleSet = plainToInstance(RuleSet, createRuleSetDto);
    return this.service.create$(ruleSet).pipe(map((createdRuleSet) => plainToInstance(RuleSetDto, createdRuleSet)));
  }

  @Get(':id')
  findOne(@Param('id') id: string): Observable<RuleSetDto | null> {
    return this.service.findById$(id).pipe(map((ruleSet) => (ruleSet ? plainToInstance(RuleSetDto, ruleSet) : null)));
  }

  @Put(':id')
  update(@Param('id') id: string, @Body() updateRuleSetDto: UpdateRuleSetDto): Observable<RuleSetDto> {
    const ruleSet = plainToInstance(RuleSet, updateRuleSetDto);
    return this.service.update$(id, ruleSet).pipe(map((updatedRuleSet) => plainToInstance(RuleSetDto, updatedRuleSet)));
  }

  @Delete(':id')
  remove(@Param('id') id: string): Observable<void> {
    return this.service.delete$(id);
  }

  @Get()
  findAll(): Observable<RuleSetDto[]> {
    return this.service.findAll$().pipe(map((ruleSets) => ruleSets.map((ruleSet) => plainToInstance(RuleSetDto, ruleSet))));
  }

  @Get('tenant/:tenantId')
  findByTenantId(@Param('tenantId') tenantId: string): Observable<RuleSetDto[]> {
    return this.service.findByTenantId$(tenantId).pipe(map((ruleSets) => ruleSets.map((ruleSet) => plainToInstance(RuleSetDto, ruleSet))));
  }
}
