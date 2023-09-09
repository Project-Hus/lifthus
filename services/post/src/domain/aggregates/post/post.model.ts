import crypto from 'crypto';
import { UpdatePostServiceDto } from 'src/dto/inbound/post.dto';
import {
  PostContents,
  PostIds,
  PostUpdates,
} from 'src/domain/aggregates/post/post.vo';
import { Timestamps } from 'src/domain/vo';

export type CreatePostInput = {
  slug: string;
  author: bigint;
  images: string[];
  content: string;

  id: bigint;
  createdAt: Date;
  updatedAt: Date;
};

export class Post {
  private id?: bigint;
  private createdAt?: Date;
  private updatedAt?: Date;
  private slug: string;
  private author: bigint;
  private imageSrcs: string[];
  private content: string;

  public static readonly SLUG_MAX_LENGTH: number = 100;

  private constructor() {}

  static create(author: bigint, contents: PostContents): Post {
    return new Post().create(author, contents);
  }

  private create(author: bigint, contents: PostContents): Post {
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

  update(updates: UpdatePostServiceDto): Post {
    this.content = updates.content;
    return this;
  }

  getUpdates(): PostUpdates {
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
    if (slugEnd == -1 || slugEnd > Post.SLUG_MAX_LENGTH) {
      slug = content.slice(0, Post.SLUG_MAX_LENGTH);
    } else {
      slug = content.slice(0, slugEnd);
    }
    slug = slug + 'code' + crypto.randomBytes(8).toString('hex');
    return slug;
  }
}
