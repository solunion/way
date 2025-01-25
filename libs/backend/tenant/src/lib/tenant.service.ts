import { entityNameToValue } from '@angular/compiler-cli/src/ngtsc/reflection';
import { Injectable, UsePipes, ValidationPipe } from '@nestjs/common';
import { plainToInstance } from 'class-transformer';
import { catchError, filter, map, Observable, throwError } from 'rxjs';
import { TenantDto } from './dto/tenant.dto';
import { TenantEntity } from './tenant.entity';
import { NewTenant, Tenant } from './tenant.model';
import { TenantRepository } from './tenant.repository';

@Injectable()
@UsePipes(new ValidationPipe({ transform: true }))
export class TenantService {
  #repository: TenantRepository;

  constructor(repository: TenantRepository) {
    this.#repository = repository;
  }

  create$(newTenant: NewTenant): Observable<Tenant> {
    return this.#repository.create$(this.#transformToEntity(newTenant)).pipe(
      map((entity: TenantEntity) => this.#transformToDto(entity)),
      catchError((error) => {
        console.error('Error creating tenant:', error);
        return throwError(() => new Error('Unable to create tenant'));
      })
    );
  }

  findAll$(): Observable<Tenant[]> {
    return this.#repository.findAll$().pipe(
      map((entities) => entities.map(entity => this.#transformToDto(entity)))
    );
  }

  findById$(id: string): Observable<TenantDto> {
    return this.#repository.findById$(id).pipe(
      filter((entity) => !!entity),
      map((entity: TenantEntity) => this.#transformToDto(entity))
    );
  }

  update$(id: string, updateTenant: Partial<Tenant>): Observable<Tenant> {
    return this.#repository.update$(id, this.#transformToEntity(updateTenant)).pipe(
      map((entity: TenantEntity) => this.#transformToDto(entity)),
      catchError((error) => {
        console.error('Error updating tenant:', error);
        return throwError(() => new Error('Unable to update tenant'));
      })
    );
  }

  delete$(id: string): Observable<void> {
    return this.#repository.softDelete$(id).pipe(
      map(() => undefined),
      catchError((error) => {
        console.error('Error deleting tenant:', error);
        return throwError(() => new Error('Unable to delete tenant'));
      })
    );
  }

  #transformToEntity(dto: Partial<Tenant>): Pick<TenantEntity, 'name' | 'description'> {
    return plainToInstance(TenantEntity, dto, {excludeExtraneousValues: true});
  }

  #transformToDto(entity: TenantEntity): Tenant {
    return plainToInstance(Tenant, entity, {excludeExtraneousValues: true});
  }
}
