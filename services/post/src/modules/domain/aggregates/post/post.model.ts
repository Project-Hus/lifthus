// task.service.ts
import { Injectable } from '@nestjs/common';
import { User } from '../user/user.model';
import { UpdatePostDto } from '../../dto(later put out)/post.dto';

import crypto from 'crypto';

export type CreatePrePostInput = {
  author: User;
  images: string[];
  content: string;
};

export type CreatePostInput = {
  slug: string;
  author: User;
  images: string[];
  content: string;

  id: bigint;
  createdAt: Date;
  updatedAt: Date;
};

export type UpdatePostInput = {
  id: bigint;
  content: string;
};

@Injectable()
export class Post {
  private constructor(
    private slug: string,
    private author: User,
    private images: string[],
    private content: string,

    private id?: bigint,
    private createdAt?: Date,
    private updatedAt?: Date,

    private querier?: User,
    private querierLiked?: boolean,
  ) {}

  static createPre(p: CreatePrePostInput): Post {
    return new Post(Post.getSlug(p.content), p.author, p.images, p.content);
  }

  static create(postInput: CreatePostInput): Post {
    return new Post(
      postInput.slug,
      postInput.author,
      postInput.images,
      postInput.content,
      postInput.id,
      postInput.createdAt,
      postInput.updatedAt,
    );
  }

  isPre(post: Post): boolean {
    return !post.id;
  }

  getID(): bigint {
    return this.id;
  }

  getSlug(): string {
    return this.slug;
  }

  getAuthor(): User {
    return this.author;
  }

  queryPost(): any {
    return {
      id: this.id,
      slug: this.slug,
      author: this.author.getID(),
      images: [...this.images],
      content: this.content,
      createdAt: this.createdAt,
      updatedAt: this.updatedAt,
    };
  }

  update(changes: UpdatePostDto): Post {
    this.content = changes.content;
    return this;
  }

  private static getSlug(content: string): string {
    let slug: string;
    const slugEnd: number = content.indexOf('\n');
    if (slugEnd == -1 || slugEnd > 30) {
      slug = content.slice(0, 30);
    } else {
      slug = content.slice(0, slugEnd);
    }
    slug = encodeURIComponent(slug + crypto.randomBytes(8).toString('hex'));
    return slug;
  }
}
