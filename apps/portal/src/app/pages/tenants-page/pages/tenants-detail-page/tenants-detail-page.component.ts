import { Component, inject, OnInit } from '@angular/core';
import { MatExpansionModule } from '@angular/material/expansion';
import { MatChipsModule } from '@angular/material/chips';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { UtilityService } from '../../../../shared/services/utility.service';

@Component({
  imports: [
    MatExpansionModule,
    MatChipsModule,
    MatFormFieldModule,
    MatInputModule
  ],
  templateUrl: './tenants-detail-page.component.html',
  standalone: true
})
export class TenantsDetailPageComponent implements OnInit {

  #utilityService = inject(UtilityService);

  unselectedItemsConfig = {
    actions: [
      {
        id: 'add',
        label: 'Add',
        icon: 'add',
        action: () => {
          console.log('Add');
        }
      }
    ]
  }


  ngOnInit() {
    this.#utilityService.showActionMenu( { ...this.unselectedItemsConfig, selectedElements: 0 });
  }

}
