import { Injectable } from '@nestjs/common';
import {
  Comment,
  UpdateCommentInput,
} from '../../../domain/aggregates/comment/comment.model';

@Injectable()
export abstract class CommentRepository {
  async getComments(pid: bigint): Promise<Comment[]> {
    return await this._getComments(pid);
  }

  async getCommentsNum(pid: bigint): Promise<number> {
    return await this._getCommentsNum(pid);
  }

  async createComment(comment: Comment): Promise<Comment> {
    const newComment = await this._createComment(comment);
    return newComment;
  }

  async deleteComment(comment: Comment): Promise<Comment> {
    return await this._deleteComment(comment);
  }

  async save(comment: Comment): Promise<void> {
    return this._save(comment.getID(), comment.getUpdates());
  }

  // Abstract methods to be implemented by the actual repository

  abstract _getComments(pid: bigint): Promise<Comment[]>;
  abstract _getCommentsNum(pid: bigint): Promise<number>;

  abstract _createComment(comment: Comment): Promise<Comment>;
  abstract _deleteComment(traget: Comment): Promise<Comment>;

  abstract _save(cid: bigint, updates: UpdateCommentInput): Promise<void>;
}
