import { UsePipes, ValidationPipe } from '@nestjs/common';
import { Args, ID, Mutation, Query, Resolver } from '@nestjs/graphql';
import { plainToInstance } from 'class-transformer';
import { Observable, map, catchError, of } from 'rxjs';
import { CreateResourceDto } from './dto/create-resource.dto';
import { ResourceDto } from './dto/resource.dto';
import { UpdateResourceDto } from './dto/update-resource.dto';
import { Resource, NewResource } from './resource.model';
import { ResourceService } from './resource.service';

@Resolver(() => ResourceDto)
@UsePipes(new ValidationPipe({ transform: true }))
export class ResourceResolver {
  #service: ResourceService;

  constructor(service: ResourceService) {
    this.#service = service;
  }

  @Query(() => ResourceDto)
  getResourceById(@Args('id') id: string): Observable<ResourceDto> {
    return this.#service
      .findById$(id)
      .pipe(map((resource) => plainToInstance(ResourceDto, resource, { excludeExtraneousValues: true })));
  }

  @Query(() => [ResourceDto])
  getResources(): Observable<ResourceDto[]> {
    return this.#service
      .findAll$()
      .pipe(map((resources) => resources.map((r) => plainToInstance(ResourceDto, r, { excludeExtraneousValues: true }))));
  }

  @Mutation(() => ResourceDto)
  createResource(@Args('resource') request: CreateResourceDto): Observable<ResourceDto> {
    const resource = this.#service.create$(plainToInstance(NewResource, request));
    return resource.pipe(map((data: ResourceDto) => plainToInstance(ResourceDto, data)));
  }

  @Mutation(() => ResourceDto)
  updateResource(
    @Args('id') id: string,
    @Args('resource') request: UpdateResourceDto
  ): Observable<ResourceDto> {
    const resource = this.#service.update$(id, plainToInstance(Resource, request));
    return resource.pipe(map((data: ResourceDto) => plainToInstance(ResourceDto, data)));
  }

  @Mutation(() => Boolean)
  deleteResource(@Args('id') id: string): Observable<boolean> {
    return this.#service.delete$(id).pipe(
      map(() => true),
      catchError(() => of(false))
    );
  }
} 