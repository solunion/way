import { Component, input } from '@angular/core';
import { MatTableModule } from '@angular/material/table';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { animate, state, style, transition, trigger } from '@angular/animations';
import { MatPaginator } from '@angular/material/paginator';

@Component({
  selector: 'way-table',
  imports: [MatTableModule, MatButtonModule, MatIconModule, MatPaginator],
  templateUrl: './table.component.html',
  styleUrl: './table.component.scss',
  standalone: true
})
export class TableComponent{

  /**
   * @description the columns to display in the table.
   */
  dataSource = input.required<any>()
  displayedColumns: string[] = ['position', 'name', 'weight', 'symbol'];
}


export interface PeriodicElement {
  name: string;
  position: number;
  weight: number;
  symbol: string;
  description: string;
}
