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
  private isUpdated(like: Like<T>, origin: string): boolean {
    return this.getCacheString(like) !== origin;
  }

  async getLike(u: User, t: T): Promise<Like<T>> {
    const like = await this._getLike(u, t);
    const cacheKey = this.getCachekey(u, t);
    this.likes.set(cacheKey, like);
    this.likeOrigins.set(cacheKey, this.getCacheString(like));
    return like;
  }

  async getLikesNum(tid: bigint): Promise<number> {
    return this._getLikesNum(tid);
  }

  async save(): Promise<void> {
    // filter out what has changed
    const changes: Set<Like<T>> = new Set();
    this.likes.forEach((like, key) => {
      const origin = this.likeOrigins.get(key);
      if (this.isUpdated(like, origin)) changes.add(like);
    });
    // pass the chages
    this.clear();
    this._save(changes);
  }

  async cancel(): Promise<void> {
    this.clear();
    return this._save(new Set());
  }

  private clear() {
    this.likes.clear();
    this.likeOrigins.clear();
  }

  // Abstract methods to be implemented by the actual repository

  abstract _getLike(u: User, t: T): Promise<Like<T>>;

  abstract _getLikesNum(tid: bigint): Promise<number>;

  abstract _save(likes: Set<Like<T>>): Promise<void>;
}

@Injectable()
export abstract class PostLikeRepository extends LikeRepository<Post> {}

@Injectable()
export abstract class CommentLikeRepository extends LikeRepository<Comment> {}
