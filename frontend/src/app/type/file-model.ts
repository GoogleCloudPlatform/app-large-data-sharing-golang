export interface FileModel {
  id: string;
  createTime: string;
  name: string;
  path: string;
  tags: Array<string>;
  thumbUrl: string;
  url: string;
  updateTime: string;
  orderNo: string;
  size: number;
  resolution?: string
}