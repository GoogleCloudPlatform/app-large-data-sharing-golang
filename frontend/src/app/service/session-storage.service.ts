import { Injectable } from '@angular/core';

const IMG_DATA = 'imgData';

@Injectable({
  providedIn: 'root'
})
export class SessionStorageService {
  constructor() { }

  clearImageData(): void {
    window.sessionStorage.clear();
  }

  saveImageData(imgData: any): void {
    window.sessionStorage.removeItem(IMG_DATA);
    window.sessionStorage.setItem(IMG_DATA, JSON.stringify(imgData));
  }

  getImageData(): any {
    const imgae = window.sessionStorage.getItem(IMG_DATA);
    if (imgae) {
      return JSON.parse(imgae);
    }

    return {};
  }
}