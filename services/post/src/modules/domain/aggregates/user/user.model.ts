// task.service.ts
import { Injectable } from '@nestjs/common';
import { Post } from '../post/post.model';

import { Comment } from '../comment/comment.model';
import { Like } from '../like/like.model';

export type UserCreatePostInput = {
  images: string[];
  content: string;
};

export type UserUpdatePostInput = {
  content: string;
};

export type UserCreateCommentInput = {
  post: Post;
  parent?: Comment;
  content: string;
};

export type UserUpdateCommentInput = {
  content: string;
};

@Injectable()
export class User {
  constructor(private id: bigint) {}

  getID() {
    return this.id;
  }

  createPost({ images, content }: UserCreatePostInput): Post {
    return Post.createPre({ author: this, images, content });
  }

  updatePost(post: Post, changes: UserUpdatePostInput): Post | undefined {
    if (this.id !== post.getAuthor().getID()) return undefined;
    return post.update(changes);
  }

  deletePost(post: Post): Post | undefined {
    if (this.id !== post.getAuthor().getID()) return undefined;
    return post;
  }

  likePost(like: Like<Post>): Like<Post> | undefined {
    if (like.liker.getID() !== this.id || like.isLiked()) return undefined;
    return like.like(this);
  }

  unlikePost(like: Like<Post>): Like<Post> | undefined {
    if (like.liker.getID() !== this.id || !like.isLiked()) return undefined;
    return like.unlike(this);
  }

  createComment(c: UserCreateCommentInput): Comment {
    return Comment.createPre({ author: this, ...c });
  }

  updateComment(comment: Comment, changes: UserUpdateCommentInput): Comment {
    if (this.id !== comment.getAuthor().getID()) return;
    return comment.update(changes);
  }

  deleteComment(comment: Comment): Comment {
    if (this.id !== comment.getAuthor().getID()) return;
    return comment;
  }

  likeComment(like: Like<Comment>): Like<Comment> | undefined {
    if (like.liker.getID() !== this.id || like.isLiked()) return undefined;
    return like.like(this);
  }

  unlikeComment(like: Like<Comment>): Like<Comment> | undefined {
    if (like.liker.getID() !== this.id || !like.isLiked()) return undefined;
    return like.unlike(this);
  }
}
