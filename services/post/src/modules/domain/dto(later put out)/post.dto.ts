export type CreatePrePostDto = {
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
