import { PostSummary } from 'src/domain/aggregates/post/postSummary.model';

export class PostSummaryDto {
  id: string;
  author: string;
  createdAt: Date;
  updatedAt: Date;
  imageSrcs: string[];
  slug: string;
  abstract: string;
  likesNum?: number;
  clientLiked?: boolean;
  commentsNum?: number;
  constructor(
    pse: PostSummary,
    likesNum?: number,
    clientLiked?: boolean,
    commentsNum?: number,
  ) {
    this.id = pse.id.toString();
    this.author = pse.author.toString();
    this.createdAt = pse.createdAt;
    this.updatedAt = pse.updatedAt;
    this.imageSrcs = pse.imageSrcs;
    this.slug = pse.slug;
    this.likesNum = likesNum;
    this.commentsNum = commentsNum;
    this.abstract = PostSummaryDto.getAbstractFromSlug(pse.slug);
    this.clientLiked = clientLiked;
  }

  static getAbstractFromSlug(slug: string) {
    slug = decodeURIComponent(slug);
    const codeIdx = slug.lastIndexOf('code');
    if (codeIdx === -1) return slug;
    return slug.slice(0, codeIdx);
  }
}
