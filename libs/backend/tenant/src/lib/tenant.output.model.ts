import { Field, ObjectType } from '@nestjs/graphql';

@ObjectType()
export class TenantOutput {
  @Field({ nullable: true })
  id?: string;

  @Field()
  name: string = '';

  @Field({ nullable: true })
  description?: string;

  constructor({ id, name, description }: TenantOutput) {
    this.id = id;
    this.name = name;
    this.description = description;
  }
}
