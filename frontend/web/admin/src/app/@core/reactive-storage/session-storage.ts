import { Observable } from 'rxjs';
import { Injectable } from '@angular/core';
import { Store } from './store';

@Injectable()
export class SessionStorage {
  constructor(private store: Store) {}

  setItem<T>(key: string, payload: T) {
    this.store.setItem(key, payload);
  }

  getItem<T>(key): Observable<T> {
    return this.store.getItem<T>(key);
  }

  removeItem(key) {
    this.store.removeItem(key);
  }

}
