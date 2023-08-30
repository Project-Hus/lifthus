import { Injectable } from '@nestjs/common';
import { Like } from '../aggregates/like/like.model';
import { Post } from '../aggregates/post/post.model';
import { User } from '../aggregates/user/user.model';
import { stringifyAny } from 'src/common/utils/utils';

@Injectable()
export abstract class LikeRepository<T> {
  private likes: Map<string, Like<T>> = new Map();
  private likeOrigins: Map<string, string> = new Map();

  constructor() {}

  private getCachekey(u: User, t: T): string {
    return stringifyAny({ u, t });
  }
  private getCacheString(like: Like<T>): string {
    return stringifyAny(like);
  }

  async getLike(u: User, t: T): Promise<Like<T>> {
    const like = await this._getLike(u, t);
    const cacheKey = this.getCachekey(u, t);
    this.likes.set(cacheKey, like);
    this.likeOrigins.set(cacheKey, this.getCacheString(like));
    return like;
  }

  async getLikeNum(t: T): Promise<number> {
    return this._getLikeNum(t);
  }

  async save(): Promise<void> {
    // filter out what has changed
    const changes: Set<Like<T>> = new Set();
    this.likes.forEach((like, key) => {
      const origin = this.likeOrigins.get(key);
      if (origin !== this.getCacheString(like)) [changes.add(like)];
    });
    // pass the chages
    this._save(changes);
    this.clear();
  }

  async cancel(): Promise<void> {
    this.clear();
    return this._save(new Set());
  }

  private clear() {
    this.likes.clear();
    this.likeOrigins.clear();
  }

  abstract _getLike(u: User, t: T): Promise<Like<T>>;

  abstract _getLikeNum(t: T): Promise<number>;

  abstract _save(likes: Set<Like<T>>): Promise<void>;
}
