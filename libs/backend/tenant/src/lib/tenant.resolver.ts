import { Args, Query, Resolver } from '@nestjs/graphql';
import { Observable } from 'rxjs';
import { TenantDto } from './dto/tenant-dto.model';
import { TenantService } from './tenant.service';

@Resolver(() => TenantDto)
export class TenantResolver {
  constructor(private readonly service: TenantService) {}

  @Query(() => TenantDto)
  tenant(@Args('id') id: string): Observable<TenantDto> {
    return this.service.getOne(id);
  }
}
