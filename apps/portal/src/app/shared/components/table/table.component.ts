import { Component, input, output } from '@angular/core';
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
   * The data source of the table, which is provided as an input.
   */
  dataSource = input.required<any>();

  /**
   * The columns to display in the table, which is provided as an input.
   */
  // columns = input.required<any>();

  allowImport = input<boolean>(false);

  allowExport = input<boolean>(false);

  allowEditColumns = input<boolean>(false);

  displayedColumns: string[] = ['position', 'name', 'weight', 'symbol'];
}


export interface PeriodicElement {
  name: string;
  position: number;
  weight: number;
  symbol: string;
  description: string;
}
