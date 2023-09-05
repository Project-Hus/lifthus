import { Comment } from 'src/domain/aggregates/comment/comment.model';

export class CommentDto {
  id: string;

  author: string;

  postId: string;
  parentId: string;

  content: string;

  createdAt: Date;
  updatedAt: Date;

  likesNum: number;

  replies?: Comment[];
  constructor(c: Comment, likesNum: number, replies?: Comment[]) {
    this.id = c.getID().toString();
    this.author = c.getAuthor().toString();

    this.postId = c.getPostID().toString();
    this.parentId = c.getParentID()?.toString();

    this.content = c.getContent();

    this.createdAt = c.getCreatedAt();
    this.updatedAt = c.getUpdatedAt();

    this.likesNum = likesNum;

    this.replies = replies;
  }
}
