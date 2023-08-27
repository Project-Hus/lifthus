// task.service.ts
import { Injectable } from '@nestjs/common';
import crypto from 'crypto';
import { Post } from './post.model';
import { Comment } from './comment.model';
import { PostLike } from './postLike.model';
import { CommentLike } from './commentLike.model';

interface IUser {
  id: bigint;
  posts?: Post[];
  comments?: Comment[];
  postLikes?: PostLike[];
  commentLikes?: CommentLike[];

  writePost(any): Post;
  updatePost(any): Post;
  deletePost(any): bigint;

  likePost(any): PostLike;
  unlikePost(any): bigint;

  writeComment(any): Comment;
  updateComment(any): Comment;
  deleteComment(any): bigint;

  likeComment(any): CommentLike;
  unlikeComment(any): bigint;
}

@Injectable()
export class User implements IUser {
  id: bigint;
  posts?: Post[];
  comments?: Comment[];
  postLikes?: PostLike[];
  commentLikes?: CommentLike[];

  constructor(user: IUser) {
    this.id = user.id;
    this.posts = user.posts;
    this.comments = user.comments;
    this.postLikes = user.postLikes;
    this.commentLikes = user.commentLikes;
  }

  writePost: () => Post;
  updatePost: () => Post;
  deletePost: () => bigint;

  likePost: () => PostLike;
  unlikePost: () => bigint;

  writeComment: () => Comment;
  updateComment: () => Comment;
  deleteComment: () => bigint;

  likeComment: () => CommentLike;
  unlikeComment: () => bigint;
}
