import { Injectable } from '@nestjs/common';
import { plainToInstance } from 'class-transformer';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { RuleSetEntity } from './rule-set.entity';
import { NewRuleSet, RuleSet } from './rule-set.model';
import { RuleSetRepository } from './rule-set.repository';

@Injectable()
export class RuleSetService {
  constructor(private readonly repository: RuleSetRepository) {}

  create$(newRuleSet: NewRuleSet): Observable<RuleSet> {
    return this.repository.create$(this.#transformToEntity(newRuleSet)).pipe(map((entity) => this.#transformToDto(entity)));
  }

  findById$(id: string): Observable<RuleSet | null> {
    return this.repository.findById$(id).pipe(map((entity) => (entity ? this.#transformToDto(entity) : null)));
  }

  update$(id: string, updateRuleSet: Partial<RuleSet>): Observable<RuleSet> {
    return this.repository.update$(id, this.#transformToEntity(updateRuleSet)).pipe(map((entity) => this.#transformToDto(entity)));
  }

  delete$(id: string): Observable<void> {
    return this.repository.softDelete$(id);
  }

  findAll$(): Observable<RuleSet[]> {
    return this.repository.findAll$().pipe(map((entities) => entities.map((entity) => this.#transformToDto(entity))));
  }

  findByTenantId$(tenantId: string): Observable<RuleSet[]> {
    return this.repository.findByTenantId$(tenantId).pipe(map((entities) => entities.map((entity) => this.#transformToDto(entity))));
  }

  #transformToEntity(ruleSet: Partial<RuleSet>): RuleSetEntity {
    return plainToInstance(RuleSetEntity, ruleSet);
  }

  #transformToDto(entity: RuleSetEntity): RuleSet {
    return plainToInstance(RuleSet, entity);
  }
}
