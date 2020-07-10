import { NgModule, ModuleWithProviders } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Store } from './store';
import { LocalStorage } from './local-storage';
import { SessionStorage } from './session-storage';

@NgModule({
  imports: [
    CommonModule
  ],
  declarations: [],
  providers: [Store, LocalStorage, SessionStorage]
})
export class ReactiveStorageModule {
  static setLocalStorageKeys(keys?: string[]): ModuleWithProviders {
    return {
      ngModule: ReactiveStorageModule,
      providers: [{provide: 'localStoragekeys', useValue: keys}]
    };
  }
 }
