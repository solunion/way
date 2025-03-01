import { Component, inject, OnInit } from '@angular/core';
import { MatExpansionModule } from '@angular/material/expansion';
import { MatChipsModule } from '@angular/material/chips';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { UtilityService } from '../../../../shared/services/utility.service';
import { FormBuilder, FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Tenant } from '../../../../shared/models/tenant.model';
import { Validator } from 'class-validator';
import { FormErrorPipe } from '../../../../shared/pipes/form-error.pipe';

interface TenantForm {
  id: null | FormControl<string>;
  name: FormControl<string>;
  description: FormControl<string>;
}

@Component({
  imports: [
    MatExpansionModule,
    MatChipsModule,
    MatFormFieldModule,
    MatInputModule,
    ReactiveFormsModule,
    FormErrorPipe
  ],
  templateUrl: './tenants-detail-page.component.html',
  standalone: true
})
export class TenantsDetailPageComponent implements OnInit {

  #utilityService = inject(UtilityService);
  #formBuilder = inject(FormBuilder);

  tenantForm: any;

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

  constructor() {
    this.tenantForm = this.#formBuilder.group<TenantForm>(
      {
        id : null,
        name: new FormControl<string>(
          '',
          {
            nonNullable: true,
            validators: [ Validators.required, Validators.minLength(3)]
          }
        ),
        description: new FormControl<string>(
          '',
          {
            nonNullable: true,
            validators: [ Validators.required ]
          }
        ),
      }
    );
  }


  ngOnInit() {
    this.#utilityService.showActionMenu( { ...this.unselectedItemsConfig, selectedElements: 0 });
  }

}
