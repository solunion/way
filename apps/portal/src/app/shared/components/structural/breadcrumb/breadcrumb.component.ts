import { Component, OnInit, signal } from '@angular/core';
import { ActivatedRoute, NavigationEnd, Router, RouterLink } from '@angular/router';
import { filter, map } from 'rxjs';
import { BreadcrumbItem } from '../../../models/breadcrumb-item.model';

@Component({
  selector: 'way-breadcrumb',
  imports: [
    RouterLink
  ],
  templateUrl: './breadcrumb.component.html',
  styleUrl: './breadcrumb.component.scss',
  standalone: true
})
export class BreadcrumbComponent implements OnInit {

  breadcrumb = signal<BreadcrumbItem[]>([])

  constructor(private router: Router, private activatedRoute: ActivatedRoute) {}

  ngOnInit() {
    this.router.events
      .pipe(
        filter((event) => event instanceof NavigationEnd),
        map(() => this.getDeepestChild(this.activatedRoute)), // Trova la route più profonda
        map((route) => route.snapshot.data) // Recupera i dati della route
      )
      .subscribe((data) => {
        if( !data ) {
          console.warn('Breadcrumb data not found');
          this.breadcrumb.set([]);
          return;
        }
        this.breadcrumb.set(data['breadcrumb']);
      });
  }

  private getDeepestChild(route: ActivatedRoute): ActivatedRoute {
    while (route.firstChild) {
      route = route.firstChild;
    }
    return route;
  }
}


