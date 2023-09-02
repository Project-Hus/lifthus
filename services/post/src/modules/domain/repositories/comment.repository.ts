import { Injectable } from '@nestjs/common';
import { Comment } from '../aggregates/comment/comment.model';
import { stringifyAny } from 'src/common/utils/utils';
import { Post } from '../aggregates/post/post.model';

@Injectable()
export abstract class CommentRepository {
  private comments: Map<bigint, Comment> = new Map();
  private commentOrigins: Map<bigint, string> = new Map();

  private getCachekey(p: Comment): bigint {
    return BigInt(p.getID());
  }

  private getCacheString(p: Comment): string {
    return stringifyAny(p);
  }

  private isUpdated(p: Comment, origin: string): boolean {
    return this.getCacheString(p) !== origin;
  }

  async getCommentByID(cid: bigint): Promise<Comment | null> {
    const comment = await this._getCommentByID(cid);
    const cacheKey = this.getCachekey(comment);
    this.comments.set(cacheKey, comment);
    this.commentOrigins.set(cacheKey, this.getCacheString(comment));
    return comment;
  }

  async getComments(post: Post): Promise<Comment[]> {
    return await this._getComments(post);
  }

  async createComment(comment: Comment): Promise<Comment> {
    const newComment = await this._createComment(comment);
    const cacheKey = this.getCachekey(newComment);
    this.comments.set(cacheKey, newComment);
    return newComment;
  }

  async deleteComment(comment: Comment): Promise<Comment> {
    const cacheKey = this.getCachekey(comment);
    this.comments.delete(cacheKey);
    this.commentOrigins.delete(cacheKey);
    return await this._deleteComment(comment);
  }

  private clear() {
    this.comments.clear();
    this.commentOrigins.clear();
  }

  async save(): Promise<void> {
    const changes: Set<Comment> = new Set();
    this.comments.forEach((comment, key) => {
      const origin = this.commentOrigins.get(key);
      if (this.isUpdated(comment, origin)) changes.add(comment);
    });
    this.clear();
    return this._save(changes);
  }

  async cancel(): Promise<void> {
    this.clear();
    return this._save(new Set());
  }

  // Abstract methods to be implemented by the actual repository

  abstract _getCommentByID(cid: bigint): Promise<Comment | null>;
  abstract _getComments(post: Post): Promise<Comment[]>;

  abstract _createComment(comment: Comment): Promise<Comment>;
  abstract _deleteComment(traget: Comment): Promise<Comment>;

  abstract _save(changes: Set<Comment>): Promise<void>;
}
