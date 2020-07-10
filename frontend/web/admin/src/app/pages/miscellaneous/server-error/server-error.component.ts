import { Component } from '@angular/core';
import { NbMenuService } from '@nebular/theme';

@Component({
  selector: 'server-error',
  styleUrls: ['./server-error.component.scss'],
  templateUrl: 'server-error.component.html',
})
export class ServerErrorComponent {

  constructor(private menuService: NbMenuService) {
  }

  goToHome() {
    this.menuService.navigateHome();
  }

}
