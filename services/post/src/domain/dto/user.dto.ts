import { Comment } from '../comment.model';
import { Post } from '../post.model';
import { User } from '../user.model';

export type CreateUserDto = {
  id: bigint;
  posts?: Post[];
  comments?: Comment[];
  postLikes?: bigint[];
  commentLikes?: bigint[];
};
