import { UsePipes, ValidationPipe } from '@nestjs/common';
import { Args, Mutation, Query, Resolver } from '@nestjs/graphql';
import { plainToInstance } from 'class-transformer';
import { map, Observable, catchError, of } from 'rxjs';
import { ComponentService } from './component.service';
import { ComponentDto } from './dto/component.dto';
import { CreateComponentDto } from './dto/create-component.dto';
import { UpdateComponentDto } from './dto/update-component.dto';
import { Component } from './component.model';

@Resolver(() => ComponentDto)
@UsePipes(new ValidationPipe({ transform: true }))
export class ComponentResolver {
  #service: ComponentService;

  constructor(service: ComponentService) {
    this.#service = service;
  }

  @Query(() => ComponentDto, { nullable: true })
  getComponentById(@Args('id') id: string): Observable<ComponentDto | null> {
    return this.#service
      .findById$(id)
      .pipe(map((model) => (model ? this.#transformToDto(model) : null)));
  }

  @Query(() => [ComponentDto])
  getComponents(): Observable<ComponentDto[]> {
    return this.#service
      .findAll$()
      .pipe(map((models) => models.map((m) => this.#transformToDto(m))));
  }

  @Mutation(() => ComponentDto)
  createComponent(
    @Args('createComponent') createComponentDto: CreateComponentDto
  ): Observable<ComponentDto> {
    return this.#service
      .create$(this.#transformToModel(createComponentDto))
      .pipe(map((model) => this.#transformToDto(model)));
  }

  @Mutation(() => ComponentDto)
  updateComponent(
    @Args('id') id: string,
    @Args('updateComponent') updateComponentDto: UpdateComponentDto
  ): Observable<ComponentDto> {
    return this.#service
      .update$(id, this.#transformToModel(updateComponentDto))
      .pipe(map((model) => this.#transformToDto(model)));
  }

  @Mutation(() => Boolean)
  deleteComponent(@Args('id') id: string): Observable<boolean> {
    return this.#service.delete$(id).pipe(
      map(() => true),
      catchError(() => of(false))
    );
  }

  #transformToDto(model: Component): ComponentDto {
    return plainToInstance(ComponentDto, model, { excludeExtraneousValues: true });
  }

  #transformToModel(dto: ComponentDto | CreateComponentDto | UpdateComponentDto): Component {
    return plainToInstance(Component, dto, { excludeExtraneousValues: true });
  }
} 