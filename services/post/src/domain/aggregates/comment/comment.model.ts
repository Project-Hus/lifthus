// task.service.ts
import { Injectable } from '@nestjs/common';
import { CreateCommentServiceDto } from 'src/dto/inbound/comment.dto';

export type CreatePreCommentInput = {
  author: bigint;
  postId: bigint;
  content: string;
};

export type CreateCommentInput = {
  id: bigint;
  author: bigint;
  content: string;
  postId: bigint;
  replies: Comment[];
  createdAt: Date;
  updatedAt: Date;
};

export type CreatePreReplyCommentInput = {
  author: bigint;
  postId: bigint;
  parentId: bigint;
  content: string;
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

@Injectable()
export class Comment {
  private author: bigint;
  private content: string;

  private postId: bigint;

  private parentId?: bigint;

  private id?: bigint;
  private createdAt?: Date;
  private updatedAt?: Date;

  private replies?: Comment[];

  static createComment(c: CreateCommentServiceDto): Comment {
    return new Comment().setNewComment(c);
  }

  private setNewComment(c: CreateCommentServiceDto): Comment {
    this.author = c.author;
    this.content = c.content;
    this.postId = c.postId;
    this.parentId = c.parentId;
    return this;
  }

  static queryComment(p: CreateCommentInput): Comment {
    return new Comment().setQueryComment(p);
  }

  private setQueryComment(p: CreateCommentInput): Comment {
    this.author = p.author;
    this.content = p.content;
    this.postId = p.postId;
    this.id = p.id;
    this.createdAt = p.createdAt;
    this.updatedAt = p.updatedAt;
    this.replies = p.replies;
    return this;
  }

  static createReply(c: CreatePreReplyCommentInput): Comment {
    return new Comment().setNewReply(c);
  }

  private setNewReply(c: CreatePreReplyCommentInput): Comment {
    this.author = c.author;
    this.content = c.content;
    this.postId = c.postId;
    this.parentId = c.parentId;
    return this;
  }

  static queryReply(p: CreateReplyInput): Comment {
    return new Comment().setQueryReply(p);
  }

  private setQueryReply(p: CreateReplyInput): Comment {
    this.author = p.author;
    this.content = p.content;
    this.postId = p.postId;
    this.parentId = p.parentId;
    this.id = p.id;
    this.createdAt = p.createdAt;
    this.updatedAt = p.updatedAt;
    return this;
  }

  isPre(): boolean {
    return !this.id;
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

  queryComment(): any {
    return {
      author: this.author,
      content: this.content,
      postId: this.postId,
      parentId: this.parentId,
      id: this.id,
      replies: this.replies,
      createdAt: this.createdAt,
      updatedAt: this.updatedAt,
    };
  }

  getReplies(): Comment[] {
    return this.replies;
  }

  update(changes: UpdateCommentInput): Comment {
    this.content = changes.content;
    return this;
  }

  getUpdates(): UpdateCommentInput {
    return {
      content: this.content,
    };
  }

  delete(): Comment {
    return this;
  }
}
