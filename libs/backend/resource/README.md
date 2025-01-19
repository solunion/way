# Resource Library

## Overview
Library for managing application resources in the Way system. A resource represents an application element (API, form, route, model) that can be subject to authorization.

## Features
- CRUD operations for resources
- Multi-tenancy support
- Soft delete
- GraphQL API

## Installation
This library is part of the Way monorepo and is installed automatically.

## Usage
```typescript
// Import the module
import { ResourceModule } from '@way/backend-resource';

@Module({
  imports: [ResourceModule],
})
export class AppModule {}

// Inject the service
constructor(private readonly resourceService: ResourceService) {}

// Use the service
this.resourceService.create$({
  name: 'My Resource',
  type: ResourceType.HTTP,
  componentId: 'uuid',
  tenantId: 'optional-tenant-uuid'
});
```

## API
### Types
```typescript
enum ResourceType {
  HTTP,
  ROUTE,
  FORM,
  MODEL
}
```

### GraphQL Operations
#### Queries
- `resource(id: ID!): Resource` - Get a resource by ID
- `resources: [Resource!]!` - Get all resources

#### Mutations
- `createResource(input: CreateResourceInput!): Resource!` - Create a new resource
- `updateResource(id: ID!, input: UpdateResourceInput!): Resource!` - Update a resource
- `deleteResource(id: ID!): Boolean!` - Delete a resource

## Development
- Run tests: `nx test backend-resource`
- Build: `nx build backend-resource` 