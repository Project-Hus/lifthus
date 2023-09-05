import { Injectable } from '@nestjs/common';
import { Like } from '../../../domain/aggregates/like/like.model';
import { Post } from '../../../domain/aggregates/post/post.model';
import { User } from '../../../domain/aggregates/user/user.model';
import { stringifyAny } from 'src/common/utils/utils';

@Injectable()
export abstract class LikeRepository<T> {
  async getLike(u: User, t: T): Promise<Like<T>> {
    const like = await this._getLike(u, t);
    return like;
  }

  async getLikesNum(tid: bigint): Promise<number> {
    return this._getLikesNum(tid);
  }

  async save(like: Like<T>): Promise<void> {
    this._save(like);
  }

  // Abstract methods to be implemented by the actual repository

  abstract _getLike(u: User, t: T): Promise<Like<T>>;

  abstract _getLikesNum(tid: bigint): Promise<number>;

  abstract _save(like: Like<T>): Promise<void>;
}

@Injectable()
export abstract class PostLikeRepository extends LikeRepository<Post> {}

@Injectable()
export abstract class CommentLikeRepository extends LikeRepository<Comment> {}
