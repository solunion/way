import { Component, computed, effect, EventEmitter, input, OnInit, Output, output } from '@angular/core';
import { MatTableModule } from '@angular/material/table';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatPaginator } from '@angular/material/paginator';
import { TableColumns, TableColumnType } from './table.models';
import { DatePipe, JsonPipe } from '@angular/common';
import { MatCheckbox } from '@angular/material/checkbox';
import {SelectionModel} from '@angular/cdk/collections';

@Component({
  selector: 'way-table',
  imports: [MatTableModule, MatButtonModule, MatIconModule, MatPaginator, DatePipe, JsonPipe, MatCheckbox],
  templateUrl: './table.component.html',
  styleUrl: './table.component.scss',
  standalone: true
})
export class TableComponent implements OnInit {

  /**
   * The data source of the table, which is provided as an input.
   */
  dataSource = input.required<any>();

  /**
   * The data source of the table, which is provided as an input.
   */
  columns = input.required<TableColumns[]>();

  /**
   * Computes the displayed columns for the table.
   * If row selection is allowed, adds a 'select' column at the beginning.
   * @returns {string[]} The list of column IDs to be displayed.
   */
  displayedColumns = computed(() => {
    const cols = this.columns().map(column => column.id);

    if( this.allowSelection() ) {
      cols.unshift('select');
    }

    return cols;
  });

  /**
   * Whether importing is allowed, provided as an input.
   * @type {boolean}
   */
  allowImport = input<boolean>(false);

  /**
   * Whether exporting is allowed, provided as an input.
   * @type {boolean}
   */
  allowExport = input<boolean>(false);

  /**
   * Whether row selection is allowed, provided as an input.
   * @type {boolean}
   */
  allowSelection = input<boolean>(false);

  /**
   * Whether editing columns is allowed, provided as an input.
   * @type {boolean}
   */
  allowEditColumns = input<boolean>(false);

  /**
   * The type of table columns, which is a read-only property.
   */
  protected readonly TableColumnType = TableColumnType;

  /**
   * The selection model for the table, which manages the selection state.
   * @type {SelectionModel<PeriodicElement>}
   */
  selection = new SelectionModel<PeriodicElement>(true, []);

  @Output() elementSelectedChange = new EventEmitter<PeriodicElement[]>();

  /**
   * @constructor
   */
  constructor() {
  }


  ngOnInit() {

  }


  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.dataSource().length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  toggleAllRows() {
    if (this.isAllSelected()) {
      this.selection.clear();
      this.elementSelectedChange.emit(this.selection.selected || []);
      return;
    }

    this.elementSelectedChange.emit([...this.dataSource()]);
    this.selection.select(...this.dataSource());
  }

  /** The label for the checkbox on the passed row */
  checkboxLabel(row?: PeriodicElement): string {
    if (!row) {
      return `${this.isAllSelected() ? 'deselect' : 'select'} all`;
    }
    return `${this.selection.isSelected(row) ? 'deselect' : 'select'} row ${row.position + 1}`;
  }

  onToggleRow(row: PeriodicElement) {
    this.selection.toggle(row);
    this.elementSelectedChange.emit(this.selection.selected || []);
  }
}


export interface PeriodicElement {
  name: string;
  position: number;
  weight: number;
  symbol: string;
  description: string;
}
