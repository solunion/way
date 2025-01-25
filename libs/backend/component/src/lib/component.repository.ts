import { Injectable } from '@nestjs/common';
import { DatabaseService } from '@way/backend-database';
import { from, Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { ComponentEntity } from './component.entity';

@Injectable()
export class ComponentRepository {
  #db: DatabaseService;

  constructor(db: DatabaseService) {
    this.#db = db;
  }

  create$(data: ComponentEntity): Observable<ComponentEntity> {
    return from(this.#db.component.create({ data }));
  }

  findAll$(): Observable<ComponentEntity[]> {
    return from(
      this.#db.component.findMany({
        where: { deletedAt: null }
      })
    );
  }

  findById$(id: string): Observable<ComponentEntity | null> {
    return from(
      this.#db.component.findFirst({
        where: { id, deletedAt: null }
      })
    );
  }

  update$(id: string, data: Partial<ComponentEntity>): Observable<ComponentEntity> {
    return from(
      this.#db.component.update({
        where: { id },
        data
      })
    );
  }

  softDelete$(id: string): Observable<void> {
    return from(
      this.#db.component.update({
        where: { id },
        data: { deletedAt: new Date() }
      })
    ).pipe(map(() => undefined));
  }
} 