// task.service.ts
import { BadRequestException, ForbiddenException } from '@nestjs/common';
import {
  CommentParents,
  CommentUpdates,
} from 'src/domain/aggregates/comment/comment.vo';
import { Timestamps } from 'src/domain/vo';
import { COMMENT_MAX_CONTENT_LENGTH } from 'src/shared/constraints';

export type CreateCommentInput = {
  id: bigint;
  author: bigint;
  content: string;
  postId: bigint;
  replies: Comment[];
  createdAt: Date;
  updatedAt: Date;
};

export type CreateReplyInput = {
  id: bigint;
  author: bigint;
  content: string;
  postId: bigint;
  parentId: bigint;
  createdAt: Date;
  updatedAt: Date;
};

export type UpdateCommentInput = {
  content: string;
};

export class Comment {
  private id?: bigint;
  private createdAt?: Date;
  private updatedAt?: Date;

  private author: bigint;
  private postId: bigint;
  private parentId?: bigint;

  private content: string;
  private replies?: Comment[];

  private static readonly MAX_CONTENT_LENGTH = COMMENT_MAX_CONTENT_LENGTH;

  constructor() {}

  static create(author: bigint, parents: CommentParents, content: string) {
    return new Comment().create(author, parents, content);
  }

  private create(
    author: bigint,
    parents: CommentParents,
    content: string,
  ): Comment {
    if (content.length > Comment.MAX_CONTENT_LENGTH) throw BadRequestException;
    this.author = author;
    this.postId = parents.postId;
    this.parentId = parents.parentId;
    this.content = content;
    return this;
  }

  isPersisted(): boolean {
    return !!this.id;
  }

  isReply(): boolean {
    return !!this.parentId;
  }

  static from(
    author: bigint,
    id: bigint,
    parents: CommentParents,
    content: string,
    timestamps: Timestamps,
    replies?: Comment[],
  ): Comment {
    return new Comment().from(
      author,
      id,
      parents,
      content,
      timestamps,
      replies,
    );
  }

  private from(
    author: bigint,
    id: bigint,
    parents: CommentParents,
    content: string,
    timestamps: Timestamps,
    replies?: Comment[],
  ) {
    this.id = id;
    this.author = author;
    this.postId = parents.postId;
    this.parentId = parents.parentId;
    this.content = content;
    this.createdAt = timestamps.createdAt;
    this.updatedAt = timestamps.updatedAt;
    this.replies = replies;
    return this;
  }

  getID(): bigint {
    return this.id;
  }

  getAuthor(): bigint {
    return this.author;
  }

  getPostID(): bigint {
    return this.postId;
  }

  getParentID(): bigint {
    return this.parentId;
  }

  getCreatedAt(): Date {
    return this.createdAt;
  }

  getUpdatedAt(): Date {
    return this.updatedAt;
  }

  getContent(): string {
    return this.content;
  }

  getReplies(): Comment[] {
    return this.replies;
  }

  update(changes: CommentUpdates): Comment {
    if (changes.content.length > Comment.MAX_CONTENT_LENGTH)
      throw BadRequestException;
    this.content = changes.content;
    return this;
  }

  getUpdates(): CommentUpdates {
    return {
      content: this.content,
    };
  }

  delete(deleter: bigint): Comment {
    if (this.author !== deleter) throw ForbiddenException;
    return this;
  }
}
