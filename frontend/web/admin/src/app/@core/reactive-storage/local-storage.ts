import { Injectable } from '@angular/core';
import { Store } from './store';

@Injectable()
export class LocalStorage {
  constructor(private store: Store) {}

  setItem<T>(key: string, payload: T) {
    this.syncWithLocalStorage(key, payload);
    this.store.setItem(key, payload);
  }

  getItem<T>(key) {
    return this.store.getItem(key);
  }

  removeItem(key) {
    this.syncWithLocalStorage(key, null);
  }

  private syncWithLocalStorage(key: string, payload: any) {
    try {
      if (!!payload) {
        localStorage.setItem(key, JSON.stringify(payload));
      } else  {
        localStorage.removeItem(key);
      }
    } catch (error) {

    }
  }
}
