import { Component, AfterViewInit } from '@angular/core';
import { Chart, registerables } from 'chart.js';
import { addBarBorderRadius } from './bar-border-radius';

Chart.register(...registerables); // Registra i moduli di Chart.js

@Component({
  selector: 'app-bar-chart',
  template: '<canvas id="barChart"></canvas>',
  standalone: true
})
export class BarChartsComponent implements AfterViewInit {
  ngAfterViewInit(): void {
    addBarBorderRadius(); // Applica la personalizzazione

    new Chart('barChart', {
      type: 'bar',
      data: {
        labels: ['Gen', 'Feb', 'Mar', 'Apr', 'Mag'],
        datasets: [
          {
            label: 'Vendite',
            data: [12, 19, 3, 5, 2],
            backgroundColor: 'rgba(54, 162, 235, 0.6)',
            borderColor: 'rgba(54, 162, 235, 1)',
            borderWidth: 2,
            borderRadius: 20, // Arrotonda gli angoli delle barre
          },
        ],
      },
      options: {
        responsive: true,
        plugins: {
          legend: { display: false },
        },
        scales: {
          x: { grid: { display: false } },
          y: { beginAtZero: true },
        },
      },
    });
  }
}
