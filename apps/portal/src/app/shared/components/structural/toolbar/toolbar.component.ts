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
  template: `
    <mat-toolbar color="primary" class="way-toolbar">

      <!-- Menu button -->
      <div>
        <button mat-icon-button class="example-icon" aria-label="Example icon-button with menu icon">
          <mat-icon>menu</mat-icon>
        </button>
      </div>

      <!-- Search Bar -->
      <mat-form-field appearance="outline">
        <mat-label>Cerca</mat-label>
        <input matInput placeholder="Inserisci una parola chiave">
      </mat-form-field>


      <!-- User Opt-->
      <div>
        <button mat-icon-button
                class="example-icon favorite-icon"
                aria-label="Example icon-button with heart icon"
                [matMenuTriggerFor]="menu"
        >
          <mat-icon>person</mat-icon>
        </button>

        <mat-menu #menu="matMenu">

          <mat-list role="list">

            <mat-list-item role="listitem">
              <mat-slide-toggle
                labelPosition="before"
                [checked]="darkModeOn()"
                (change)="setDarkMode( $event )"
              >
                <span style="padding-right: 10px">Dark Mode</span>
              </mat-slide-toggle>
            </mat-list-item>

            <mat-list-item role="listitem">
              <button mat-button>
                <mat-icon>logout</mat-icon>
                Logout
              </button>
            </mat-list-item>

          </mat-list>
        </mat-menu>


      </div>
    </mat-toolbar>
  `,
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
