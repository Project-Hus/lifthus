// task.service.ts
import { Injectable } from '@nestjs/common';
import { Comment } from './comment.model';
import crypto from 'crypto';
import { User } from './user.model';

interface IPost {
  id?: bigint;
  createdAt?: Date;
  updatedAt?: Date;

  images: IImage[];
  userGroup?: bigint;
  author: User;
  slug: string;
  content: string;
  likenum: number;
  likes: IPostLike[];
  comments: Comment[];
}

interface IImage {
  id: bigint; // id and order unique
  order: number;
  src: string;
}

interface IPostLike {
  postId: bigint; // id with user
  user: bigint;
  createdAt: Date;
}

@Injectable()
export class Post {
  private post: IPost;

  constructor(post: IPost) {
    this.post = post;
  }

  private slugify(content: string): string {
    return encodeURIComponent(content + crypto.randomBytes(8).toString('hex'));
  }

  public isCreated(): boolean {
    return !!this.post.createdAt;
  }

  public getPost() {
    return this.post;
  }
}
