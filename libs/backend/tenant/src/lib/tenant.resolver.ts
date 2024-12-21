import { Args, Query, Resolver } from '@nestjs/graphql';
import { Observable } from 'rxjs';
import { Tenant } from './tenant.model';
import { TenantService } from './tenant.service';

@Resolver(() => Tenant)
export class TenantResolver {
  constructor(private readonly service: TenantService) {}

  @Query(() => Tenant)
  tenant(@Args('id') id: string): Observable<Tenant> {
    return this.service.getOne(id);
  }
}
