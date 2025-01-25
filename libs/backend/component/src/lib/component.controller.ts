import { Body, Controller, Delete, Get, Param, Patch, Post, UsePipes, ValidationPipe } from '@nestjs/common';
import { plainToInstance } from 'class-transformer';
import { filter, map, Observable } from 'rxjs';
import { ComponentService } from './component.service';
import { ComponentDto } from './dto/component.dto';
import { CreateComponentDto } from './dto/create-component.dto';
import { UpdateComponentDto } from './dto/update-component.dto';
import { Component } from './component.model';

@Controller('components')
@UsePipes(new ValidationPipe({ transform: true }))
export class ComponentController {
  #service: ComponentService;

  constructor(service: ComponentService) {
    this.#service = service;
  }

  @Post()
  create$(@Body() createComponentDto: CreateComponentDto): Observable<ComponentDto> {
    return this.#service
      .create$(this.#transformToModel(createComponentDto))
      .pipe(map((model) => this.#transformToDto(model)));
  }

  @Get()
  findAll$(): Observable<ComponentDto[]> {
    return this.#service
      .findAll$()
      .pipe(map((models) => models.map((m) => this.#transformToDto(m))));
  }

  @Get(':id')
  findOne$(@Param('id') id: string): Observable<ComponentDto | null> {
    return this.#service
      .findById$(id)
      .pipe(
        filter((model) => !!model),
        map((model) => this.#transformToDto(model))
      );
  }

  @Patch(':id')
  update$(
    @Param('id') id: string,
    @Body() updateComponentDto: UpdateComponentDto
  ): Observable<ComponentDto> {
    return this.#service
      .update$(id, this.#transformToModel(updateComponentDto))
      .pipe(map((model) => this.#transformToDto(model)));
  }

  @Delete(':id')
  remove$(@Param('id') id: string): Observable<void> {
    return this.#service.delete$(id);
  }

  #transformToDto(model: Component): ComponentDto {
    return plainToInstance(ComponentDto, model, { excludeExtraneousValues: true });
  }

  #transformToModel(dto: ComponentDto | CreateComponentDto | UpdateComponentDto): Component {
    return plainToInstance(Component, dto, { excludeExtraneousValues: true });
  }
} 