export interface PostQueryDto {
  id: number;
  author: number;
  createdAt: Date;
  updatedAt: Date;
  slug: string;
  content: string;
  likenum: number;
}
