import { Body, Controller, Get, Param, Post } from '@nestjs/common';
import { Observable } from 'rxjs';
import { Tenant } from './tenant.model';
import { TenantService } from './tenant.service';

@Controller('tenants')
export class TenantController {
  constructor(private readonly service: TenantService) {
  }

  @Post()
  create(@Body() request: Tenant): Observable<Tenant> {
    return this.service.create(request);
  }

  @Get(':id')
  getOne(@Param('id') id: string): Observable<Tenant> {
    return this.service.getOne(id);
  }
}
