// task.service.ts
import { Injectable } from '@nestjs/common';
import { Comment } from './comment.model';
import { User } from '../user/user.model';
import { UpdatePostDto } from '../../dto(later put out)/post.dto';

import crypto from 'crypto';

type IPost = {
  getID(): bigint | undefined;
  getAuthor(): User;

  update(updateData: UpdatePostDto): Post;
  getUpdatePostForm(): UpdatePostDto;

  isLikedBy(user: User): boolean;
  like(user: User): void;
  unlike(user: User): void;
};
@Injectable()
export class Post implements IPost {
  private constructor(
    private author: User,
    private images: string[],
    private content: string,
    private slug: string,

    private id?: bigint,

    private likenum?: number,
    private likers?: bigint[],

    private comments?: Comment[],

    private createdAt?: Date,
    private updatedAt?: Date,
  ) {}

  static PrePost(author: User, images: string[], content: string) {
    return new Post(author, images, content, this.getSlug(content));
  }

  static CreatedPost(
    author: User,
    images: string[],
    content: string,
    slug: string,
    id: bigint,
    likenum: number,
    likers: bigint[],
    comments: Comment[],
    createdAt: Date,
    updatedAt: Date,
  ) {
    return new Post(
      author,
      images,
      content,
      slug,
      id,
      likenum,
      likers,
      comments,
      createdAt,
      updatedAt,
    );
  }

  getID(): bigint | undefined {
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
