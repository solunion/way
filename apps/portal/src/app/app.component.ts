import { Component, inject, OnInit, signal } from '@angular/core';
import { RouterModule } from '@angular/router';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatTooltipModule } from '@angular/material/tooltip';
import { NavbarComponent } from './shared/components/structural/navbar/navbar.component';
import { ActionMenuComponent } from './shared/components/structural/action-menu/action-menu.component';
import { ToolbarComponent } from './shared/components/structural/toolbar/toolbar.component';
import { UtilityService } from './shared/services/utility.service';
import { NgClass } from '@angular/common';
import { BreadcrumbComponent } from './shared/components/structural/breadcrumb/breadcrumb.component';

@Component({
  imports: [
    RouterModule,
    MatToolbarModule,
    MatIconModule,
    MatButtonModule,
    MatFormFieldModule,
    MatInputModule,
    MatTooltipModule,
    NavbarComponent,
    ActionMenuComponent,
    ToolbarComponent,
    NgClass,
    BreadcrumbComponent
  ],
  selector: 'way-root',
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss',
  standalone: true
})
export class AppComponent implements OnInit {
  //Injection
  #utilityService = inject(UtilityService);

  //States
  darkModeOn = signal( false );

  title = 'portal';

  /**
   *
   */
  ngOnInit() {
    this.#utilityService.darkModeOn$
      .subscribe((value : boolean) => {
        this.darkModeOn.set( value );
      });
  }
}
