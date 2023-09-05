export class PostSummaryDto {
  id: string;
  author: string;
  createdAt: Date;
  updatedAt: Date;
  images: string[];
  slug: string;
  abstract: string;
  likesNum: number;
  commentsNum: number;
  constructor(psinput: PostSumamryDtoInput) {
    this.id = psinput.id.toString();
    this.author = psinput.author.toString();
    this.createdAt = psinput.createdAt;
    this.updatedAt = psinput.updatedAt;
    this.images = psinput.images;
    this.slug = psinput.slug;
    this.likesNum = psinput.likesNum;
    this.commentsNum = psinput.commentsNum;
    this.abstract = psinput.abstract;
  }

  static getAbstractFromSlug(slug: string) {
    slug = decodeURIComponent(slug);
    const codeIdx = slug.lastIndexOf('code');
    if (codeIdx === -1) return slug;
    return slug.slice(0, codeIdx);
  }
}

export type PostSumamryDtoInput = {
  id: bigint;
  author: bigint;
  createdAt: Date;
  updatedAt: Date;
  images: string[];
  slug: string;
  abstract: string;
  likesNum: number;
  commentsNum: number;
};
