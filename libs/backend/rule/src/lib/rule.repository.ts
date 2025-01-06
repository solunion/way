import { Injectable } from '@nestjs/common';
import { Prisma } from '@prisma/client';
import { DatabaseService } from '@way/backend-database';
import { from, Observable } from 'rxjs';
import { RuleEntity } from './rule.entity';

@Injectable()
export class RuleRepository {
  #db: DatabaseService;

  constructor(private readonly db: DatabaseService) {
    this.#db = db;
  }

  create$(data: Pick<RuleEntity, 'name' | 'value' | 'tenantId'>): Observable<RuleEntity> {
    return from(this.#db.rule.create({ data: this.#transformData(data) }));
  }

  findById$(id: string): Observable<RuleEntity | null> {
    return from(this.#db.rule.findUnique({ where: { id, deletedAt: null } }));
  }

  update$(id: string, data: Pick<RuleEntity, 'name' | 'value' | 'tenantId'>): Observable<RuleEntity> {
    return from(this.#db.rule.update({ where: { id, deletedAt: null }, data: this.#transformData(data) }));
  }

  delete$(id: string): Observable<RuleEntity> {
    return from(this.#db.rule.delete({ where: { id, deletedAt: null } }));
  }

  softDelete$(id: string): Observable<RuleEntity> {
    return from(
      this.db.rule.update({
        where: { id },
        data: { deletedAt: new Date() },
      })
    );
  }

  findAll$(): Observable<RuleEntity[]> {
    return from(this.#db.rule.findMany({ where: { deletedAt: null } }));
  }

  findByTenantId$(tenantId: string): Observable<RuleEntity[]> {
    return from(this.#db.rule.findMany({ where: { tenantId, deletedAt: null } }));
  }

  #transformData(data: Pick<RuleEntity, 'name' | 'value' | 'tenantId'>) {
    return {
      ...data,
      value: !data.value ? Prisma.JsonNull : data.value,
    };
  }
}
