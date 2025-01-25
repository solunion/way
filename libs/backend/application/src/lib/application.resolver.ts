import { Resolver, Query, Mutation, Args } from '@nestjs/graphql';
import { plainToInstance } from 'class-transformer';
import { ApplicationService } from './application.service';
import { ApplicationDto } from './dto/application.dto';
import { CreateApplicationDto } from './dto/create-application.dto';
import { UpdateApplicationDto } from './dto/update-application.dto';
import { catchError, filter, map, Observable, of } from 'rxjs';
import { Application } from './application.model';

@Resolver(() => ApplicationDto)
export class ApplicationResolver {
  #service: ApplicationService;
  constructor(applicationService: ApplicationService) {
    this.#service = applicationService;
  }

  @Mutation(() => ApplicationDto)
  createApplication(@Args('createApplicationInput') createApplicationDto: CreateApplicationDto): Observable<ApplicationDto> {
    return this.#service.create$(this.#transformToModel(createApplicationDto)).pipe(
      map(model => this.#transformToDto(model)),
    );
  }

  @Query(() => [ApplicationDto])
  getApplications(): Observable<ApplicationDto[]> {
    return this.#service.findAll$().pipe(
      map(models => models.map(model => this.#transformToDto(model))),
    );
  }

  @Query(() => ApplicationDto, { nullable: true })
  getApplicationById(@Args('id') id: string): Observable<ApplicationDto | null> {
    return this.#service.findById$(id).pipe(
      filter(model => !!model),
      map(model => this.#transformToDto(model)),
    );
  }

  @Mutation(() => ApplicationDto)
  updateApplication(
    @Args('id') id: string,
    @Args('updateApplicationInput') updateApplicationDto: UpdateApplicationDto
  ): Observable<ApplicationDto> {
    return this.#service.update$(id, this.#transformToModel(updateApplicationDto)).pipe(
      map(model => this.#transformToDto(model)),
    );
  }

  @Mutation(() => Boolean)
  deleteApplication(@Args('id') id: string): Observable<boolean> {
    return this.#service.delete$(id).pipe(
      map(() => true),
      catchError(() => of(false))
    );
  }

  #transformToDto(model: Application): ApplicationDto {
    return plainToInstance(ApplicationDto, model, {excludeExtraneousValues: true});
  }

  #transformToModel(dto: ApplicationDto | CreateApplicationDto | UpdateApplicationDto): Application {
    return plainToInstance(Application, dto, {excludeExtraneousValues: true});
  }
}
