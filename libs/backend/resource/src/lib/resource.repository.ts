import { Injectable } from '@nestjs/common';
import { DatabaseService } from '@way/backend-database';
import { from, Observable } from 'rxjs';
import { ResourceEntity } from './resource.entity';

@Injectable()
export class ResourceRepository {
  #db: DatabaseService;

  constructor(db: DatabaseService) {
    this.#db = db;
  }

  create$(data: Pick<ResourceEntity, 'name' | 'type' | 'componentId' | 'tenantId'>): Observable<ResourceEntity> {
    return from(this.#db.resource.create({ data }));
  }

  findById$(id: string): Observable<ResourceEntity | null> {
    return from(
      this.#db.resource.findFirst({
        where: { id, deletedAt: null },
      })
    );
  }

  findAll$(): Observable<ResourceEntity[]> {
    return from(
      this.#db.resource.findMany({
        where: { deletedAt: null },
      })
    );
  }

  update$(id: string, data: Partial<Pick<ResourceEntity, 'name' | 'type' | 'componentId' | 'tenantId'>>): Observable<ResourceEntity> {
    return from(
      this.#db.resource.update({
        where: { id },
        data,
      })
    );
  }

  softDelete$(id: string): Observable<ResourceEntity> {
    return from(
      this.#db.resource.update({
        where: { id },
        data: { deletedAt: new Date() },
      })
    );
  }
} 