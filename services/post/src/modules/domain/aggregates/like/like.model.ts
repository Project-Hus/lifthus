import { Injectable } from '@nestjs/common';
import { User } from '../user/user.model';
import { Post } from '../post/post.model';

@Injectable()
export class Like<T> {
  constructor(
    readonly liker: User,
    readonly target: T,
    private liked: boolean,
  ) {}

  like(u: User): Like<T> | undefined {
    if (u.getID() !== this.liker.getID() || this.liked) return undefined;
    this.liked = true;
    return this;
  }

  unlike(u: User): Like<T> | undefined {
    if (u.getID() !== this.liker.getID() || !this.liked) return undefined;
    this.liked = false;
    return this;
  }
}
