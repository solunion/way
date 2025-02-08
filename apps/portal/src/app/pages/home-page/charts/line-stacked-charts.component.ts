import { Component, AfterViewInit } from '@angular/core';
import { Chart, registerables } from 'chart.js';
import { addStackedLineBorderRadius } from './bar-border-radius';

Chart.register(...registerables); // Registra i moduli di Chart.js

@Component({
  selector: 'app-stacked-line-chart',
  template: '<canvas id="stackedLineChart"></canvas>',
  standalone: true
})
export class StackedLineChartComponent implements AfterViewInit {
  ngAfterViewInit(): void {
    addStackedLineBorderRadius(); // Applica la personalizzazione

    new Chart('stackedLineChart', {
      type: 'line',
      data: {
        labels: ['Gen', 'Feb', 'Mar', 'Apr', 'Mag', 'Giu'],
        datasets: [
          {
            label: 'Serie A',
            data: [10, 20, 15, 25, 30, 40],
            borderColor: 'rgba(54, 162, 235, 1)',
            backgroundColor: 'rgba(54, 162, 235, 0.2)',
            borderWidth: 2,
            fill: true, // Area sotto la linea
            pointRadius: 8, // Raggio dei punti
            pointBorderWidth: 2,
            pointBackgroundColor: 'rgba(54, 162, 235, 1)',
            pointBorderColor: 'rgba(255, 255, 255, 1)',
          },
          {
            label: 'Serie B',
            data: [5, 15, 10, 20, 25, 35],
            borderColor: 'rgba(255, 99, 132, 1)',
            backgroundColor: 'rgba(255, 99, 132, 0.2)',
            borderWidth: 2,
            fill: true, // Stack delle aree
            pointRadius: 8,
            pointBorderWidth: 2,
            pointBackgroundColor: 'rgba(255, 99, 132, 1)',
            pointBorderColor: 'rgba(255, 255, 255, 1)',
          },
        ],
      },
      options: {
        responsive: true,
        plugins: {
          legend: { display: true },
        },
        scales: {
          x: { grid: { display: false } },
          y: {
            stacked: true, // Abilita stacking
            beginAtZero: true,
          },
        },
      },
    });
  }
}
