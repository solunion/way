import { Body, Controller, Get, Param, Post } from '@nestjs/common';
import { Observable } from 'rxjs';
import { TenantCreateInput } from './tenant-create.input.model';
import { TenantOutput } from './tenant.output.model';
import { TenantService } from './tenant.service';

@Controller('tenants')
export class TenantController {
  #service: TenantService;

  constructor(service: TenantService) {
    this.#service = service;
  }

  @Post()
  create$(@Body() request: TenantCreateInput): Observable<TenantOutput> {
    return this.#service.create$(request);
  }

  @Get(':id')
  getOne$(@Param('id') id: string): Observable<TenantOutput> {
    return this.#service.getOne$(id);
  }
}
