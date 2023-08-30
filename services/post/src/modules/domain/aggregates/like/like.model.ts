import { Injectable } from '@nestjs/common';
import { User } from '../user/user.model';
import { Post } from '../post/post.model';

@Injectable()
export class Like {
  constructor(
    private user: User,
    private liked: boolean,
    private post?: Post,
    private comment?: Comment,
  ) {}

  like(l: Like): Like {
    l.liked = true;
    return l;
  }

  unlike(l: Like): Like {
    l.liked = false;
    return l;
  }
}
