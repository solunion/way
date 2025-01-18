import { Field, InputType, ObjectType } from '@nestjs/graphql';
import { IsNotEmpty, IsOptional, IsUUID, MaxLength, MinLength } from 'class-validator';

@InputType()
@ObjectType('ApplicationOutput')
export class ApplicationDto {
  @Field()
  @IsUUID()
  id: string;

  @Field()
  @IsNotEmpty()
  @MinLength(3, { message: 'Name must be at least 3 characters long' })
  @MaxLength(50, { message: 'Name must not exceed 50 characters' })
  name: string;

  @Field({ nullable: true })
  @IsOptional()
  @MaxLength(200, { message: 'Description must not exceed 200 characters' })
  description?: string;

  @Field({ nullable: true })
  @IsOptional()
  @IsUUID()
  tenantId?: string;

  constructor(id: string, name: string, description: string, tenantId: string) {
    this.id = id;
    this.name = name;
    this.description = description;
    this.tenantId = tenantId;
  }
}
