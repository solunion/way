import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatCardModule } from '@angular/material/card';
import { BarChartsComponent } from './charts/bar-charts.component';
import { SteppedLineChartComponent } from './charts/stepped-charts.component';
import { StackedLineChartComponent } from './charts/line-stacked-charts.component';

@Component({
  selector: 'way-home-page',
  standalone: true,
  imports: [CommonModule, MatCardModule, BarChartsComponent, SteppedLineChartComponent, StackedLineChartComponent],
  templateUrl: './home-page.component.html',
  styleUrl: './home-page.component.scss',
})
export class HomePageComponent {}
