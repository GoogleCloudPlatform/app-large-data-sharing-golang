import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class MainService {
  tags: string[] = [];
  tagsSubject = new Subject<string[]>();


  updateTags(newTags: string[]) {
    this.tags = [...newTags];
    this.tagsSubject.next(this.tags);
  }

  constructor() { }

  getTags() {
    return this.tags;
  }

  clearTag() {
    this.tags = [];
    this.tagsSubject.next(this.tags);
  }
  checkFileType(fileName: string) {
    if (!fileName) {
      return 'unknown';
    }
    const fileExtension = (fileName.toLowerCase().split('.').pop() ?? '');

    const photoExtensions = ['jpg', 'jpeg', 'png', 'bmp', 'gif'];
    const videoExtensions = ['mp4', 'mkv', 'avi', 'flv', 'wmv', 'mov'];
    const pdfExtension = 'pdf';

    if (photoExtensions.includes(fileExtension)) {
      return 'photo';
    } else if (videoExtensions.includes(fileExtension)) {
      return 'video';
    } else if (fileExtension === pdfExtension) {
      return 'pdf';
    } else {
      return 'unknown';
    }
  }
}
