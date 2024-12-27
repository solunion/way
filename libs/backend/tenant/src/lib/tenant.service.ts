import { Injectable, UsePipes, ValidationPipe } from '@nestjs/common';
import { TenantEntity } from '@prisma/client';
import { plainToInstance } from 'class-transformer';
import { catchError, filter, map, mapTo, Observable, throwError } from 'rxjs';
import { TenantDto } from './dto/tenant.dto';
import { NewTenant, Tenant } from './tenant.model';
import { TenantRepository } from './tenant.repository';

@Injectable()
@UsePipes(new ValidationPipe({transform: true}))
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
      mapTo(void 0),
      catchError((error) => {
        console.error('Error deleting tenant:', error);
        return throwError(() => new Error('Unable to delete tenant'));
      })
    );
  }

  #transformToEntity(dto: Partial<Tenant>): Pick<TenantEntity, 'name' | 'description'> {
    // @ts-expect-error Generated type by Prisma
    return plainToInstance(TenantEntity, dto);
  }

  #transformToDto(entity: TenantEntity): Tenant {
    return plainToInstance(Tenant, entity);
  }
}
