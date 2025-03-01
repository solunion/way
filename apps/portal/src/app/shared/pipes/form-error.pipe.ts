import { Pipe, PipeTransform } from '@angular/core';
import { ValidationErrors } from "@angular/forms";

@Pipe({
  name: 'formError',
  pure: true,
  standalone: true
})
export class FormErrorPipe implements PipeTransform {

  transform(errors: ValidationErrors | null | undefined, status: any = null): string {
    if (!errors) return "";

    const errorMessages: { [key: string]: string | ((errors: any) => string) } = {
      'pattern': "Il valore inserito non è in un formato valido",
      'matDatepickerParse': "Inserire una data nel formato gg/mm/aaaa",
      'matDatetimePickerParse': "Inserire il dato nel formato gg/mm/aaaa hh:mm",
      'maxlength': (errors: any) => `Inserire massimo ${errors?.['maxlength'].requiredLength} caratteri`,
      'minlength': (errors: any) => `Inserire minimo ${errors?.['minlength'].requiredLength} caratteri`,
      'max': (errors: any) => {
        const max = (typeof errors?.['max'] === 'object') ? errors?.['max'].max : errors?.['max'];
        return 'Inserire un valore minore o uguale a ' + `${max}`;
      },
      'min': (errors: any) => {
        const min = (typeof errors?.['min'] === 'object') ? errors?.['min'].min : errors?.['min'];
        return 'Inserire un valore maggiore o uguale a ' + `${min}`;
      },
      'matDatepickerMax': (errors: any) => `Inserire una data precedente al ${errors?.['matDatepickerMax'].max.format("DD/MM/YYYY")}`,
      'matDatepickerMin': (errors: any) => `Inserire una data posteriore al ${errors?.['matDatepickerMin'].min.format("DD/MM/YYYY")}`,
      'minArrayLength': (errors: any) => status === 'INVALID' ? "Inserire almeno " + errors?.['minArrayLength'] + " elementi" : '',
      'mutuallyInclusive': "Compilare tutti i campi",
      'required': "Questo campo è obbligatorio"
    };



    for (const key in errorMessages) {
      if (key in errors) {
        const message = errorMessages[key];
        return typeof message === 'function' ? message(errors) : message;
      }
    }

    debugger;
    return "Errore sconosciuto";
  }

}
