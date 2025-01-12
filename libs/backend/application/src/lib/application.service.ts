import { Injectable } from '@nestjs/common';
import { plainToInstance } from 'class-transformer';
import { catchError, map, Observable, throwError } from 'rxjs';
import { ApplicationEntity } from './application.entity';
import { Application, NewApplication } from './application.model';
import { ApplicationRepository } from './application.repository';

@Injectable()
export class ApplicationService {
  #repository: ApplicationRepository;

  constructor(readonly repository: ApplicationRepository) {
    this.#repository = repository;
  }

  create$(application: NewApplication): Observable<Application> {
    return this.#repository.create$(this.#transformToEntity(application)).pipe(
      map((entity: ApplicationEntity) => this.#transformToDto(entity)),
      catchError((error) => {
        console.error('Error creating application:', error);
        return throwError(() => new Error('Unable to create application'));
      })
    );
  }

  findAll$(): Observable<Application[]> {
    return this.#repository.findAll$().pipe(map((entities) => entities.map((entity) => this.#transformToDto(entity))));
  }

  findById$(id: string): Observable<Application | null> {
    return this.#repository.findOne$(id).pipe(map((entity) => (entity ? this.#transformToDto(entity) : null)));
  }

  update$(id: string, application: Partial<NewApplication>): Observable<Application> {
    return this.#repository.update$(id, this.#transformToEntity(application)).pipe(map((entity) => this.#transformToDto(entity)));
  }

  delete$(id: string): Observable<void> {
    return this.#repository.softDelete$(id).pipe(
      catchError((error) => {
        console.error('Error deleting application:', error);
        return throwError(() => new Error('Unable to delete application'));
      })
    );
  }

  #transformToEntity(dto: Partial<Application>): ApplicationEntity {
    return plainToInstance(ApplicationEntity, dto, { excludeExtraneousValues: true });
  }

  #transformToDto(entity: ApplicationEntity): Application {
    return plainToInstance(Application, entity, { excludeExtraneousValues: true });
  }
}
