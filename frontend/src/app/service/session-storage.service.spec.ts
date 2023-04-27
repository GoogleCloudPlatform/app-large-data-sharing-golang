import { TestBed } from '@angular/core/testing';

import { SessionStorageService } from './session-storage.service';

describe('SessionStorageService', () => {
  let service: SessionStorageService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SessionStorageService);
  });

  it('should create the serivce', () => {
    expect(service).toBeTruthy();
  });

  const imgData = {
    id: 'id',
    createTime: 'createTime',
    name: 'name',
    path: 'path',
    tags: ['tag1', 'tag2'],
    thumbUrl: 'thumbUrl',
    url: 'url',
    updateTime: 'updateTime',
    orderNo: 'orderNo',
    size: 100
  };
  it('should save the image data correctly', () => {
    service.saveImageData(imgData);
    const imageData = service.getImageData();
    expect(imageData).toEqual(imgData);
  });

  it('should retrieve image data correctly', () => {
    const imageData = service.getImageData();  
    expect(imageData).toEqual(imgData);
  });
});
