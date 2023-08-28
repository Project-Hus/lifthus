import { Injectable } from '@nestjs/common';
import { Post } from '../aggregates/post/post.model';

export interface PostRepository {
  createPost();
  updatePost();
  deletePost();
  likePost();
  unlikePost();

  createComment();
  updateComment();
  deleteComment();
  likeComment();
  unlikeComment();
}

@Injectable()
export class AbstractPostRepository implements PostRepository {
  private posts: Map<bigint, Post> = new Map(); // in-memory post storage

  constructor() {}
  createPost() {}
  updatePost() {}
  deletePost() {}
  likePost() {}
  unlikePost() {}

  createComment() {}
  updateComment() {}
  deleteComment() {}
  likeComment() {}
  unlikeComment() {}
}
