import { Injectable } from '@angular/core';
import { ApolloClient, ApolloQueryResult } from '@apollo/client';
import { Apollo } from 'apollo-angular';
import { Tenant_GetAll_Query } from '../utils/graph-ql/queries/tentant-get-all.graphql-query';
import { Tenant } from '../models/tenant.model';
import { map, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class TenantService {

  constructor(
    private readonly apollo: Apollo
  ) {
  }

  /**
   * @description This method will be used to get all tenants.
   * @returns {Observable<any>} - An observable that will emit the tenants.
   */
  getAll$(): Observable<Tenant[]> {
    return this.apollo.query({
      query: Tenant_GetAll_Query,
      variables: {} as any
    }).pipe(
      map( (response: ApolloQueryResult<unknown>) => (response.data as { getTenants: Tenant[] }).getTenants || [] )
    )
  }

}
