import { Injectable } from '@angular/core';
import { FileModel } from '../type/file-model';

const IMG_DATA = 'imgData';

@Injectable({
  providedIn: 'root'
})
export class SessionStorageService {
  constructor() { }

  clearImageData(): void {
    window.sessionStorage.clear();
  }

  saveImageData(imgData: FileModel): void {
    window.sessionStorage.removeItem(IMG_DATA);
    window.sessionStorage.setItem(IMG_DATA, JSON.stringify(imgData));
  }

  getImageData(): FileModel {
    const image = window.sessionStorage.getItem(IMG_DATA);
    if (image) {
      return JSON.parse(image);
    }

    return {
      id: '',
      createTime: '',
      name: '',
      path: '',
      tags: [],
      thumbUrl: '',
      url: '',
      updateTime: '',
      orderNo: '',
      size: 0
    };
  }
}