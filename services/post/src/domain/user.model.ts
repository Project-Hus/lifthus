// task.service.ts
import { Injectable } from '@nestjs/common';
import { CreateWaitingPostModelInput, Post, WaitingPost } from './post.model';
import { Comment } from './comment.model';

import {
  CreatePostDto,
  CreateWaitingPostDto,
  PostLikeDto,
  PostUnlikeDto,
  UpdatePostDto,
} from './dto/post.dto';

interface IUser {
  getId(): bigint;

  createTmpPost(post: CreatePostDto): WaitingPost;
  createPost(post: Post): Post;
  updatePost(post: Post, updateData: UpdatePostDto): Post;
  deletePost(post: Post): bigint;

  likePost(post: Post): PostLikeDto;
  unlikePost(post: Post): PostUnlikeDto;

  // writeComment(comment: CreateCommentDto): Comment;
  // updateComment(comment: UpdateCommentDto): Comment;
  // deleteComment(comment: QueryCommentDto): bigint;

  // likeComment(comment: QueryCommentDto): BigInt;
  // unlikeComment(comment: QueryCommentDto): BigInt;
}

export type CreateUserModelInput = {
  id: bigint;
  posts?: Post[];
  comments?: Comment[];
  postLikes?: bigint[];
  commentLikes?: bigint[];
};

@Injectable()
export class User implements IUser {
  private readonly id: bigint;
  private waitingPosts: WaitingPost[];
  private posts: Post[];
  private comments: Comment[];
  private postLikes: bigint[];
  private commentLikes: bigint[];

  constructor(user: CreateUserModelInput) {
    this.id = user.id;
    this.posts = user.posts;
    this.comments = user.comments;
  }

  getId() {
    return this.id;
  }

  createTmpPost(post: CreateWaitingPostDto): WaitingPost {
    if (this.id !== post.author) throw new Error('invalid author');
    const newWaitingPost = new WaitingPost({
      author: this,
      content: post.content,
      images: post.srcs,
    });
    this.waitingPosts.push(newWaitingPost);
    return newWaitingPost;
  }
  createPost(post: Post): Post {
    if (this.id !== post.getAuthor().getId()) throw new Error('invalid author');
    this.posts.push(post);
    return post;
  }

  updatePost(post: Post, updateData: UpdatePostDto): Post {
    if (this.id !== post.getAuthor().getId()) throw new Error('invalid author');
    return post.update(updateData);
  }
  deletePost(post: Post): bigint {
    if (this.id !== post.getAuthor().getId()) throw new Error('invalid author');
    return post.getId();
  }

  likePost(post: Post): PostLikeDto {
    if (post.isLikedBy(this)) throw new Error('already liked');
    post.like(this);
    return {
      userId: this.id,
      postId: post.getId(),
    };
  }
  unlikePost(post: Post): PostUnlikeDto {
    if (!post.isLikedBy(this)) throw new Error('not liked');
    post.unlike(this);
    return {
      userId: this.id,
      postId: post.getId(),
    };
  }

  // writeComment = (comment: CreateCommentDto): Comment => {
  //   return Promise.reject(new Error('Method not implemented.'));
  // };
  // updateComment = (comment: UpdateCommentDto): Comment => {
  //   return Promise.reject(new Error('Method not implemented.'));
  // };
  // deleteComment = (comment: QueryCommentDto): bigint => {
  //   return Promise.reject(new Error('Method not implemented.'));
  // };
  // likeComment = (comment: QueryCommentDto): BigInt => {
  //   return Promise.reject(new Error('Method not implemented.'));
  // };
  // unlikeComment = (comment: QueryCommentDto): BigInt => {
  //   return Promise.reject(new Error('Method not implemented.'));
  // };
}
