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

  // updateComment(comment: Comment, updateData: UpdateCommentDto): Comment {
  //   if (this.id !== comment.getAuthor().getID()) return;
  //   return comment.update(updateData);
  // }

  // deleteComment(comment: Comment): Comment {
  //   if (this.id !== comment.getAuthor().getID()) return;
  //   return comment;
  // }

  // likeComment(comment: Comment): CommentLikeDto {
  //   if (comment.isLikedBy(this)) return;
  //   comment.like(this);
  //   return {
  //     userId: this.id,
  //     commentId: comment.getID(),
  //   };
  // }

  // unlikeComment(comment: Comment): CommentUnlikeDto {
  //   if (!comment.isLikedBy(this)) throw new Error('not liked');
  //   comment.unlike(this);
  //   return {
  //     userId: this.id,
  //     commentId: comment.getID(),
  //   };
  // }
}
