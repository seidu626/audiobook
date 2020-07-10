import { Injectable, Inject, Optional } from '@angular/core';
import { BehaviorSubject, Subject, Observable } from 'rxjs';
import { startWith, scan, map } from 'rxjs/operators';

@Injectable()
export class Store extends BehaviorSubject<any> {
  private dispatcher = new Subject();
  constructor( @Inject('localStoragekeys') @Optional() keys) {
    super({});
    const initialState = keys && keys.length > 0 ? this.getFromLocalStorage(keys) : {};

    this.dispatcher.pipe(
      startWith({}),
      scan((state, payload) => this.reducer(state, payload), initialState)
    ).subscribe(state => super.next(state));
  }


  getAll() {
    return this;
  }

  getItem<T>(key: string): Observable <T | null>  {
    if (!!key && typeof key === 'string') {
      return this.pipe(map(state => state[key]));
    }
  }

  setItem(key: string, payload) {
    if (!key) { return; }
    this.dispatcher.next({ key, payload });
  }

  removeItem(key: string) {
    if (!key) { return; }
    this.dispatcher.next({ key, payload: null });
  }

  private reducer(state: any, payload: any) {
    return {
      ...state,
      [payload.key]: payload.payload
    };
  }


  private getFromLocalStorage(keys: string[]) {
    const initialState = {};
    keys.forEach(key => {
      try {
        const payload = JSON.parse(localStorage.getItem(key));
        initialState[key] = payload;
      } catch (error) {

      }
    });
    return initialState;
  }

}

