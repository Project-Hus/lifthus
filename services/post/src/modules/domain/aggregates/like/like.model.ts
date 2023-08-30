import { Injectable } from '@nestjs/common';
import { User } from '../user/user.model';
import { Post } from '../post/post.model';

@Injectable()
export class Like {
  constructor(
    readonly user: User,
    private liked: boolean,
    readonly post?: Post,
    readonly comment?: Comment,
  ) {}

  like(l: Like): Like | undefined {
    if (l.liked) return;
    l.liked = true;
    return l;
  }

  unlike(l: Like): Like | undefined {
    if (!l.liked) return;
    l.liked = false;
    return l;
  }
}
