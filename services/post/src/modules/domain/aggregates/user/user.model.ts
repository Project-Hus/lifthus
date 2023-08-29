// task.service.ts
import { Injectable } from '@nestjs/common';
import { Post } from '../post/post.model';

import {
  CreatePrePostDto,
  PostLikeDto,
  PostUnlikeDto,
  UpdatePostDto,
} from '../../dto(later put out)/post.dto';
import { Comment } from '../post/comment.model';
import {
  CommentLikeDto,
  CommentUnlikeDto,
  CreatePreCommentDto,
  UpdateCommentDto,
} from '../../dto(later put out)/comment.dto';

interface IUser {
  getID(): bigint;

  updatePost(post: Post, updateData: UpdatePostDto): Post;
  deletePost(post: Post): Post;

  likePost(post: Post): PostLikeDto;
  unlikePost(post: Post): PostUnlikeDto;

  updateComment(comment: Comment, updateData: UpdateCommentDto): Comment;
  deleteComment(comment: Comment): Comment;

  likeComment(comment: Comment): CommentLikeDto;
  unlikeComment(comment: Comment): CommentUnlikeDto;
}

export type CreateUserModelInput = {
  id: bigint;
};

/**
 * @description
 * User domain model
 * @class User
 * @implements {IUser}
 */
@Injectable()
export class User implements IUser {
  private readonly id: bigint;

  constructor(user: CreateUserModelInput) {
    this.id = user.id;
  }

  getID() {
    return this.id;
  }

  /**
   * @description
   * takes the data to update the given post model and returns the updated post model's reference.
   * @param post
   * @param updateData
   * @returns
   */
  updatePost(post: Post, updateData: UpdatePostDto): Post {
    if (this.id !== post.getAuthor().getID()) return;
    return post.update(updateData);
  }

  /**
   * @description
   * returns the post model's reference to be deleted if the user is the author of the post.
   * @param post
   * @returns
   */
  deletePost(post: Post): Post {
    if (this.id !== post.getAuthor().getID()) return;
    return post;
  }

  /**
   * @description
   * returns the PostLikeDto if the user has not liked the post yet.
   * @param post
   * @returns
   */
  likePost(post: Post): PostLikeDto {
    if (post.isLikedBy(this)) return;
    post.like(this);
    return {
      userId: this.id,
      postId: post.getID(),
    };
  }

  /**
   * @description
   * returns the PostUnlikeDto if the user has liked the post.
   * @param post
   * @returns
   */
  unlikePost(post: Post): PostUnlikeDto {
    if (!post.isLikedBy(this)) return;
    post.unlike(this);
    return {
      userId: this.id,
      postId: post.getID(),
    };
  }

  /**
   * @description
   * takes the data to update the given comment model and returns the updated comment model's reference.
   * @param comment
   * @param updateData
   * @returns
   */
  updateComment(comment: Comment, updateData: UpdateCommentDto): Comment {
    if (this.id !== comment.getAuthor().getID()) return;
    return comment.update(updateData);
  }

  /**
   * @description
   * returns the comment model's reference to be deleted if the user is the author of the comment.
   * @param comment
   * @returns
   */
  deleteComment(comment: Comment): Comment {
    if (this.id !== comment.getAuthor().getID()) return;
    return comment;
  }

  /**
   * @description
   * returns the CommentLikeDto if the user has not liked the comment yet.
   * @param comment
   * @returns
   */
  likeComment(comment: Comment): CommentLikeDto {
    if (comment.isLikedBy(this)) return;
    comment.like(this);
    return {
      userId: this.id,
      commentId: comment.getID(),
    };
  }

  /**
   * @description
   * returns the CommentUnlikeDto if the user has liked the comment.
   * @param comment
   * @returns
   */
  unlikeComment(comment: Comment): CommentUnlikeDto {
    if (!comment.isLikedBy(this)) throw new Error('not liked');
    comment.unlike(this);
    return {
      userId: this.id,
      commentId: comment.getID(),
    };
  }
}
