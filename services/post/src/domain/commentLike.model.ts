import { Injectable } from '@nestjs/common';
import { User } from './user.model';

interface ICommentLike {
  commentId: bigint; // id with user
  user: User;
  createdAt: Date;
}

@Injectable()
export class CommentLike {
  private commentLike: ICommentLike;
}
