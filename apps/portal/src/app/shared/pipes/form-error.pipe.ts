import { Pipe, PipeTransform } from '@angular/core';
import { ValidationErrors } from "@angular/forms";
import { TranslateService } from '@ngx-translate/core';

@Pipe({
  name: 'formError',
  pure: true,
  standalone: true
})
export class FormErrorPipe implements PipeTransform {

  constructor(
    private translate: TranslateService
  ) {
  }

  transform(errors: ValidationErrors | null | undefined, status: any = null): string {
    if (!errors) return "";

    const errorMessages: { [key: string]: string | ((errors: any) => string) } = {
      'pattern':
        this.translate.instant('form-validator.pattern'),
      'matDatepickerParse':
        this.translate.instant('form-validator.matDatepickerParse'),
      'matDatetimePickerParse':
        this.translate.instant('form-validator.matDatetimePickerParse'),
      'maxlength':
        (errors: any) =>
          this.translate.instant('form-validator.maxlength', { value: errors?.['maxlength'].requiredLength}),
      'minlength':
        (errors: any) =>
          this.translate.instant('form-validator.minlength', { value: errors?.['minlength'].requiredLength}),
      'max':
        (errors: any) => {
        const max = (typeof errors?.['max'] === 'object') ? errors?.['max'].max : errors?.['max'];
        return this.translate.instant('form-validator.max' , { value: max });
      },
      'min':
        (errors: any) => {
        const min = (typeof errors?.['min'] === 'object') ? errors?.['min'].min : errors?.['min'];
        return this.translate.instant('form-validator.min', { value: min });
      },
      'matDatepickerMax':
        (errors: any) =>
          this.translate.instant('form-validator.matDatepickerMax', { value: errors?.['matDatepickerMax'].max.format("DD/MM/YYYY")}),
      'matDatepickerMin':
        (errors: any) =>
          this.translate.instant('form-validator.matDatepickerMin', { value: errors?.['matDatepickerMin'].min.format("DD/MM/YYYY")}),
      'minArrayLength':
        (errors: any) =>
          status === 'INVALID' ? this.translate.instant('form-validator.minArrayLength', { value: errors?.['minArrayLength']}) : '',
      'mutuallyInclusive':
        this.translate.instant('form-validator.mutuallyInclusive'),
      'required':
        this.translate.instant('form-validator.required')
    };

    for (const key in errorMessages) {
      if (key in errors) {
        const message = errorMessages[key];
        return typeof message === 'function' ? message(errors) : message;
      }
    }

    return this.translate.instant('required');
  }

}
