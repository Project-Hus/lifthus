import { ForbiddenException, Injectable } from '@nestjs/common';
import { Post } from '../post/post.model';

import { Comment } from '../comment/comment.model';
import { Like } from '../like/like.model';

import { PostContents, PostUpdates } from 'src/domain/aggregates/post/post.vo';
import {
  CommentParents,
  CommentUpdates,
} from 'src/domain/aggregates/comment/comment.vo';

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

  createPost(author: bigint, contents: PostContents): Post {
    if (author !== this.id) throw ForbiddenException;
    return Post.create(author, contents);
  }

  updatePost(originalPost: Post, updates: PostUpdates): Post {
    if (this.id !== originalPost.getAuthor()) throw ForbiddenException;
    return originalPost.update(updates);
  }

  deletePost(post: Post): Post {
    if (this.id !== post.getAuthor()) throw ForbiddenException;
    return post.delete(this.id);
  }

  likePost(like: Like<Post>): Like<Post> {
    if (like.getLiker() !== this.id || like.isLiked()) throw ForbiddenException;
    return like.like(this);
  }

  unlikePost(like: Like<Post>): Like<Post> {
    if (like.getLiker() !== this.id || !like.isLiked())
      throw ForbiddenException;
    return like.unlike(this);
  }

  createComment(
    author: bigint,
    parents: CommentParents,
    content: string,
  ): Comment {
    if (this.id !== author) throw ForbiddenException;
    return Comment.create(author, parents, content);
  }

  updateComment(comment: Comment, changes: CommentUpdates): Comment {
    if (this.id !== comment.getAuthor()) throw ForbiddenException;
    return comment.update(changes);
  }

  deleteComment(comment: Comment): Comment {
    if (this.id !== comment.getAuthor()) throw ForbiddenException;
    return comment.delete(this.id);
  }

  likeComment(like: Like<Comment>): Like<Comment> {
    if (like.getLiker() !== this.id || like.isLiked()) throw ForbiddenException;
    return like.like(this);
  }

  unlikeComment(like: Like<Comment>): Like<Comment> {
    if (like.getLiker() !== this.id || !like.isLiked())
      throw ForbiddenException;
    return like.unlike(this);
  }
}
