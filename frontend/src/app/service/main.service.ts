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
  checkFileType(singleFilename: string) {
    if (!singleFilename) {
      return 'unknown';
    }
    const fileExtension = (singleFilename.toLowerCase().split('.').pop() ?? '');

    switch(fileExtension) {
      case 'jpg':
      case 'jpeg':
      case 'png':
      case 'bmp':
      case 'gif':
        return 'photo';
      case 'mp4':
      case 'mkv':
      case 'avi':
      case 'flv':
      case 'wmv':
      case 'mov':
        return 'video';
      case 'pdf':
        return 'pdf';
      default:
        return 'unknown';
    }
  }
}
