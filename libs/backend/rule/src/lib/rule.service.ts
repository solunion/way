import { Injectable } from '@nestjs/common';
import { Observable } from 'rxjs';
import { RuleRepository } from './rule.repository';
import { NewRule, Rule } from './rule.model';
import { plainToInstance } from 'class-transformer';
import { RuleEntity } from './rule.entity';
import { map } from 'rxjs/operators';

@Injectable()
export class RuleService {
  constructor(private readonly ruleRepository: RuleRepository) {}

  create$(newRule: NewRule): Observable<Rule> {
    return this.ruleRepository.create$(this.#transformToEntity(newRule)).pipe(
      map(entity => this.#transformToDto(entity)),
    );
  }

  findById$(id: string): Observable<Rule | null> {
    return this.ruleRepository.findById$(id).pipe(
      map(entity => entity ? this.#transformToDto(entity) : null),
    );
  }

  update$(id: string, updateRule: Partial<Rule>): Observable<Rule> {
    return this.ruleRepository.update$(id, this.#transformToEntity(updateRule)).pipe(
      map(entity => this.#transformToDto(entity)),
    );
  }

  delete$(id: string): Observable<Rule> {
    return this.ruleRepository.delete$(id).pipe(
      map(entity => this.#transformToDto(entity)),
    );
  }

  findAll$(): Observable<Rule[]> {
    return this.ruleRepository.findAll$().pipe(
      map(entities => entities.map(entity => this.#transformToDto(entity))),
    );
  }

  findByTenantId$(tenantId: string): Observable<Rule[]> {
    return this.ruleRepository.findByTenantId$(tenantId).pipe(
      map(entities => entities.map(entity => this.#transformToDto(entity))),
    );
  }

  #transformToEntity(dto: Partial<Rule>): Pick<RuleEntity, 'name' | 'type' | 'value' | 'tenantId'> {
    return plainToInstance(RuleEntity, dto);
  }

  #transformToDto(entity: RuleEntity): Rule {
    return plainToInstance(Rule, entity);
  }
} 