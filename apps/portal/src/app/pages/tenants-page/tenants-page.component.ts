import { Component, inject, OnInit, signal } from '@angular/core';
import { PeriodicElement, TableComponent } from '../../shared/components/table/table.component';
import { MatExpansionModule } from '@angular/material/expansion';
import { MatChipsModule } from '@angular/material/chips';
import { MatIcon } from '@angular/material/icon';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { TenantService } from '../../shared/services/tenant.service';
import { Tenant } from '../../shared/models/tenant.model';
import { TENANTS_COLUMS } from './utils/tenants-columns.const';
import { UtilityService } from '../../shared/services/utility.service';
import { Router } from '@angular/router';

@Component({
  imports: [
    TableComponent,
    MatExpansionModule,
    MatChipsModule,
    MatIcon,
    MatFormFieldModule,
    MatInputModule
  ],
  templateUrl: './tenants-page.component.html',
  styleUrl: './tenants-page.component.html',
  standalone: true
})
export class TenantsPageComponent implements OnInit {

  #tenantService= inject(TenantService);
  #utilityService = inject(UtilityService);
  #router = inject(Router);

  readonly panelOpenState = signal(false);

  tenants = signal<Tenant[]>([]);
  tableColumns = TENANTS_COLUMS;

  selectedItemsConfig = {
    selectedElements: 0,
    actions: [
      {
        id: 'delete',
        label: 'Delete',
        icon: 'delete',
        action: () => {
          console.log('Delete');
        }
      }
    ]
  }

  unselectedItemsConfig = {
    actions: [
      {
        id: 'add',
        label: 'Add',
        icon: 'add',
        action: () => {
          this.#router.navigate(['tenants', 'detail'])
        }
      }
    ]
  }


  ngOnInit() {
    this.#tenantService.getAll$().subscribe(tenants => {
      this.tenants.set(tenants);
    })

    this.#utilityService.showActionMenu( { ...this.unselectedItemsConfig, selectedElements: 0 });
  }

  selectionChange(event: PeriodicElement[]){

    if( event.length > 0 ) {
      this.#utilityService.showActionMenu( { ...this.selectedItemsConfig, selectedElements: event.length });
    } else {
      this.#utilityService.showActionMenu( { ...this.unselectedItemsConfig, selectedElements: 0 });
    }
  }

}
