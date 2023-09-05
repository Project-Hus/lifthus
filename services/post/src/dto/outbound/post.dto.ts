import { Post } from 'src/domain/aggregates/post/post.model';

export class PostDto {
  id: string;
  slug: string;

  author: string;

  images: string[];
  content: string;

  createdAt: Date;
  updatedAt: Date;

  likesNum: number;
  commentsNum: number;
  constructor(post: Post, likesNum: number, commentsNum: number) {
    this.id = post.getID().toString();
    this.author = post.getAuthor().toString();
    this.createdAt = post.getCreatedAt();
    this.updatedAt = post.getUpdatedAt();
    this.images = post.getImageSrcs();
    this.slug = post.getSlug();
    this.content = post.getContent();
    this.likesNum = likesNum;
    this.commentsNum = commentsNum;
  }
}
