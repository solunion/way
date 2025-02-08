import { Component, computed, input, output } from '@angular/core';
import { MatTableModule } from '@angular/material/table';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { animate, state, style, transition, trigger } from '@angular/animations';
import { MatPaginator } from '@angular/material/paginator';
import { TableColumns, TableColumnType } from './table.models';
import { DatePipe, JsonPipe } from '@angular/common';

@Component({
  selector: 'way-table',
  imports: [MatTableModule, MatButtonModule, MatIconModule, MatPaginator, DatePipe, JsonPipe],
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
   * The data source of the table, which is provided as an input.
   */
  columns = input.required<TableColumns[]>();

  displayedColumns = computed(() => this.columns().map(column => column.id));

  /**
   * The columns to display in the table, which is provided as an input.
   */
  // columns = input.required<any>();

  allowImport = input<boolean>(false);

  allowExport = input<boolean>(false);

  allowEditColumns = input<boolean>(false);
  protected readonly TableColumnType = TableColumnType;
}


export interface PeriodicElement {
  name: string;
  position: number;
  weight: number;
  symbol: string;
  description: string;
}
