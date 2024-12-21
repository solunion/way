import { Injectable } from '@nestjs/common';
import { DatabaseService } from '@way/backend-database';
import { from, Observable } from 'rxjs';
import { Tenant } from './tenant.model';
import { TenantEntity } from '@prisma/client';

@Injectable()
export class TenantService {
  constructor(private db: DatabaseService) {}

  create(newTenant: Tenant): Observable<Tenant> {
    return from(
      this.db.tenantEntity
        .create({
          select: {
            id: true,
            name: true,
            description: true,
          },
          data: newTenant,
        })
        .then(
          (entity: TenantEntity | null) =>
            <Tenant>{
              id: entity?.id,
              name: entity?.name,
              description: entity?.description,
            }
        )
    );
  }

  getOne(id: string): Observable<Tenant> {
    return from(
      this.db.tenantEntity.findUnique({
        where: {
          id: id,
        },
      }).then(
        (entity: TenantEntity | null) =>
          <Tenant>{
            id: entity?.id,
            name: entity?.name,
            description: entity?.description,
          }
      )
    );
  }
}
