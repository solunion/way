import { Field, ObjectType } from '@nestjs/graphql';

@ObjectType()
export class Tenant {
  @Field({ nullable: true })
  id?: string;

  @Field()
  name: string = '';

  @Field({ nullable: true })
  description?: string;
}
