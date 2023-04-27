import { Component, EventEmitter, Output } from '@angular/core';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
})

export class HeaderComponent {
  @Output() refreshToHome = new EventEmitter();
  backToHome() {
    this.refreshToHome.emit();
  }
}
