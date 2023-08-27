// task.service.ts
import { Injectable } from '@nestjs/common';
import { Comment } from './comment.model';

import { User } from './user.model';
import { CreatePostDto } from './dto/post.dto';

import crypto from 'crypto';

/* model for created posts */

export type CreatePostModelInput = {
  id: bigint;
  slug: string;

  author: User;

  images: string[];
  content: string;

  likenum: number;
  likers: bigint[];

  comments?: Comment[];

  createdAt: Date;
  updatedAt: Date;
};

interface IPost {
  getID(): bigint;
  getAuthor(): User;

  update(updateData: UpdatePostForm);
  getUpdatePostForm(): UpdatePostForm;

  isLikedBy(user: User): boolean;
  like(user: User): void;
  unlike(user: User): void;
}

export type UpdatePostForm = {
  id: bigint;
  content: string;
};
@Injectable()
export class Post implements IPost {
  private id: bigint;
  private slug: string;

  private author: User;

  private images: string[];
  private content: string;

  private likenum: number;
  likers: bigint[];

  private comments?: Comment[];

  private createdAt: Date;
  private updatedAt: Date;

  constructor(post: CreatePostModelInput) {
    this.id = post.id;
    this.slug = post.slug;

    this.author = post.author;

    this.images = post.images;
    this.content = post.content;

    this.likenum = post.likenum;

    this.comments = post.comments;

    this.createdAt = post.createdAt;
    this.updatedAt = post.updatedAt;
  }

  getID(): bigint {
    return this.id;
  }

  getAuthor(): User {
    return this.author;
  }

  update(updateData: UpdatePostForm) {
    this.content = updateData.content;
    return this;
  }

  getUpdatePostForm(): UpdatePostForm {
    return {
      id: this.id,
      content: this.content,
    };
  }

  isLikedBy(user: User): boolean {
    return this.likers.includes(user.getID());
  }
  like(user: User): void {
    if (this.likers.includes(user.getID())) throw new Error('already liked');
    this.likers.push(user.getID());
    this.likenum++;
  }
  unlike(user: User): void {
    if (!this.likers.includes(user.getID())) throw new Error('not liked');
    this.likers = this.likers.filter((liker) => liker !== user.getID());
    this.likenum--;
  }
}

/* model for posts waiting to be created */
interface IPrePost {
  getCreatePostForm(): CreatePostDto;
}

export type CreatePrePostModelInput = {
  author: User;
  images: string[];
  content: string;
};

@Injectable()
export class PrePost implements IPrePost {
  private readonly author: User;
  private readonly images: string[];
  private readonly content: string;

  private readonly slug: string;
  private readonly likenum: number;

  constructor(post: CreatePrePostModelInput) {
    this.author = post.author;
    this.images = post.images;
    this.content = post.content;
    this.slug = this.getSlug(post.content);
    this.likenum = 0;
  }

  public getCreatePostForm(): CreatePostDto {
    return {
      author: this.author.getID(),
      srcs: [...this.images],
      content: this.content,
      slug: this.slug,
      likenum: this.likenum,
    };
  }

  private getSlug(content: string): string {
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
