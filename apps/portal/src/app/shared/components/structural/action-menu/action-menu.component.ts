import { Component } from '@angular/core';
import { MatIconButton } from '@angular/material/button';

@Component({
  selector: 'way-action-menu',
  imports: [
    MatIconButton
  ],
  standalone: true,
  template: `
    <nav aria-label="Dock menu" class="way-action-menu curved-corner ">
      <button mat-icon-button color="primary" aria-label="Go to Home">
        <span class="material-icons">home</span>
      </button>
      <button mat-icon-button color="primary" aria-label="Learn more About us">
        <span class="material-icons">info</span>
      </button>
      <button mat-icon-button color="primary" aria-label="Explore our Services">
        <span class="material-icons">build</span>
      </button>
      <button mat-icon-button color="primary" aria-label="Get in touch with us">
        <span class="material-icons">contact_mail</span>
      </button>
    </nav>
  `,
})
export class ActionMenuComponent {

}
