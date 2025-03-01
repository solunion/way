import { HttpClient } from '@angular/common/http';
import { APP_INITIALIZER, Provider } from '@angular/core';
import { TranslateLoader, TranslateModule, TranslateService } from '@ngx-translate/core';
import { TranslateHttpLoader } from '@ngx-translate/http-loader';

export function createTranslateLoader(http: any) {
  return new TranslateHttpLoader(http, './i18n/', '.json');
}

export function appInitializerFactory(translate: TranslateService) {
  return () => {
    translate.setDefaultLang('en');
    return translate.use('en').toPromise();
  };
}

// Recupera i provider in modo sicuro
const translateProviders = TranslateModule.forRoot({
  loader: {
    provide: TranslateLoader,
    useFactory: createTranslateLoader,
    deps: [HttpClient] // ✅ Usa HttpClient invece di fetch
  }
}).providers as Provider[];

export const provideTranslation = [
  ...translateProviders,
  {
    provide: APP_INITIALIZER,
    useFactory: appInitializerFactory,
    deps: [TranslateService],
    multi: true
  }
];
