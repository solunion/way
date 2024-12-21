import { Field, ObjectType } from '@nestjs/graphql';
// import { v4 as uuidv4 } from 'uuid';

@ObjectType()
export class TenantDto {
  // FIXME: convert to UUIDv4
  @Field(type => String, {nullable: true})
  id?: string;

  @Field(type => String)
  name: string = '';

  @Field(type => String, {nullable: true})
  description?: string;
}
