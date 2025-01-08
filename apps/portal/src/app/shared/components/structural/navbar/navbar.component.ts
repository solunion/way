import { Component } from '@angular/core';
import { MatIcon } from '@angular/material/icon';
import { MatIconButton } from '@angular/material/button';
import { MatTooltip } from '@angular/material/tooltip';

@Component({
  selector: 'way-navbar',
  imports: [
    MatIcon,
    MatIconButton,
    MatTooltip
  ],
  standalone: true,
  template: `
    <nav aria-label="Main menu" class="way-menu">
      <button mat-icon-button color="primary" aria-label="Go to Home"
              matTooltip="Example" matTooltipPosition="right"
              matTooltipClass="way-menu-tooltip"><mat-icon>view_comfy_alt</mat-icon></button>

      <button mat-icon-button color="primary" aria-label="Go to Home"
              matTooltip="Example" matTooltipPosition="right"
              matTooltipClass="way-menu-tooltip"><mat-icon>view_comfy_alt</mat-icon></button>


      <div class="way-menu__divisor"></div>
      <button mat-icon-button color="primary" aria-label="Go to Home"
              matTooltip="Example" matTooltipPosition="right"
              matTooltipClass="way-menu-tooltip"><mat-icon>view_comfy_alt</mat-icon></button>
    </nav>
  `
})
export class NavbarComponent {

}
