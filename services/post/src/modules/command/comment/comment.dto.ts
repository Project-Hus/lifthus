export type CreateCommentDto = {
  postId?: number;
  parentId?: number;
  author: number;
  content: string;
};

export type UpdateCommentDto = {
  id: number;
  author: number;
  content: string;
};

export type CommentDto = {
  id: number;
  author: number;
  createdAt: Date;
  updatedAt: Date;
  postId: number;
  parentId: number | null;
  content: string;
  likenum: number;
};
