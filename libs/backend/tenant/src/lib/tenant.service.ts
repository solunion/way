import { Injectable } from '@nestjs/common';
import { DatabaseService } from '@way/backend-database';
import { filter, from, map, Observable } from 'rxjs';
import { TenantOutput } from './tenant.output.model';
import { TenantEntity } from '@prisma/client';

@Injectable()
export class TenantService {
  #db: DatabaseService;

  constructor(db: DatabaseService) {
    this.#db = db;
  }

  create$(newTenant: TenantOutput): Observable<TenantOutput> {
    return from(
      this.#db.tenantEntity.create({
        select: {
          id: true,
          name: true,
          description: true,
        },
        data: newTenant,
      })
    ).pipe(
      filter((entity) => !!entity),
      map(
        (entity: TenantEntity) =>
          new TenantOutput({
            id: entity.id,
            name: entity.name,
            description: entity.description ? entity.description : undefined,
          })
      )
    );
  }

  getOne$(id: string): Observable<TenantOutput> {
    return from(
      this.#db.tenantEntity.findUnique({
        where: {
          id: id,
        },
      })
    ).pipe(
      filter((entity) => !!entity),
      map(
        (entity: TenantEntity) =>
          new TenantOutput({
            id: entity.id,
            name: entity.name,
            description: entity.description ? entity.description : undefined,
          })
      )
    );
  }
}
