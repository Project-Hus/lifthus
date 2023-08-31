// task.service.ts
import { Injectable } from '@nestjs/common';

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
  constructor(
    private author: bigint,
    private content: string,

    private postId: bigint,

    private parentId?: bigint,

    private id?: bigint,
    private createdAt?: Date,
    private updatedAt?: Date,

    private replies?: Comment[],
  ) {}

  static createComment(c: CreatePreCommentInput): Comment {
    return new Comment(c.author, c.content, c.postId);
  }

  static queryComment(p: CreateCommentInput): Comment {
    return new Comment(
      p.author,
      p.content,
      p.postId,
      undefined,
      p.id,
      p.createdAt,
      p.updatedAt,
      p.replies,
    );
  }

  static createReply(c: CreatePreReplyCommentInput): Comment {
    return new Comment(c.author, c.content, c.postId, c.parentId);
  }

  static queryReply(p: CreateReplyInput): Comment {
    return new Comment(
      p.author,
      p.content,
      p.postId,
      p.parentId,
      p.id,
      p.createdAt,
      p.updatedAt,
    );
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

  update(changes: UpdateCommentInput): Comment {
    this.content = changes.content;
    return this;
  }

  delete(): Comment {
    return this;
  }
}
