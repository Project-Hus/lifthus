import { Injectable } from '@nestjs/common';

type PostSumm = {
  id: bigint;
  author: bigint;
  createdAt: Date;
  slug: string;
  likenum: number;
  commentnum: number;
  clientLiked: boolean;
};

@Injectable()
export class PostSummary {
  constructor(
    private id: bigint,
    private author: bigint,
    private createdAt: Date,
    private slug: string,
    private likenum: number,
    private commentnum: number,
    private clientLiked: boolean,
  ) {}
  getSumm(): PostSumm {
    return {
      id: this.id,
      author: this.author,
      createdAt: this.createdAt,
      slug: this.slug,
      likenum: this.likenum,
      commentnum: this.commentnum,
      clientLiked: this.clientLiked,
    };
  }
}
