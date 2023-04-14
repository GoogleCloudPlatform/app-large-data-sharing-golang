import { Component, EventEmitter, Input, Output } from '@angular/core';
import { MatChipInputEvent } from '@angular/material/chips';
import { SPACE, ENTER } from '@angular/cdk/keycodes';
import { FormGroup } from '@angular/forms';
import { MainService } from '../service/main.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-search-bar',
  templateUrl: './search-bar.component.html',
  styleUrls: ['./search-bar.component.scss']
})
export class SearchBarComponent {
  @Output() clicked: EventEmitter<any> = new EventEmitter();
  @Input() formGroup: FormGroup | undefined;
  tags: string[] = [];
  private tagsSubscription: Subscription = new Subscription;
  readonly separatorKeysCodes: number[] = [SPACE, ENTER];
  constructor(private mainService: MainService) {

  }
  ngOnInit() {
    this.tags = this.mainService.tags;
    this.tagsSubscription = this.mainService.tagsSubject.subscribe(
      (tags: string[]) => {
        this.tags = tags;
      }
    );
  }
  searchTags() {
    this.clicked.emit({ eventName: "searchTags", tags: this.tags });
  }
  clickSearchTags() {
    const input = this.formGroup?.get('tags');
    const value = (input?.value || '').trim();
    const currentTags = [...this.tags];
    if (this.tags.length && value) {
      this.mainService.updateTags([...currentTags]);
    }
    input?.setValue('');
    this.clicked.emit({ eventName: "searchTags", tags: this.tags });
  }
  clearTag() {

    this.mainService.clearTag();
  }

  addTags(event: MatChipInputEvent) {
    const value = (event.value || '').trim();
    const newTags = this.tags.filter(tag => tag !== value);
    if (value) {
      newTags.push(value.toLowerCase());
      this.mainService.updateTags(newTags);
    }

    event.chipInput!.clear();
  }
  removeTag(tag: string) {
    const index = this.tags.indexOf(tag);
    const currentTags = [...this.tags];
    if (index >= 0) {
      currentTags.splice(index, 1);
      this.mainService.updateTags(currentTags);
    }
  }

}
