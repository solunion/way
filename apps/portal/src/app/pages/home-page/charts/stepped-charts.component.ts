import { Component, AfterViewInit } from '@angular/core';
import { Chart, registerables } from 'chart.js';
import { addSteppedLineBorderRadius } from './bar-border-radius';

Chart.register(...registerables); // Registra i moduli di Chart.js

@Component({
  selector: 'app-stepped-line-chart',
  template: '<canvas id="steppedLineChart"></canvas>',
  standalone: true
})
export class SteppedLineChartComponent implements AfterViewInit {
  ngAfterViewInit(): void {
    addSteppedLineBorderRadius(); // Applica la personalizzazione

    new Chart('steppedLineChart', {
      type: 'line',
      data: {
        labels: ['Lun', 'Mar', 'Mer', 'Gio', 'Ven', 'Sab', 'Dom'],
        datasets: [
          {
            label: 'Temperatura',
            data: [10, 15, 12, 20, 18, 25, 22],
            borderColor: 'rgba(54, 162, 235, 1)',
            borderWidth: 2,
            backgroundColor: 'rgba(54, 162, 235, 0.2)',
            fill: false,
            stepped: true, // Linea scalettata
            pointRadius: 8, // Raggio dei punti
            pointBorderWidth: 2, // Bordo dei punti
            pointBackgroundColor: 'rgba(255, 99, 132, 1)', // Colore interno dei punti
            pointBorderColor: 'rgba(255, 255, 255, 1)', // Bordo bianco per contrasto
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
          y: { beginAtZero: false },
        },
      },
    });
  }
}
