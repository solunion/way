import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { ActionMenuConfig, ActionMenuService } from '../components/structural/action-menu/action-menu.service';

@Injectable({
  providedIn: 'root'
})
export class UtilityService {

  /**
   * @description This is a private BehaviorSubject that will be used to store the current state of the dark mode.
   * @private
   */
  #BSDdarkModeOn = new BehaviorSubject<boolean>(false);

  /**
   * @description This is a public Observable that will be used to subscribe to the dark mode state.
   * @public
   */
  darkModeOn$ = this.#BSDdarkModeOn.asObservable();

  constructor(
    private actionMenuService: ActionMenuService
  ) {
  }

  /**
   * @description This method will be used to toggle the dark mode state.
   * @param state {boolean} - The state of the dark mode.
   */
  public setDarkMode( state : boolean ) {
    this.#BSDdarkModeOn.next( state );
  }

  // region ACTION MENU
    showActionMenu( config: ActionMenuConfig ) {
      this.actionMenuService.setVisibility(true, config );
    }

    hideActionMenu() {
      this.actionMenuService.setVisibility(false);
    }
  // endregion

}
