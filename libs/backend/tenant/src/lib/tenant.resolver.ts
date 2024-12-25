import { Args, Mutation, Query, Resolver } from '@nestjs/graphql';
import { Observable } from 'rxjs';
import { TenantCreateInput } from './tenant-create.input.model';
import { TenantOutput } from './tenant.output.model';
import { TenantService } from './tenant.service';

@Resolver(() => TenantOutput)
export class TenantResolver {
  #service: TenantService;

  constructor(service: TenantService) {
    this.#service = service;
  }

  @Query(() => TenantOutput)
  tenant(@Args('id') id: string): Observable<TenantOutput> {
    return this.#service.getOne$(id);
  }

  @Mutation(() => TenantOutput)
  createTenant(@Args('tenant') tenant: TenantCreateInput): Observable<TenantOutput> {
    return this.#service.create$(tenant);
  }
}
