// task.service.ts
import { Injectable } from '@nestjs/common';
import { User } from '../user/user.model';
import { Post } from '../post/post.model';

export type CreatePreCommentInput = {
  author: User;
  content: string;
  post: Post;
};

export type CreateCommentInput = {
  id: bigint;
  author: User;
  content: string;
  post: Post;
  parent?: Comment;
  replies: Comment[];
  createdAt: Date;
  updatedAt: Date;
};

export type UpdateCommentInput = {
  content: string;
};
@Injectable()
export class Comment {
  constructor(
    private author: User,
    private content: string,
    private post: Post,
    private parent?: Comment,

    private id?: bigint,
    private replies?: Comment[],
    private createdAt?: Date,
    private updatedAt?: Date,
  ) {}

  static createPre(p: CreatePreCommentInput): Comment {
    return new Comment(p.author, p.content, p.post);
  }

  isPre(p: Comment): boolean {
    return !p.id;
  }

  static create(p: CreateCommentInput): Comment {
    return new Comment(
      p.author,
      p.content,
      p.post,
      p.parent,
      p.id,
      p.replies,
      p.createdAt,
      p.updatedAt,
    );
  }

  getID(): bigint {
    return this.id;
  }

  getAuthor(): User {
    return this.author;
  }

  queryComment(): any {
    return {
      author: this.author,
      content: this.content,
      post: this.post,
      parent: this.parent,
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
}
