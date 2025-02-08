import { Component, AfterViewInit } from '@angular/core';
import { Chart, registerables } from 'chart.js';
import { addRadarBorderRadius } from './bar-border-radius';

Chart.register(...registerables); // Registra i moduli di Chart.js

@Component({
  selector: 'app-radar-chart',
  template: '<canvas id="radarChart"></canvas>',
  standalone: true
})
export class RadarChartComponent implements AfterViewInit {
  ngAfterViewInit(): void {
    addRadarBorderRadius(); // Applica la personalizzazione

    new Chart('radarChart', {
      type: 'radar',
      data: {
        labels: ['Forza', 'Velocità', 'Resistenza', 'Agilità', 'Precisione'],
        datasets: [
          {
            label: 'Atleta A',
            data: [80, 60, 90, 70, 85],
            fill: true,
            backgroundColor: 'rgba(54, 162, 235, 0.2)',
            borderColor: 'rgba(54, 162, 235, 1)',
            borderWidth: 2,
            pointRadius: 10, // Imposta il raggio dei punti
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
          r: {
            beginAtZero: true,
            grid: { color: 'rgba(200, 200, 200, 0.3)' },
            angleLines: { color: 'rgba(200, 200, 200, 0.3)' },
          },
        },
      },
    });
  }
}
