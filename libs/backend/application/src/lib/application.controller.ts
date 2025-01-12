import { Controller, Post, Body, Get, Param, Patch, Delete } from '@nestjs/common';
import { plainToInstance } from 'class-transformer';
import { ApplicationService } from './application.service';
import { ApplicationDto } from './dto/application.dto';
import { CreateApplicationDto } from './dto/create-application.dto';
import { UpdateApplicationDto } from './dto/update-application.dto';
import { filter, map, Observable } from 'rxjs';
import { Application } from './application.model';

@Controller('applications')
export class ApplicationController {
  #service: ApplicationService;
  constructor(applicationService: ApplicationService) {
    this.#service = applicationService;
  }

  @Post()
  create(@Body() createApplicationDto: CreateApplicationDto): Observable<ApplicationDto> {
    return this.#service.create$(this.#transformToModel(createApplicationDto)).pipe(
      map(model => this.#transformToDto(model)),
    );
  }

  @Get()
  findAll(): Observable<ApplicationDto[]> {
    return this.#service.findAll$().pipe(
      map(models => models.map(model => this.#transformToDto(model))),
    );
  }

  @Get(':id')
  findOne(@Param('id') id: string): Observable<ApplicationDto | null> {
    return this.#service.findById$(id).pipe(
      filter(model => !!model),
      map(model => this.#transformToDto(model)),
    );
  }

  @Patch(':id')
  update(
    @Param('id') id: string,
    @Body() updateApplicationDto: UpdateApplicationDto
  ): Observable<Application> {
    return this.#service.update$(id, updateApplicationDto).pipe(
      map(model => this.#transformToDto(model)),
    );
  }

  @Delete(':id')
  remove(@Param('id') id: string): Observable<void> {
    return this.#service.delete$(id);
  }

  #transformToDto(model: Application): ApplicationDto {
    return plainToInstance(ApplicationDto, model, {excludeExtraneousValues: true});
  }

  #transformToModel(dto: ApplicationDto | CreateApplicationDto | UpdateApplicationDto): Application {
    return plainToInstance(Application, dto, {excludeExtraneousValues: true});
  }
}
