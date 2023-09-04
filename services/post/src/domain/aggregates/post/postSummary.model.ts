import { Injectable } from '@nestjs/common';

export type PostSumm = {
  id: bigint;
  author: bigint;
  createdAt: Date;
  images: string[];
  slug: string;
};

@Injectable()
export class PostSummary {
  private id: bigint;
  private author: bigint;
  private createdAt: Date;
  private images: string[];
  private slug: string;

  static create(p: PostSumm): PostSummary {
    return new PostSummary().setPostSummary(p);
  }

  private setPostSummary(p: PostSumm): PostSummary {
    this.id = p.id;
    this.author = p.author;
    this.createdAt = p.createdAt;
    this.images = p.images;
    this.slug = p.slug;
    return this;
  }

  getSumm(): PostSumm {
    return {
      id: this.id,
      author: this.author,
      createdAt: this.createdAt,
      images: this.images,
      slug: this.slug,
    };
  }
}
