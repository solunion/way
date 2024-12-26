import { Field, InputType, ObjectType } from '@nestjs/graphql';
import { IsNotEmpty, MinLength, MaxLength } from 'class-validator';

@ObjectType()
@InputType()
export class TenantCreateInput {
  @Field()
  @IsNotEmpty({ message: 'Name is required' })
  @MinLength(3, { message: 'Name must be at least 3 characters long' })
  @MaxLength(50, { message: 'Name must not exceed 50 characters' })
  name: string;

  @Field({ nullable: true })
  @MaxLength(200, { message: 'Description must not exceed 200 characters' })
  description?: string;

  constructor(name: string, description?: string) {
    this.name = name;
    this.description = description;
  }
}
