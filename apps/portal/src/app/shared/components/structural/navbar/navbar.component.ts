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
  templateUrl: './navbar.component.html',
})
export class NavbarComponent {

}
