import { Component } from '@angular/core';
import { MatIcon } from '@angular/material/icon';
import { MatIconButton } from '@angular/material/button';
import { MatTooltip } from '@angular/material/tooltip';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'way-navbar',
  imports: [
    MatIcon,
    MatIconButton,
    MatTooltip,
    RouterLink
  ],
  standalone: true,
  templateUrl: './navbar.component.html',
})
export class NavbarComponent {

  links = [
    {
      path: '/',
      label: 'Home',
      icon: 'home'
    },
    {

      path: '/tenants',
      label: 'Tenants',
      icon: 'dns'
    }
  ];

}
