import { Field, InputType } from '@nestjs/graphql';

@InputType()
export class TenantCreateInput {
  @Field()
  name: string;

  @Field()
  description?: string;

  constructor(name: string, description?: string) {
    this.name = name;
    this.description = description;
  }
}
