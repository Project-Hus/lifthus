// task.service.ts
import { Injectable } from '@nestjs/common';
import { UserUpdatePostInput } from '../user/user.model';

import crypto from 'crypto';
import { SLUG_MAX_LENGTH } from 'src/common/constraints';
import { CreatePostServiceDto } from 'src/dto/inbound/post.dto';

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
  private imageSrcs: string[];
  private content: string;

  private id?: bigint;
  private createdAt?: Date;
  private updatedAt?: Date;

  static create(p: CreatePostServiceDto): Post {
    return new Post().setNewPost(p);
  }

  private setNewPost(p: CreatePostServiceDto): Post {
    this.slug = Post.generateSlug(p.content);
    this.author = p.author;
    this.imageSrcs = p.imageSrcs;
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
    this.imageSrcs = p.images;
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

  getCreatedAt(): Date {
    return this.createdAt;
  }

  getUpdatedAt(): Date {
    return this.updatedAt;
  }

  getImageSrcs(): string[] {
    return this.imageSrcs;
  }

  queryPost(): any {
    return {
      id: this.id,
      slug: this.slug,
      author: this.author,
      images: [...this.imageSrcs],
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

  private static generateSlug(content: string): string {
    let slug: string;
    const slugEnd: number = content.indexOf('\n');
    if (slugEnd == -1 || slugEnd > SLUG_MAX_LENGTH) {
      slug = content.slice(0, SLUG_MAX_LENGTH);
    } else {
      slug = content.slice(0, slugEnd);
    }
    slug = slug + 'code' + crypto.randomBytes(8).toString('hex');
    return slug;
  }
}
