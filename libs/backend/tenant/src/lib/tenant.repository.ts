import { Injectable } from '@nestjs/common';

import { DatabaseService } from '@way/backend-database';
import { from, Observable } from 'rxjs';
import { TenantEntity } from './tenant.entity';

@Injectable()
export class TenantRepository {
  constructor(private readonly db: DatabaseService) {}

  create$(data: Pick<TenantEntity, 'name' | 'description'>): Observable<TenantEntity> {
    return from(this.db.tenant.create({ data }));
  }

  findById$(id: string): Observable<TenantEntity | null> {
    return from(
      this.db.tenant.findFirst({
        where: { id, deletedAt: null },
      })
    );
  }

  update$(id: string, data: Partial<Pick<TenantEntity, 'name' | 'description'>>): Observable<TenantEntity> {
    return from(
      this.db.tenant.update({
        where: { id },
        data,
      })
    );
  }

  softDelete$(id: string): Observable<TenantEntity> {
    return from(
      this.db.tenant.update({
        where: { id },
        data: { deletedAt: new Date() },
      })
    );
  }

  findAll$(): Observable<TenantEntity[]> {
    return from(
      this.db.tenant.findMany({
        where: { deletedAt: null },
      })
    );
  }
}
