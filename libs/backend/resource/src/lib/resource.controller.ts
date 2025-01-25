import { Body, Controller, Delete, Get, Param, Post, Put } from '@nestjs/common';
import { plainToInstance } from 'class-transformer';
import { Observable, map } from 'rxjs';
import { CreateResourceDto } from './dto/create-resource.dto';
import { ResourceDto } from './dto/resource.dto';
import { UpdateResourceDto } from './dto/update-resource.dto';
import { Resource } from './resource.model';
import { ResourceService } from './resource.service';

@Controller('resources')
export class ResourceController {
  #service: ResourceService;

  constructor(service: ResourceService) {
    this.#service = service;
  }

  @Get(':id')
  findById(@Param('id') id: string): Observable<ResourceDto | null> {
    return this.#service
      .findById$(id)
      .pipe(map((resource) => (resource ? plainToInstance(ResourceDto, resource, { excludeExtraneousValues: true }) : null)));
  }

  @Get()
  findAll(): Observable<ResourceDto[]> {
    return this.#service
      .findAll$()
      .pipe(map((resources) => resources.map((r) => plainToInstance(ResourceDto, r, { excludeExtraneousValues: true }))));
  }

  @Post()
  create(@Body() input: CreateResourceDto): Observable<ResourceDto> {
    return this.#service
      .create$(plainToInstance(Resource, input, { excludeExtraneousValues: true }))
      .pipe(map((resource) => plainToInstance(ResourceDto, resource, { excludeExtraneousValues: true })));
  }

  @Put(':id')
  update(@Param('id') id: string, @Body() input: UpdateResourceDto): Observable<ResourceDto> {
    return this.#service
      .update$(id, plainToInstance(Resource, input, { excludeExtraneousValues: true }))
      .pipe(map((resource) => plainToInstance(ResourceDto, resource, { excludeExtraneousValues: true })));
  }

  @Delete(':id')
  delete(@Param('id') id: string): Observable<void> {
    return this.#service.delete$(id);
  }
} 