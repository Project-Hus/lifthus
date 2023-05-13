export interface PostQueryDto {
  id: PrimaryKey;
  author: PrimaryKey;
  createdAt: Date;
  updatedAt: Date;
  slug: string;
  content: string;
  likenum: number;
}
