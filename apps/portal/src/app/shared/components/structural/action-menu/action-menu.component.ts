import { Component } from '@angular/core';
import { MatIconButton } from '@angular/material/button';
import { ActionMenuService } from './action-menu.service';
import { JsonPipe } from '@angular/common';
import { MatTooltip } from '@angular/material/tooltip';

@Component({
  selector: 'way-action-menu',
  imports: [
    MatIconButton,
    JsonPipe,
    MatTooltip
  ],
  standalone: true,
  templateUrl: './action-menu.component.html',
})
export class ActionMenuComponent {

  constructor(
    public actionMenuService: ActionMenuService
  ) {

  }

}
