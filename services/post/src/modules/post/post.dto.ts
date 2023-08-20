export type CreatePostDto = {
  userGroup?: number;
  author: number;
  content: string;
  images?: Express.Multer.File[];
};

export type UpdatePostDto = {
  id: number;
  author: number;
  content: string;
};

export type PostDto = {
  id: number;
  userGroup: number;
  author: number;
  createdAt: Date;
  updatedAt: Date;
  slug: string;
  content: string;
  likenum: number;
};
