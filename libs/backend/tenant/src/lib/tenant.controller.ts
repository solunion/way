import { Body, Controller, Get, Param, Post, UsePipes, ValidationPipe, Patch } from '@nestjs/common';
import { plainToInstance } from 'class-transformer';
import { map, Observable } from 'rxjs';
import { CreateTenantDto } from './dto/create-tenant.dto';
import { TenantDto } from './dto/tenant.dto';
import { NewTenant } from './tenant.model';
import { TenantService } from './tenant.service';
import { UpdateTenantDto } from './dto/update-tenant.dto';
import { Tenant } from './tenant.model';

@Controller('tenants')
@UsePipes(new ValidationPipe({transform: true}))
export class TenantController {
  #service: TenantService;

  constructor(service: TenantService) {
    this.#service = service;
  }

  @Post()
  create$(@Body() request: CreateTenantDto): Observable<TenantDto> {
    const tenant = this.#service.create$(plainToInstance(NewTenant, request));
    return tenant.pipe(map((data: TenantDto) => plainToInstance(TenantDto, data)));
  }

  @Get(':id')
  getOne$(@Param('id') id: string): Observable<TenantDto> {
    return this.#service.findById$(id);
  }

  @Patch(':id')
  update$(@Param('id') id: string, @Body() request: UpdateTenantDto): Observable<TenantDto> {
    const tenant = this.#service.update$(id, plainToInstance(Tenant, request));
    return tenant.pipe(map((data: TenantDto) => plainToInstance(TenantDto, data)));
  }
}
