import { Component, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-confirm-dialog',
  templateUrl: './confirm-dialog.component.html',
})
export class ConfirmDialogComponent {
  @Input() deleteId: string = '';
  @Output() confirmDelete = new EventEmitter();
  @Output() confirmCancel = new EventEmitter();
  cancel(){
    this.confirmCancel.emit();
  }
  delete(event: any){
    event.target.disabled = true;
    this.confirmDelete.emit();
  }
}
