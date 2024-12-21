import { Body, Controller, Post } from '@nestjs/common';
import { Observable } from 'rxjs';
import { TenantDto } from './dto/tenant-dto.model';
import { TenantService } from './tenant.service';

@Controller('tenants')
export class TenantRestController {
  constructor(private readonly service: TenantService) {
  }

  @Post()
  create(@Body() request: TenantDto): Observable<TenantDto> {
    return this.service.create(request);
  }
}
