import { Component, EventEmitter, Output } from '@angular/core';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})

export class HeaderComponent {
  @Output() refreshToHome = new EventEmitter();
  backToHome() {
    this.refreshToHome.emit();
  }
}
