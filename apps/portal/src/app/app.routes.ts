import { Route } from '@angular/router';

export const appRoutes: Route[] = [
  {
    path: '',
    loadComponent: () => import('./pages/home-page/home-page.component').then(m => m.HomePageComponent),
    data: {
      title: 'Home'
    }
  },
  {
    path: 'tenants',
    loadComponent: () => import('./pages/tenants-page/tenants-page.component').then(m => m.TenantsPageComponent),
    data: {
      title: 'Tenants'
    }
  }
];
