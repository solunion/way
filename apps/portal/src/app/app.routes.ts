import { Route } from '@angular/router';

export const appRoutes: Route[] = [
  {
    path: '',
    loadComponent: () => import('./pages/tenants-page/tenants-page.component').then(m => m.TenantsPageComponent),
    data: {
      title: 'Home'
    }
  }
];
