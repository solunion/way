import { ApolloServerPluginLandingPageLocalDefault } from '@apollo/server/plugin/landingPage/default';
import { ApolloDriver, ApolloDriverConfig } from '@nestjs/apollo';
import { Module } from '@nestjs/common';
import { GraphQLModule } from '@nestjs/graphql';
import { ApplicationModule } from '@way/backend-application';
import { ComponentModule } from '@way/backend-component';
import { ResourceModule } from '@way/backend-resource';
import { RuleModule } from '@way/backend-rule';
import { TenantModule } from '@way/backend-tenant';
import GraphQLJSON from 'graphql-type-json';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { join } from 'path';

@Module({
  imports: [
    GraphQLModule.forRoot<ApolloDriverConfig>({
      driver: ApolloDriver,
      resolvers: { JSON: GraphQLJSON },
      autoSchemaFile: join(process.cwd(), 'schema.gql'),
      sortSchema: true,
      playground: false,
      plugins: [ApolloServerPluginLandingPageLocalDefault()],

    }),
    TenantModule,
    RuleModule,
    ApplicationModule,
    ComponentModule,
    ResourceModule
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
