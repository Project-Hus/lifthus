import { User } from './user.model';
import { Injectable } from '@nestjs/common';
import crypto from 'crypto';
import { CreatePostDto } from './dto/post.dto';

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
