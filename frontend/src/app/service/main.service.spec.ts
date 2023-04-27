import { TestBed } from '@angular/core/testing';

import { MainService } from './main.service';

describe('MainService', () => {
  let service: MainService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(MainService);
  });

  it('should create the serivce', () => {
    expect(service).toBeTruthy();
  });

  it('should clear the tags', () => {
    service.updateTags(['test', 'another test'])
    service.clearTag();
    expect(service.tags).toEqual([]);
  });

  it('checks if the file is an image', () => {
    const type = service.checkFileType('test.jpg');
    expect(type).toEqual('photo');
  });

  it('checks if the file is a video', () => {
    const type = service.checkFileType('test.mp4');
    expect(type).toEqual('video');
  });

  it('checks if the file is a pdf document', () => {
    const type = service.checkFileType('test.pdf');
    expect(type).toEqual('pdf');
  });

  it('checks if the file type is unknown', () => {
    const type = service.checkFileType('test');
    expect(type).toEqual('unknown');
  });
});
