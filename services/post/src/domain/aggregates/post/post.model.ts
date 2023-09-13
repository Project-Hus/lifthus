import crypto from 'crypto';
import {
  PostContents,
  PostIds,
  PostUpdates,
} from 'src/domain/aggregates/post/post.vo';
import { Timestamps } from 'src/domain/vo';
import { BadRequestException, ForbiddenException } from '@nestjs/common';
import {
  POST_MAX_CONTENT_LENGTH,
  POST_SLUG_LENGTH,
} from 'src/shared/constraints';

export class Post {
  private id?: bigint;
  private createdAt?: Date;
  private updatedAt?: Date;
  private slug: string;
  private author: bigint;
  private imageSrcs: string[];
  private content: string;

  private static readonly SLUG_LENGTH: number = POST_SLUG_LENGTH;
  private static readonly MAX_CONTENT_LENGTH = POST_MAX_CONTENT_LENGTH;

  private constructor() {}

  static create(author: bigint, contents: PostContents): Post {
    return new Post().create(author, contents);
  }

  private create(author: bigint, contents: PostContents): Post {
    if (contents.content.length > Post.MAX_CONTENT_LENGTH) {
      throw BadRequestException;
    }
    this.slug = Post.generateSlug(contents.content);
    this.author = author;
    this.imageSrcs = contents.imageSrcs;
    this.content = contents.content;
    return this;
  }

  static from(
    author: bigint,
    postIds: PostIds,
    postContents: PostContents,
    timestamps: Timestamps,
  ): Post {
    return new Post().from(author, postIds, postContents, timestamps);
  }

  private from(
    author: bigint,
    postIds: PostIds,
    postContents: PostContents,
    timestamps: Timestamps,
  ): Post {
    this.id = postIds.id;
    this.slug = postIds.slug;
    this.author = author;
    this.imageSrcs = postContents.imageSrcs;
    this.content = postContents.content;
    this.createdAt = timestamps.createdAt;
    this.updatedAt = timestamps.updatedAt;
    return this;
  }

  isPersisted(): boolean {
    return !!this.id;
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

  update(updates: PostUpdates): Post {
    if (updates.content.length > Post.MAX_CONTENT_LENGTH)
      throw BadRequestException;
    this.content = updates.content;
    return this;
  }

  getUpdates(): PostUpdates {
    return {
      content: this.content,
    };
  }

  delete(deleter: bigint): Post {
    if (this.author !== deleter) throw ForbiddenException;
    return this;
  }

  private static generateSlug(content: string): string {
    let slug: string;
    const slugEnd: number = content.indexOf('\n');
    if (slugEnd == -1 || slugEnd > Post.SLUG_LENGTH) {
      slug = content.slice(0, Post.SLUG_LENGTH);
    } else {
      slug = content.slice(0, slugEnd);
    }
    slug = slug + 'code' + crypto.randomBytes(8).toString('hex');
    return slug;
  }
}
