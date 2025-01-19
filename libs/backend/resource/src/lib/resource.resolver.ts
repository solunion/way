import { Args, ID, Mutation, Query, Resolver } from '@nestjs/graphql';
import { plainToInstance } from 'class-transformer';
import { Observable, map } from 'rxjs';
import { CreateResourceDto } from './dto/create-resource.dto';
import { ResourceDto } from './dto/resource.dto';
import { UpdateResourceDto } from './dto/update-resource.dto';
import { Resource } from './resource.model';
import { ResourceService } from './resource.service';

@Resolver(() => ResourceDto)
export class ResourceResolver {
  #service: ResourceService;

  constructor(service: ResourceService) {
    this.#service = service;
  }

  @Query(() => ResourceDto, { nullable: true })
  resource(@Args('id', { type: () => ID }) id: string): Observable<ResourceDto | null> {
    return this.#service
      .findById$(id)
      .pipe(map((resource) => (resource ? plainToInstance(ResourceDto, resource, { excludeExtraneousValues: true }) : null)));
  }

  @Query(() => [ResourceDto])
  resources(): Observable<ResourceDto[]> {
    return this.#service
      .findAll$()
      .pipe(map((resources) => resources.map((r) => plainToInstance(ResourceDto, r, { excludeExtraneousValues: true }))));
  }

  @Mutation(() => ResourceDto)
  createResource(@Args('input') input: CreateResourceDto): Observable<ResourceDto> {
    return this.#service
      .create$(plainToInstance(Resource, input, { excludeExtraneousValues: true }))
      .pipe(map((resource) => plainToInstance(ResourceDto, resource, { excludeExtraneousValues: true })));
  }

  @Mutation(() => ResourceDto)
  updateResource(@Args('id', { type: () => ID }) id: string, @Args('input') input: UpdateResourceDto): Observable<ResourceDto> {
    return this.#service
      .update$(id, plainToInstance(Resource, input, { excludeExtraneousValues: true }))
      .pipe(map((resource) => plainToInstance(ResourceDto, resource, { excludeExtraneousValues: true })));
  }

  @Mutation(() => Boolean)
  deleteResource(@Args('id', { type: () => ID }) id: string): Observable<boolean> {
    return this.#service.delete$(id).pipe(map(() => true));
  }
} 