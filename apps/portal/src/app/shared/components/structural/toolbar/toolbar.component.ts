import { Component, inject, OnInit, signal } from '@angular/core';
import { MatFormField, MatLabel } from '@angular/material/form-field';
import { MatIcon } from '@angular/material/icon';
import { MatButton, MatIconButton } from '@angular/material/button';
import { MatInput } from '@angular/material/input';
import { MatToolbar } from '@angular/material/toolbar';
import { MatMenuModule } from '@angular/material/menu';
import { MatListModule } from '@angular/material/list';
import { MatSlideToggle, MatSlideToggleChange } from '@angular/material/slide-toggle';
import { UtilityService } from '../../../services/utility.service';

@Component({
  selector: 'way-toolbar',
  imports: [
    MatFormField,
    MatIcon,
    MatIconButton,
    MatInput,
    MatLabel,
    MatToolbar,
    MatMenuModule,
    MatListModule,
    MatButton,
    MatSlideToggle
  ],
  templateUrl: './toolbar.component.html',
  styles: ``,
  standalone: true
})
export class ToolbarComponent implements OnInit {
  //Injection
  #utilityService = inject(UtilityService);

  //States
  darkModeOn = signal( false );

  /**
   *
   */
  ngOnInit() {
    this.#utilityService.darkModeOn$.subscribe((value : boolean) => {
      this.darkModeOn.set( value );
    });
  }

  /**
   *
   */
  setDarkMode( state: MatSlideToggleChange ) {
    this.#utilityService.setDarkMode(state.checked);
  }
}
