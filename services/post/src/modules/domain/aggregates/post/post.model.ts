// task.service.ts
import { Injectable } from '@nestjs/common';
import { Comment } from './comment.model';
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

export type LikePostInput = {
  id: bigint;
  userId: bigint;
};

export type UnlikePostInput = {
  id: bigint;
  userId: bigint;
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

    private likenum: number,
    private likers: Set<bigint>,

    private createdAt: Date,
    private updatedAt: Date,

    private comments?: Comment[],
  ) {}

  static getInsertInput(prePostInput: CreatePostInput): InsertPostInput {
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

  isLikedBy(user: User): boolean {
    return this.likers.has(user.getID());
  }

  update(updatePostInput: UpdatePostDto): UpdatePostInput {
    this.content = updatePostInput.content;
    return updatePostInput;
  }

  like(user: User): LikePostInput {
    if (this.likers.has(user.getID())) throw new Error('already liked');
    this.likenum++;
    return {
      id: this.getID(),
      userId: user.getID(),
    };
  }
  unlike(user: User): UnlikePostInput {
    if (!this.likers.has(user.getID())) throw new Error('not liked');
    this.likenum--;
    return {
      id: this.getID(),
      userId: user.getID(),
    };
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
