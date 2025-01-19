import { Injectable } from '@nestjs/common';
import { plainToInstance } from 'class-transformer';
import { Observable, map } from 'rxjs';
import { ResourceEntity } from './resource.entity';
import { NewResource, Resource } from './resource.model';
import { ResourceRepository } from './resource.repository';

@Injectable()
export class ResourceService {
  #repository: ResourceRepository;

  constructor(repository: ResourceRepository) {
    this.#repository = repository;
  }

  create$(newResource: NewResource): Observable<Resource> {
    return this.#repository
      .create$(this.#transformToEntity(newResource))
      .pipe(map((entity) => this.#transformToModel(entity)));
  }

  findById$(id: string): Observable<Resource | null> {
    return this.#repository
      .findById$(id)
      .pipe(map((entity) => (entity ? this.#transformToModel(entity) : null)));
  }

  findAll$(): Observable<Resource[]> {
    return this.#repository
      .findAll$()
      .pipe(map((entities) => entities.map((e) => this.#transformToModel(e))));
  }

  update$(id: string, resource: Partial<Resource>): Observable<Resource> {
    return this.#repository
      .update$(id, this.#transformToEntity(resource))
      .pipe(map((entity) => this.#transformToModel(entity)));
  }

  delete$(id: string): Observable<void> {
    return this.#repository.softDelete$(id).pipe(map(() => undefined));
  }

  #transformToEntity(model: Partial<Resource>): ResourceEntity {
    return plainToInstance(ResourceEntity, model, { excludeExtraneousValues: true });
  }

  #transformToModel(entity: ResourceEntity): Resource {
    return plainToInstance(Resource, entity, { excludeExtraneousValues: true });
  }
} 