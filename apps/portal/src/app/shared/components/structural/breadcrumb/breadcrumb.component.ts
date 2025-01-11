import { Component } from '@angular/core';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'way-breadcrumb',
  imports: [
    RouterLink
  ],
  templateUrl: './breadcrumb.component.html',
  styleUrl: './breadcrumb.component.scss',
  standalone: true
})
export class BreadcrumbComponent {

}
