// task.service.ts
import { Injectable } from '@nestjs/common';
import { Comment } from './comment.model';
import { User } from './user.model';
import { UpdatePostDto } from '../../dto/post.dto';

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

  update(updateData: UpdatePostDto): Post;
  getUpdatePostForm(): UpdatePostDto;

  isLikedBy(user: User): boolean;
  like(user: User): void;
  unlike(user: User): void;
}
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

  update(updateData: UpdatePostDto) {
    this.content = updateData.content;
    return this;
  }

  getUpdatePostForm(): UpdatePostDto {
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
