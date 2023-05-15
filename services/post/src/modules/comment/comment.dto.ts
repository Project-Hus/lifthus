export type CreateCommentDto = {
  postId: PrimaryKey;
  parentId?: PrimaryKey;
  author: PrimaryKey;
  content: string;
};

export type UpdateCommentDto = {
  id: PrimaryKey;
  author: PrimaryKey;
  content: string;
};
