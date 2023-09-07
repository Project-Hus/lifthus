// task.service.ts
import {
  BadRequestException,
  ForbiddenException,
  Injectable,
} from '@nestjs/common';
import { Post } from '../post/post.model';

import { Comment, CreatePreCommentInput } from '../comment/comment.model';
import { Like } from '../like/like.model';
import {
  CreatePostServiceDto,
  UpdatePostServiceDto,
} from 'src/dto/inbound/post.dto';

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

  createPost(np: CreatePostServiceDto): Post {
    if (np.author !== this.id) throw ForbiddenException;
    return Post.create(np);
  }

  updatePost(postOriginal: Post, updates: UpdatePostServiceDto): Post {
    if (this.id !== postOriginal.getAuthor()) throw ForbiddenException;
    if (postOriginal.getID() !== updates.id) throw BadRequestException;
    return postOriginal.update(updates);
  }

  deletePost(post: Post): Post {
    if (this.id !== post.getAuthor()) throw ForbiddenException;
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
