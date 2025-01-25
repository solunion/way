import { HttpLink } from 'apollo-angular/http';
import { ApolloClientOptions, InMemoryCache } from '@apollo/client';
import { environment } from '../../../environments/environment';
import { APOLLO_OPTIONS } from 'apollo-angular';


export function createApollo(httpLink: HttpLink): ApolloClientOptions<any> {
  return {
    link: httpLink.create({ uri: environment.apiUrl }),
    cache: new InMemoryCache(),
  };
}


export const APOLLO_PROVIDER = {
  provide: APOLLO_OPTIONS,
  useFactory: createApollo,
  deps: [HttpLink],
}
