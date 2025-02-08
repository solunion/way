import { Component, inject, OnInit, signal } from '@angular/core';
import { PeriodicElement, TableComponent } from '../../shared/components/table/table.component';
import { MatExpansionModule } from '@angular/material/expansion';
import { MatChipsModule } from '@angular/material/chips';
import { MatIcon } from '@angular/material/icon';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { TenantService } from '../../shared/services/tenant.service';
import { Tenant } from '../../shared/models/tenant.model';
import { ApolloModule } from 'apollo-angular';
import { TENANTS_COLUMS } from './utils/tenants-columns.const';

@Component({
  selector: 'way-tenants-page',
  imports: [
    TableComponent,
    MatExpansionModule,
    MatChipsModule,
    MatIcon,
    MatFormFieldModule,
    MatInputModule
  ],
  templateUrl: './tenants-page.component.html',
  styleUrl: './tenants-page.component.scss',
  standalone: true
})
export class TenantsPageComponent implements OnInit {

  #tenantService= inject(TenantService);

  readonly panelOpenState = signal(false);

  tenants = signal<Tenant[]>([]);
  tableColumns = TENANTS_COLUMS;

  ngOnInit() {
    this.#tenantService.getAll$().subscribe(tenants => {
      this.tenants.set(tenants);
    })
  }

}
