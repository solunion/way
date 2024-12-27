import { Injectable } from '@nestjs/common';
import { DatabaseService } from '@way/backend-database';
import { TenantEntity } from '@prisma/client';
import { Observable, from } from 'rxjs';

@Injectable()
export class TenantRepository {
  constructor(private readonly db: DatabaseService) {}

  create(data: Pick<TenantEntity, 'name' | 'description'>): Observable<TenantEntity> {
    return from(
      this.db.tenantEntity.create({
        data,
        select: {
          id: true,
          name: true,
          description: true,
        },
      })
    );
  }

  findById(id: string): Observable<TenantEntity | null> {
    return from(
      this.db.tenantEntity.findUnique({
        where: { id },
      })
    );
  }

  update(id: string, data: Partial<Pick<TenantEntity, 'name' | 'description'>>): Observable<TenantEntity> {
    return from(
      this.db.tenantEntity.update({
        where: { id },
        data,
        select: {
          id: true,
          name: true,
          description: true,
        },
      })
    );
  }

}
