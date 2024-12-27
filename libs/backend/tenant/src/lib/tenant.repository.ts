import { Injectable } from '@nestjs/common';
import { DatabaseService } from '@way/backend-database';
import { TenantEntity } from '@prisma/client';
import { Observable, from } from 'rxjs';

@Injectable()
export class TenantRepository {
  constructor(private readonly db: DatabaseService) {}

  create(data: Pick<TenantEntity, 'name' | 'description'>): Observable<TenantEntity> {
    return from(this.db.tenantEntity.create({ data }));
  }

  findById(id: string): Observable<TenantEntity | null> {
    return from(
      this.db.tenantEntity.findFirst({
        where: { id, deletedAt: null },
      })
    );
  }

  update(id: string, data: Partial<Pick<TenantEntity, 'name' | 'description'>>): Observable<TenantEntity> {
    return from(
      this.db.tenantEntity.update({
        where: { id },
        data,
      })
    );
  }

  softDelete(id: string): Observable<TenantEntity> {
    return from(
      this.db.tenantEntity.update({
        where: { id },
        data: { deletedAt: new Date() },
      })
    );
  }

  findAll(): Observable<TenantEntity[]> {
    return from(
      this.db.tenantEntity.findMany({
        where: { deletedAt: null },
      })
    );
  }
}
