import { Field, InputType, ObjectType, registerEnumType } from '@nestjs/graphql';
import { Expose } from 'class-transformer';
import { IsNotEmpty, IsUUID, MaxLength, MinLength, IsEnum } from 'class-validator';
import { ResourceType } from '../resource.model';

@InputType()
@ObjectType('ResourceOutput')
export class ResourceDto {
  @Field()
  @IsUUID()
  id: string;

  @Field()
  @Expose()
  @IsNotEmpty({ message: 'Name is required' })
  @MinLength(3, { message: 'Name must be at least 3 characters long' })
  @MaxLength(50, { message: 'Name must not exceed 50 characters' })
  name: string;

  @Field(() => ResourceType)
  @Expose()
  @IsEnum(ResourceType)
  type: ResourceType;

  @Field()
  @Expose()
  @IsUUID()
  componentId: string;

  @Field({ nullable: true })
  @Expose()
  @IsUUID()
  tenantId?: string;

  constructor(id: string, name: string, type: ResourceType, componentId: string, tenantId?: string) {
    this.id = id;
    this.name = name;
    this.type = type;
    this.componentId = componentId;
    this.tenantId = tenantId;
  }
}

registerEnumType(ResourceType, {
  name: 'ResourceType',
});
