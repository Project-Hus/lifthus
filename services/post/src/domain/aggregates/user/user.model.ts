// task.service.ts
import { BadRequestException, Injectable } from '@nestjs/common';
import { Post } from '../post/post.model';

import { Comment, CreatePreCommentInput } from '../comment/comment.model';
import { Like } from '../like/like.model';

export type UserCreatePostInput = {
  images: string[];
  content: string;
};

export type UserUpdatePostInput = {
  content: string;
};

export type UserCreatePreCommentInput = {
  postId: bigint;
  content: string;
};

export type UserUpdateCommentInput = {
  content: string;
};

@Injectable()
export class User {
  private id: bigint;

  static create(id: bigint): User {
    return new User().setID(id);
  }

  private setID(id: bigint): User {
    this.id = id;
    return this;
  }

  getID() {
    return this.id;
  }

  createPost({ images, content }: UserCreatePostInput): Post {
    return Post.create({ author: this.id, images, content });
  }

  updatePost(post: Post, changes: UserUpdatePostInput): Post {
    if (this.id !== post.getAuthor()) throw BadRequestException;
    return post.update(changes);
  }

  deletePost(post: Post): Post {
    if (this.id !== post.getAuthor()) throw BadRequestException;
    return post;
  }

  likePost(like: Like<Post>): Like<Post> {
    if (like.getLiker() !== this.id || like.isLiked())
      throw BadRequestException;
    return like.like(this);
  }

  unlikePost(like: Like<Post>): Like<Post> {
    if (like.getLiker() !== this.id || !like.isLiked())
      throw BadRequestException;
    return like.unlike(this);
  }

  createComment(c: UserCreatePreCommentInput): Comment {
    return Comment.createComment({ author: this.id, ...c });
  }

  updateComment(comment: Comment, changes: UserUpdateCommentInput): Comment {
    if (this.id !== comment.getAuthor()) return;
    return comment.update(changes);
  }

  deleteComment(comment: Comment): Comment {
    if (this.id !== comment.getAuthor()) return;
    return comment;
  }

  likeComment(like: Like<Comment>): Like<Comment> {
    if (like.getLiker() !== this.id || like.isLiked())
      throw BadRequestException;
    return like.like(this);
  }

  unlikeComment(like: Like<Comment>): Like<Comment> {
    if (like.getLiker() !== this.id || !like.isLiked())
      throw BadRequestException;
    return like.unlike(this);
  }
}
