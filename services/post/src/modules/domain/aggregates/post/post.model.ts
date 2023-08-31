// task.service.ts
import { Injectable } from '@nestjs/common';
import { UserUpdatePostInput } from '../user/user.model';

import crypto from 'crypto';

export type CreatePrePostInput = {
  author: bigint;
  images: string[];
  content: string;
};

export type CreatePostInput = {
  slug: string;
  author: bigint;
  images: string[];
  content: string;

  id: bigint;
  createdAt: Date;
  updatedAt: Date;
};

export type UpdatePostInput = {
  content: string;
};

@Injectable()
export class Post {
  private slug: string;
  private author: bigint;
  private images: string[];
  private content: string;

  private id?: bigint;
  private createdAt?: Date;
  private updatedAt?: Date;

  static create(p: CreatePrePostInput): Post {
    return new Post().setNewPost(p);
  }

  private setNewPost(p: CreatePrePostInput): Post {
    this.slug = Post.getSlug(p.content);
    this.author = p.author;
    this.images = p.images;
    this.content = p.content;
    return this;
  }

  isPre(): boolean {
    return !this.id;
  }

  static query(postInput: CreatePostInput): Post {
    return new Post().setQuery(postInput);
  }

  private setQuery(p: CreatePostInput): Post {
    this.slug = p.slug;
    this.author = p.author;
    this.images = p.images;
    this.content = p.content;
    this.id = p.id;
    this.createdAt = p.createdAt;
    this.updatedAt = p.updatedAt;
    return this;
  }

  getID(): bigint {
    return this.id;
  }

  getSlug(): string {
    return this.slug;
  }

  getAuthor(): bigint {
    return this.author;
  }

  getContent(): string {
    return this.content;
  }

  getImageSrcs(): string[] {
    return this.images;
  }

  queryPost(): any {
    return {
      id: this.id,
      slug: this.slug,
      author: this.author,
      images: [...this.images],
      content: this.content,
      createdAt: this.createdAt,
      updatedAt: this.updatedAt,
    };
  }

  update(changes: UserUpdatePostInput): Post {
    this.content = changes.content;
    return this;
  }

  getUpdates(): UpdatePostInput {
    return {
      content: this.content,
    };
  }

  delete(): Post {
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
