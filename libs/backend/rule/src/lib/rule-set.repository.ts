import { Injectable } from '@nestjs/common';
import { DatabaseService } from '@way/backend-database';
import { from, Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { RuleSetEntity } from './rule-set.entity';

@Injectable()
export class RuleSetRepository {
  constructor(private readonly db: DatabaseService) {}

  create$(data: RuleSetEntity): Observable<RuleSetEntity> {
    return from(this.db.ruleSet.create({ data: data }));
  }

  findById$(id: string): Observable<RuleSetEntity | null> {
    return from(this.db.ruleSet.findUnique({ where: { id } }));
  }

  update$(id: string, data: Partial<RuleSetEntity>): Observable<RuleSetEntity> {
    return from(this.db.ruleSet.update({ where: { id }, data }));
  }

  delete$(id: string): Observable<void> {
    return from(this.db.ruleSet.delete({ where: { id } })).pipe(map(() => undefined));
  }

  softDelete$(id: string): Observable<void> {
    return from(this.db.ruleSet.delete({ where: { id, deletedAt: null } })).pipe(map(() => undefined));
  }

  findAll$(): Observable<RuleSetEntity[]> {
    return from(this.db.ruleSet.findMany());
  }

  findByTenantId$(tenantId: string): Observable<RuleSetEntity[]> {
    return from(this.db.ruleSet.findMany({ where: { tenantId } }));
  }

  findByParentId$(parentId: string): Observable<RuleSetEntity[]> {
    return from(this.db.ruleSet.findMany({ where: { parentId } }));
  }
}
