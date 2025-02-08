import { Injectable, signal } from '@angular/core';


export interface ActionMenuConfig {
  selectedElements: number;
  actions?: {
    id: string;
    label: string;
    icon: string;
    action: () => void;
  }[];
}

@Injectable({
  providedIn: 'root'
})
export class ActionMenuService {

  show = signal( false );

  config = signal<ActionMenuConfig | null>( null );

  // Overload della funzione per gestire i casi con e senza 'config'
  setVisibility(visibility: true, config: ActionMenuConfig): void; // Quando visibility è true, config è obbligatorio
  setVisibility(visibility: false): void; // Quando visibility è false, config non viene passato
  setVisibility( visibility: boolean, config?: ActionMenuConfig ) {
    this.show.set( visibility );
    this.config.set( config || null );
  }
}
