import { Route } from '@angular/router';

export enum PathRoutes {
  HOME = '',
  TENANTS = 'tenants',
  TENANT_NEW = 'tenants/detail',
  TENANT_MODIFY = 'tenants/detail/:id',
}


export const appRoutes: Route[] = [
  {
    path: PathRoutes.HOME,
    loadComponent: () => import('./pages/home-page/home-page.component').then(m => m.HomePageComponent),
    data: {
      title: 'Home',
      breadcrumb: []
    }
  },
  {
    path: PathRoutes.TENANTS,
    loadComponent: () => import('./pages/tenants-page/tenants-page.component').then(m => m.TenantsPageComponent),
    data: {
      title: 'Tenants',
      breadcrumb: [
        {
          label: 'Tenants',
          path: PathRoutes.TENANTS
        }
      ]
    }
  },
  {
    path: PathRoutes.TENANT_NEW,
    loadComponent: () => import('./pages/tenants-page/pages/tenants-detail-page/tenants-detail-page.component').then(m => m.TenantsDetailPageComponent),
    data: {
      title: 'Tenants',
      breadcrumb: [
        {
          label: 'Tenants',
          path: PathRoutes.TENANTS
        },
        {
          label: 'New Tenant',
          path: PathRoutes.TENANT_NEW
        }
      ]
    }
  }
];
