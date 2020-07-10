import { Injectable } from '@angular/core';
import { NgProgress, NgProgressRef } from '@ngx-progressbar/core';

@Injectable()

export class LoaderService {

  progressRef: NgProgressRef;

  // https://www.npmjs.com/package/ngx-progressbar#ngprogressservice-options-functions
    constructor(public progress: NgProgress) {
      this.progressRef = this.progress.ref();
    }

    show() {
      /** request started */
      this.progressRef.start();
       // this.loaderSubject.next(<LoaderState>{show: true});
    }

    // NgProgress.set(n) Sets a percentage n (where n is between 0- 1)

    // NgProgress.inc(n) Increments by n (where n is between 0- 1)

    hide() {
      /** request completed */
      this.progressRef.complete();
        // this.loaderSubject.next(<LoaderState>{show: false});
    }
}
