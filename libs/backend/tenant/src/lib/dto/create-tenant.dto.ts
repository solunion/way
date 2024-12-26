import { Field, InputType } from '@nestjs/graphql';
import { Expose } from 'class-transformer';
import { IsNotEmpty, MaxLength, MinLength } from 'class-validator';

@InputType()
export class CreateTenantDto {
  @Field()
  @Expose()
  @IsNotEmpty({ message: 'Name is required' })
  @MinLength(3, { message: 'Name must be at least 3 characters long' })
  @MaxLength(50, { message: 'Name must not exceed 50 characters' })
  name: string;

  @Expose()
  @Field({ nullable: true })
  @MaxLength(200, { message: 'Description must not exceed 200 characters' })
  description?: string;

  constructor(name: string, description?: string) {
    this.name = name;
    this.description = description;
  }
}