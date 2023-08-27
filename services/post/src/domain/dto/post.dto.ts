export type CreateWaitingPostDto = {
  author: bigint;
  srcs: string[];
  content: string;
};

export type CreatePostDto = {
  author: bigint;
  srcs: string[];
  content: string;
  slug: string;
  likenum: number;
};

export type UpdatePostDto = {
  id: bigint;
  content: string;
};

export type PostLikeDto = {
  userId: bigint;
  postId: bigint;
};

export type PostUnlikeDto = {
  userId: bigint;
  postId: bigint;
};
