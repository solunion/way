# Resource Library

## Overview
Library for managing application resources in the Way system. A resource represents an application element (API, form, route, model) that can be subject to authorization.

## Features
- CRUD operations for resources
- Multi-tenancy support
- Soft delete
- GraphQL and REST APIs

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
  HTTP = 'HTTP',
  ROUTE = 'ROUTE',
  FORM = 'FORM',
  MODEL = 'MODEL'
}
```

### GraphQL API
#### Queries
- `getResourceById(id: String!): Resource` - Get a resource by ID
- `getResources: [Resource!]!` - Get all resources

#### Mutations
- `createResource(resource: CreateResourceInput!): Resource!` - Create a new resource
- `updateResource(id: String!, resource: UpdateResourceInput!): Resource!` - Update a resource
- `deleteResource(id: String!): Boolean!` - Delete a resource (returns false if fails)

### REST API
#### Endpoints
- `GET /resources/:id` - Get a resource by ID
- `GET /resources` - Get all resources
- `POST /resources` - Create a new resource
- `PUT /resources/:id` - Update a resource
- `DELETE /resources/:id` - Delete a resource

## Development
- Run tests: `nx test backend-resource`
- Build: `nx build backend-resource` 
