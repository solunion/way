import { UsePipes, ValidationPipe } from '@nestjs/common';
import { Args, Mutation, Query, Resolver } from '@nestjs/graphql';
import { plainToInstance } from 'class-transformer';
import { map, Observable } from 'rxjs';
import { CreateTenantDto } from './dto/create-tenant.dto';
import { TenantDto } from './dto/tenant.dto';
import { NewTenant } from './tenant.model';
import { TenantService } from './tenant.service';

@Resolver(() => TenantDto)
@UsePipes(new ValidationPipe({transform: true}))
export class TenantResolver {
  #service: TenantService;

  constructor(service: TenantService) {
    this.#service = service;
  }

  @Query(() => TenantDto)
  tenant(@Args('id') id: string): Observable<TenantDto> {
    return this.#service.findById$(id);
  }

  @Mutation(() => TenantDto)
  createTenant(@Args('tenant') request: CreateTenantDto): Observable<TenantDto> {
    const tenant = this.#service.create$(plainToInstance(NewTenant, request));
    return tenant.pipe(map((data: TenantDto) => plainToInstance(TenantDto, data)));
  }
}
