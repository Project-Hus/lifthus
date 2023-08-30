// task.service.ts
import { Injectable } from '@nestjs/common';
import { Post, PrePostInput } from '../post/post.model';

import { UpdatePostDto } from '../../dto(later put out)/post.dto';
import { Comment } from '../comment/comment.model';
import {
  CommentLikeDto,
  CommentUnlikeDto,
  UpdateCommentDto,
} from '../../dto(later put out)/comment.dto';
import { bool } from 'aws-sdk/clients/signer';
import { UserPostLike } from '../../repositories/post.repository';

@Injectable()
export class User {
  constructor(private id: bigint) {}

  getID() {
    return this.id;
  }

  createPost(postData: PrePostInput): Post {
    return Post.createPre(postData, this);
  }

  updatePost(post: Post, updateData: UpdatePostDto): Post {
    if (this.id !== post.getAuthor().getID()) return;
    return post.update(updateData);
  }

  deletePost(post: Post): DeletePostInput {
    if (this.id !== post.getAuthor().getID()) return;
    return {
      deleter: this,
      post,
    };
  }

  likePost(upl: UserPostLike): LikePostInput {
    if (upl.liked) return { liker: this, post: u };
    post.like(this);
    return {
      userId: this.id,
      postId: post.getID(),
    };
  }

  unlikePost(post: Post): UnlikePostInput {
    if (!post.isLikedBy(this)) return;
    post.unlike(this);
    return {
      userId: this.id,
      postId: post.getID(),
    };
  }

  updateComment(comment: Comment, updateData: UpdateCommentDto): Comment {
    if (this.id !== comment.getAuthor().getID()) return;
    return comment.update(updateData);
  }

  deleteComment(comment: Comment): Comment {
    if (this.id !== comment.getAuthor().getID()) return;
    return comment;
  }

  likeComment(comment: Comment): CommentLikeDto {
    if (comment.isLikedBy(this)) return;
    comment.like(this);
    return {
      userId: this.id,
      commentId: comment.getID(),
    };
  }

  unlikeComment(comment: Comment): CommentUnlikeDto {
    if (!comment.isLikedBy(this)) throw new Error('not liked');
    comment.unlike(this);
    return {
      userId: this.id,
      commentId: comment.getID(),
    };
  }
}
