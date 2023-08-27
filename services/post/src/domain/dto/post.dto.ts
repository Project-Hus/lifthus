import { Comment } from '../comment.model';
import { Post } from '../post.model';
import { User } from '../user.model';

export type CreatePostDto = {
  author: User;
  images: string[];
  content: string;
};

export type QueryPostDto = {
  id: bigint;
  slug: string;

  author: User;

  images: string[];
  content: string;

  likenum: number;
  likers: bigint[];

  comments?: Comment[];

  createdAt: Date;
  updatedAt: Date;
};

export type UpdatePostDto = {
  id: bigint;
  content: string;
};

export type PostLikeDto = {
  user: User;
  post: Post;
};

export type PostUnlikeDto = {
  user: User;
  post: Post;
};
