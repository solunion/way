import { Injectable } from '@nestjs/common';
import { Prisma } from '@prisma/client';
import { DatabaseService } from '@way/backend-database';
import { from, Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { RuleEntity } from './rule.entity';

@Injectable()
export class RuleRepository {
  #db: DatabaseService;

  constructor(private readonly db: DatabaseService) {
    this.#db = db;
  }

  create$(data: Pick<RuleEntity, 'name' | 'value' | 'tenantId'>): Observable<RuleEntity> {
    return from(this.#db.rule.create({ data: this.#transformDataForCreate(data) }));
  }

  findById$(id: string): Observable<RuleEntity | null> {
    return from(this.#db.rule.findUnique({ where: { id, deletedAt: null } }));
  }

  update$(id: string, data: RuleEntity): Observable<RuleEntity> {
    return from(this.#db.rule.update({ where: { id, deletedAt: null }, data: this.#transformDataForUpdate(data) }));
  }

  delete$(id: string): Observable<void> {
    return from(this.#db.rule.delete({ where: { id, deletedAt: null } })).pipe(map(() => undefined));
  }

  softDelete$(id: string): Observable<void> {
    return from(
      this.db.rule.update({
        where: { id },
        data: { deletedAt: new Date() },
      })
    ).pipe(map(() => undefined));
  }

  findAll$(): Observable<RuleEntity[]> {
    return from(this.#db.rule.findMany({ where: { deletedAt: null } }));
  }

  findByTenantId$(tenantId: string): Observable<RuleEntity[]> {
    return from(this.#db.rule.findMany({ where: { tenantId, deletedAt: null } }));
  }

  #transformDataForCreate(data: Pick<RuleEntity, 'name' | 'value' | 'tenantId'>) {
    return {
      ...data,
      value: !data.value ? Prisma.JsonNull : data.value,
    };
  }

  #transformDataForUpdate(data: Pick<RuleEntity, 'name' | 'value' | 'tenantId'>) {
    return {
      ...data,
      value: !data.value ? undefined : data.value,
    };
  }
}
