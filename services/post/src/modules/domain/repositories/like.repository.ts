import { Injectable } from '@nestjs/common';
import { Like } from '../aggregates/like/like.model';
import { Post } from '../aggregates/post/post.model';
import { User } from '../aggregates/user/user.model';

@Injectable()
export abstract class LikeRepository<T> {
  private likes: Set<Like<T>> = new Set();

  constructor() {}

  async getLike(u: User, t: T): Promise<Like<T>> {
    const like = await this._getLike(u, t);
    this.likes.add(like);
    return like;
  }

  async getLikeNum(t: T): Promise<number> {
    return this._getLikeNum(t);
  }

  async save(): Promise<void> {
    this._save(this.likes);
    this.clear();
  }

  async cancel(): Promise<void> {
    this.clear();
    return;
  }

  private clear() {
    this.likes.clear();
  }

  abstract _getLike(u: User, t: T): Promise<Like<T>>;

  abstract _getLikeNum(t: T): Promise<number>;

  abstract _save(likes: Set<Like<T>>): Promise<void>;
}
