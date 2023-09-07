import { Injectable } from '@nestjs/common';

export type PostSumm = {
  id: bigint;
  author: bigint;
  createdAt: Date;
  updatedAt: Date;
  imageSrcs: string[];
  slug: string;
};

@Injectable()
export class PostSummary {
  private id: bigint;
  private author: bigint;
  private createdAt: Date;
  private updatedAt: Date;
  private imageSrcs: string[];
  private slug: string;

  static create(p: PostSumm): PostSummary {
    return new PostSummary().setPostSummary(p);
  }

  private setPostSummary(p: PostSumm): PostSummary {
    this.id = p.id;
    this.author = p.author;
    this.createdAt = p.createdAt;
    this.updatedAt = p.updatedAt;
    this.imageSrcs = p.imageSrcs;
    this.slug = p.slug;
    return this;
  }

  getSumm(): PostSumm {
    return {
      id: this.id,
      author: this.author,
      createdAt: this.createdAt,
      updatedAt: this.updatedAt,
      imageSrcs: this.imageSrcs,
      slug: this.slug,
    };
  }
}
