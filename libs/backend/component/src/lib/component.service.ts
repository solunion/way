import { Injectable } from '@nestjs/common';
import { plainToInstance } from 'class-transformer';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { ComponentEntity } from './component.entity';
import { Component, NewComponent } from './component.model';
import { ComponentRepository } from './component.repository';

@Injectable()
export class ComponentService {
  #repository: ComponentRepository;

  constructor(repository: ComponentRepository) {
    this.#repository = repository;
  }

  create$(newComponent: NewComponent): Observable<Component> {
    return this.#repository
      .create$(this.#transformToEntity(newComponent))
      .pipe(map((entity) => this.#transformToModel(entity)));
  }

  findById$(id: string): Observable<Component | null> {
    return this.#repository
      .findById$(id)
      .pipe(
        map((entity) => (entity ? this.#transformToModel(entity) : null))
      );
  }

  findAll$(): Observable<Component[]> {
    return this.#repository
      .findAll$()
      .pipe(map((entities) => entities.map((e) => this.#transformToModel(e))));
  }

  update$(id: string, component: Partial<Component>): Observable<Component> {
    return this.#repository
      .update$(id, this.#transformToEntity(component))
      .pipe(map((entity) => this.#transformToModel(entity)));
  }

  delete$(id: string): Observable<void> {
    return this.#repository.softDelete$(id);
  }

  #transformToEntity(model: Partial<Component>): ComponentEntity {
    return plainToInstance(ComponentEntity, model, { excludeExtraneousValues: true });
  }

  #transformToModel(entity: ComponentEntity): Component {
    return plainToInstance(Component, entity, { excludeExtraneousValues: true });
  }
} 