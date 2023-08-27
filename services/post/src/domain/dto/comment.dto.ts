export type CreateCommentDto = {
  postId: bigint; // id with user
  parentId?: bigint;
  author: bigint;
  content: string;
};

export type UpdateCommentDto = {
  id: bigint;
  content: string;
};
