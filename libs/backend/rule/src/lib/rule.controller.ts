import { Body, Controller, Delete, Get, Param, Post, Put } from '@nestjs/common';
import { plainToInstance } from 'class-transformer';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { CreateRuleDto, RuleDto, UpdateRuleDto } from './dto/rule.dto';
import { Rule } from './rule.model';
import { RuleService } from './rule.service';

@Controller('rules')
export class RuleController {
  constructor(private readonly ruleService: RuleService) {}

  @Post()
  create(@Body() createRuleDto: CreateRuleDto): Observable<RuleDto> {
    const rule = plainToInstance(Rule, createRuleDto);
    return this.ruleService.create$(rule).pipe(
      map(createdRule => plainToInstance(RuleDto, createdRule)),
    );
  }

  @Get(':id')
  findOne(@Param('id') id: string): Observable<RuleDto | null> {
    return this.ruleService.findById$(id).pipe(
      map(rule => rule ? plainToInstance(RuleDto, rule) : null),
    );
  }

  @Put(':id')
  update(@Param('id') id: string, @Body() updateRuleDto: UpdateRuleDto): Observable<RuleDto> {
    const rule = plainToInstance(Rule, updateRuleDto);
    return this.ruleService.update$(id, rule).pipe(
      map(updatedRule => plainToInstance(RuleDto, updatedRule)),
    );
  }

  @Delete(':id')
  remove(@Param('id') id: string): Observable<RuleDto> {
    return this.ruleService.delete$(id).pipe(
      map(deletedRule => plainToInstance(RuleDto, deletedRule)),
    );
  }

  @Get()
  findAll(): Observable<RuleDto[]> {
    return this.ruleService.findAll$().pipe(
      map(rules => rules.map(rule => plainToInstance(RuleDto, rule))),
    );
  }

  @Get('tenant/:tenantId')
  findByTenantId(@Param('tenantId') tenantId: string): Observable<RuleDto[]> {
    return this.ruleService.findByTenantId$(tenantId).pipe(
      map(rules => rules.map(rule => plainToInstance(RuleDto, rule))),
    );
  }
} 