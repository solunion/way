import { Component } from '@angular/core';
import { TableComponent } from '../../shared/components/table/table.component';

@Component({
  selector: 'way-tenants-page',
  imports: [
    TableComponent
  ],
  templateUrl: './tenants-page.component.html',
  styleUrl: './tenants-page.component.scss',
  standalone: true
})
export class TenantsPageComponent {

}
