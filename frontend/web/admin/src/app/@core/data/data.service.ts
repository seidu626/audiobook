import { Injectable } from '@angular/core';
import { HttpClient} from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { map, catchError } from 'rxjs/operators';
import { LoaderService } from '../utils/loader.service';

// Implementing a Retry-Circuit breaker policy
// is pending to do for the SPA app
@Injectable()
export class DataService {
  constructor(private http: HttpClient, private loader: LoaderService) {}

  get(url: string, params?: any): Observable<any> {
    this.loader.show();
    const options = {};
    this.setHeaders(options);

    return this.http.get(url, options).pipe(
      // retry(3), // retry a failed request up to 3 times
      map((res: any) => {
        this.loader.hide();
        return res;
      }),
      catchError(this.handleError),
    );
  }

  postWithId(url: string, data: any, params?: any): Observable<any> {
    return this.doPost(url, data, true, params);
  }

  post(url: string, data: any, params?: any): Observable<any> {
    return this.doPost(url, data, false, params);
  }

  put(url: string, data: any, params?: any): Observable<any> {
    return this.doPut(url, data, true, params);
  }

  putWithId(url: string, data: any, params?: any): Observable<any> {
    return this.doPut(url, data, true, params);
  }

  private doPost(
    url: string,
    data: any,
    needId: boolean,
    params?: any,
  ): Observable<any> {
    this.loader.show();
    const options = {};
    this.setHeaders(options, needId);

    return this.http.post(url, data, options).pipe(
      map((res: any) => {
        this.loader.hide();
        return res;
      }),
      catchError(this.handleError),
    );
  }

  delete(url: string, params?: any): Observable<any> {
    this.loader.show();
    const options = {};
    this.setHeaders(options);
    return this.http.delete(url, options).pipe(
      // retry(3), // retry a failed request up to 3 times
      map((res: any) => {
        this.loader.hide();
        return res;
      }),
      catchError(this.handleError),
    );
  }

  private handleError(error: any) {
    if (error.error instanceof ErrorEvent) {
      // A client-side or network error occurred. Handle it accordingly.
      console.log(error);
      console.error('Client side network error occurred:', error.error.message);
    } else {
      // The backend returned an unsuccessful response code.
      // The response body may contain clues as to what went wrong,
      console.error(
        'Backend - ' +
          `status: ${error.status}, ` +
          `statusText: ${error.statusText}, ` +
          `message: ${error.error.message}`,
      );
    }

    // return an observable with a user-facing error message
    return throwError(error || 'server error');
  }

  private doPut(
    url: string,
    data: any,
    needId: boolean,
    params?: any,
  ): Observable<any> {
    this.loader.show();
    const options = {};
    this.setHeaders(options, needId);

    return this.http.put(url, data, options).pipe(
      map((res: any) => {
        this.loader.hide();
        return res;
      }),
      catchError(this.handleError),
    );
  }

  private setHeaders(options: any, needId?: boolean) {
    /*
    if (needId && this.securityService) {
      options['headers'] = new HttpHeaders()
        .append('authorization', 'Bearer ' + this.securityService.GetToken())
        .append('x-requestid', Guid.newGuid());
    } else if (this.securityService) {
      options['headers'] = new HttpHeaders().append(
        'authorization',
        'Bearer ' + this.securityService.GetToken()
      );
    }
    */
  }
}
