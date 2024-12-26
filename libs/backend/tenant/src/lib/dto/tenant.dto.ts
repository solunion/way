import { Field, ObjectType } from '@nestjs/graphql';

@ObjectType('Tenant')
export class TenantDto {
  @Field()
  id: string;

  @Field()
  name: string;

  @Field({ nullable: true })
  description?: string;

  constructor(id: string, name: string, description: string) {
    this.id = id;
    this.name = name;
    this.description = description;
  }
}
