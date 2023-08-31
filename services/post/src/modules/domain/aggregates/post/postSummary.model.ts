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
  constructor(
    private id: bigint,
    private author: bigint,
    private createdAt: Date,
    private images: string[],
    private slug: string,
  ) {}
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
