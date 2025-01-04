import { Injectable } from '@nestjs/common';
import { plainToInstance } from 'class-transformer';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { RuleSetEntity } from './rule-set.entity';
import { NewRuleSet, RuleSet } from './rule-set.model';
import { RuleSetRepository } from './rule-set.repository';

@Injectable()
export class RuleSetService {
  constructor(private readonly ruleSetRepository: RuleSetRepository) {}

  create$(newRuleSet: NewRuleSet): Observable<RuleSet> {
    return this.ruleSetRepository.create$(this.#transformToEntity(newRuleSet)).pipe(map((entity) => this.#transformToDto(entity)));
  }

  findById$(id: string): Observable<RuleSet | null> {
    return this.ruleSetRepository.findById$(id).pipe(map((entity) => (entity ? this.#transformToDto(entity) : null)));
  }

  update$(id: string, updateRuleSet: Partial<RuleSet>): Observable<RuleSet> {
    return this.ruleSetRepository.update$(id, this.#transformToEntity(updateRuleSet)).pipe(map((entity) => this.#transformToDto(entity)));
  }

  delete$(id: string): Observable<RuleSet> {
    return this.ruleSetRepository.delete$(id).pipe(map((entity) => this.#transformToDto(entity)));
  }

  findAll$(): Observable<RuleSet[]> {
    return this.ruleSetRepository.findAll$().pipe(map((entities) => entities.map((entity) => this.#transformToDto(entity))));
  }

  findByTenantId$(tenantId: string): Observable<RuleSet[]> {
    return this.ruleSetRepository.findByTenantId$(tenantId).pipe(map((entities) => entities.map((entity) => this.#transformToDto(entity))));
  }

  findByParentId$(parentId: string): Observable<RuleSet[]> {
    return this.ruleSetRepository.findByParentId$(parentId).pipe(map((entities) => entities.map((entity) => this.#transformToDto(entity))));
  }

  #transformToEntity(dto: Partial<RuleSet>): RuleSetEntity {
    return {
      ...plainToInstance(RuleSetEntity, dto),
      parentId: dto.parent?.id || null,
    };
  }

  #transformToDto(entity: RuleSetEntity): RuleSet {
    return plainToInstance(RuleSet, entity);
  }
}
