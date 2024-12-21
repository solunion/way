import { Injectable } from '@nestjs/common';
import { DatabaseService } from '@way/backend-database';
import { from, Observable } from 'rxjs';
import { TenantDto } from './dto/tenant-dto.model';
import { Tenant } from '@prisma/client';

@Injectable()
export class TenantService {
  constructor(private db: DatabaseService) {}

  create(newTenant: TenantDto): Observable<TenantDto> {
    return from(
      this.db.tenant
        .create({
          select: {
            id: true,
            name: true,
            description: true,
          },
          data: newTenant,
        })
        .then(
          (entity: Tenant) =>
            <TenantDto>{
              id: entity.id,
              name: entity.name,
              description: entity.description,
            }
        )
    );
  }

  getOne(id: string): Observable<TenantDto> {
    return from(
      this.db.tenant.findUnique({
        where: {
          id: id,
        },
      }).then(
        (entity: Tenant | null) =>
          <TenantDto>{
            id: entity?.id,
            name: entity?.name,
            description: entity?.description,
          }
      )
    );
  }
}
