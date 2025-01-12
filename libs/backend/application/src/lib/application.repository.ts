import { Injectable } from '@nestjs/common';
import { DatabaseService } from '@way/backend-database';
import { from, Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { ApplicationEntity } from './application.entity';

@Injectable()
export class ApplicationRepository {
  #db: DatabaseService;
  constructor(db: DatabaseService) {
    this.#db = db;
  }

  create$(data: ApplicationEntity): Observable<ApplicationEntity> {
    return from(this.#db.application.create({ data }));
  }

  findAll$(): Observable<ApplicationEntity[]> {
    return from(this.#db.application.findMany());
  }

  findOne$(id: string): Observable<ApplicationEntity | null> {
    return from(this.#db.application.findUnique({ where: { id } }));
  }

  update$(id: string, data: Partial<ApplicationEntity>): Observable<ApplicationEntity> {
    return from(
      this.#db.application.update({
        where: { id },
        data,
      })
    );
  }

  softDelete$(id: string): Observable<void> {
    return from(
      this.#db.tenant.update({
        where: { id },
        data: { deletedAt: new Date() },
      })
    ).pipe(map(() => undefined));
  }
}
