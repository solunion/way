import { Injectable } from '@nestjs/common';
import { plainToInstance } from 'class-transformer';
import { catchError, Observable, throwError } from 'rxjs';
import { map } from 'rxjs/operators';
import { RuleEntity } from './rule.entity';
import { NewRule, Rule } from './rule.model';
import { RuleRepository } from './rule.repository';
import { RuleType } from './rule-type.model';

@Injectable()
export class RuleService {
  constructor(private readonly ruleRepository: RuleRepository) {}

  create$(newRule: NewRule): Observable<Rule> {
    return this.ruleRepository.create$(this.#transformToEntity(newRule)).pipe(map((entity) => this.#transformToDto(entity)));
  }

  findById$(id: string): Observable<Rule | null> {
    return this.ruleRepository.findById$(id).pipe(map((entity) => (entity ? this.#transformToDto(entity) : null)));
  }

  update$(id: string, updateRule: Rule): Observable<Rule> {
    return this.ruleRepository.update$(id, this.#transformToEntity(updateRule)).pipe(map((entity) => this.#transformToDto(entity)));
  }

  delete$(id: string): Observable<void> {
    return this.ruleRepository.softDelete$(id).pipe(
      map(() => undefined),
      catchError((error) => {
        console.error('Error deleting tenant:', error);
        return throwError(() => new Error('Unable to delete tenant'));
      })
    );
  }

  findAll$(): Observable<Rule[]> {
    return this.ruleRepository.findAll$().pipe(map((entities) => entities.map((entity) => this.#transformToDto(entity))));
  }

  findByTenantId$(tenantId: string): Observable<Rule[]> {
    return this.ruleRepository.findByTenantId$(tenantId).pipe(map((entities) => entities.map((entity) => this.#transformToDto(entity))));
  }

  #transformToEntity(rule: Partial<Rule>): Pick<RuleEntity, 'name' | 'value' | 'tenantId'> {
    return plainToInstance(RuleEntity, rule, { excludeExtraneousValues: true });
  }

  #transformToDto(entity: RuleEntity): Rule {
    return plainToInstance(Rule, entity, { excludeExtraneousValues: true });
  }
}
