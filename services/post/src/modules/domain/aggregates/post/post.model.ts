// task.service.ts
import { Injectable } from '@nestjs/common';
import { User } from '../user/user.model';
import { UpdatePostDto } from '../../dto(later put out)/post.dto';

import crypto from 'crypto';

export type CreatePostInput = {
  author: User;
  images: string[];
  content: string;
};

export type UpdatePostInput = {
  id: bigint;
  content: string;
};

export type InsertPostInput = CreatePostInput & { slug: string };

@Injectable()
export class Post {
  private constructor(
    private id: bigint,
    private slug: string,

    private author: User,
    private images: string[],
    private content: string,

    private createdAt: Date,
    private updatedAt: Date,

    private likers: User[],
  ) {
    if (
      !id === undefined ||
      !slug === undefined ||
      !author === undefined ||
      !images === undefined ||
      !content === undefined ||
      !createdAt === undefined ||
      !updatedAt === undefined
    )
      throw new Error('invalid arguments');
  }

  static createPre(prePostInput: CreatePostInput): InsertPostInput {
    return {
      ...prePostInput,
      slug: Post.getSlug(prePostInput.content),
    };
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

  update(post: Post, changes: UpdatePostDto): Post {
    changes.content = post.content;
    return post;
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
