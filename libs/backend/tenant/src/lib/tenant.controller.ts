import { Body, Controller, Get, Param, Post } from '@nestjs/common';
import { Observable } from 'rxjs';
import { TenantDto } from './dto/tenant-dto.model';
import { TenantService } from './tenant.service';

@Controller('tenants')
export class TenantController {
  constructor(private readonly service: TenantService) {
  }

  @Post()
  create(@Body() request: TenantDto): Observable<TenantDto> {
    return this.service.create(request);
  }

  @Get(':id')
  getOne(@Param('id') id: string): Observable<TenantDto> {
    return this.service.getOne(id);
  }
}
