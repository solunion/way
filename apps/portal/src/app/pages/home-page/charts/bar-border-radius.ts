import { Chart } from 'chart.js';

export function addBarBorderRadius() {
  Chart.defaults.elements.bar.borderRadius = 10; // Imposta il border radius di default a 10px
}

export function addRadarBorderRadius() {
  Chart.defaults.elements.point.radius = 8; // Imposta il raggio dei punti
  Chart.defaults.elements.point.borderWidth = 2; // Aggiunge un bordo ai punti
  Chart.defaults.elements.point.borderColor = 'rgba(54, 162, 235, 1)'; // Colore del bordo
  Chart.defaults.elements.point.backgroundColor = 'rgba(54, 162, 235, 0.6)'; // Sfondo dei punti
}

export function addSteppedLineBorderRadius() {
  Chart.defaults.elements.point.radius = 6; // Raggio dei punti
  Chart.defaults.elements.point.borderWidth = 2; // Bordo dei punti
  Chart.defaults.elements.point.borderColor = 'rgba(255, 99, 132, 1)'; // Colore del bordo
  Chart.defaults.elements.point.backgroundColor = 'rgba(255, 99, 132, 0.6)'; // Colore di sfondo
}

export function addStackedLineBorderRadius() {
  Chart.defaults.elements.point.radius = 6; // Raggio dei punti
  Chart.defaults.elements.point.borderWidth = 2; // Bordo dei punti
  Chart.defaults.elements.point.borderColor = 'rgba(255, 99, 132, 1)'; // Colore del bordo
  Chart.defaults.elements.point.backgroundColor = 'rgba(255, 99, 132, 0.6)'; // Colore di sfondo
}
