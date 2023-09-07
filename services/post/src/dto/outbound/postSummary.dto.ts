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
    const ps = pse.getSumm();
    this.id = ps.id.toString();
    this.author = ps.author.toString();
    this.createdAt = ps.createdAt;
    this.updatedAt = ps.updatedAt;
    this.imageSrcs = ps.imageSrcs;
    this.slug = ps.slug;
    this.likesNum = likesNum;
    this.commentsNum = commentsNum;
    this.abstract = PostSummaryDto.getAbstractFromSlug(ps.slug);
    this.clientLiked = clientLiked;
  }

  static getAbstractFromSlug(slug: string) {
    slug = decodeURIComponent(slug);
    const codeIdx = slug.lastIndexOf('code');
    if (codeIdx === -1) return slug;
    return slug.slice(0, codeIdx);
  }
}
