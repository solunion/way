import { Field } from '@nestjs/graphql';
import { IsNotEmpty, IsUUID, MaxLength, MinLength } from 'class-validator';

export class NewTenant {
  @IsNotEmpty()
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

export class Tenant extends NewTenant {
  @IsUUID()
  id: string;

  constructor(id: string, name: string, description?: string) {
    super(name, description);
    this.id = id;
  }
}
