export type CreateWaitingCommentDto = {
  author: bigint;
};

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

export type CommentLikeDto = {
  userId: bigint;
  commentId: bigint;
};

export type CommentUnlikeDto = {
  userId: bigint;
  commentId: bigint;
};
