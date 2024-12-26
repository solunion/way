import { Injectable } from '@nestjs/common';
import { TenantEntity } from '@prisma/client';
import { catchError, filter, map, Observable, throwError } from 'rxjs';
import { TenantCreateInput } from './tenant-create.input.model';
import { TenantOutput } from './tenant.output.model';
import { TenantRepository } from './tenant.repository';

@Injectable()
export class TenantService {
  #repository: TenantRepository;

  constructor(repository: TenantRepository) {
    this.#repository = repository;
  }

  create$(
    newTenant: Pick<TenantCreateInput, 'name' | 'description'>
  ): Observable<TenantOutput> {
    return this.#repository.create(this.#mapToEntity(newTenant)).pipe(
      map((entity: TenantEntity) => this.#mapToDto(entity)),
      catchError((error) => {
        console.error('Error creating tenant:', error);
        return throwError(() => new Error('Unable to create tenant'));
      })
    );
  }

  findById$(id: string): Observable<TenantOutput> {
    return this.#repository.findById(id).pipe(
      filter((entity) => !!entity),
      map((entity: TenantEntity) => this.#mapToDto(entity)),
    );
  }

  #mapToEntity(
    dto: TenantCreateInput
  ): Pick<TenantEntity, 'name' | 'description'> {
    return {
      name: dto.name,
      description: dto.description ? dto.description : null,
    };
  }

  #mapToDto(entity: TenantEntity): TenantOutput {
    return new TenantOutput({
      id: entity.id,
      name: entity.name,
      description: entity.description ?? undefined,
    });
  }
}
