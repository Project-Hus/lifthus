import { Injectable } from '@nestjs/common';
import { User } from './user.model';

interface IPostLike {
  postId: bigint; // id with user
  user: User;
  createdAt: Date;
}

@Injectable()
export class PostLike {
  private postLike: IPostLike;
}
