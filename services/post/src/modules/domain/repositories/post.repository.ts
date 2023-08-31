import { Injectable } from '@nestjs/common';
import { Post } from '../aggregates/post/post.model';
import { User } from '../aggregates/user/user.model';

import { PostSummary } from '../aggregates/post/postSummary.model';
import { stringifyAny } from 'src/common/utils/utils';

@Injectable()
export abstract class PostRepository {
  private posts: Map<bigint, Post> = new Map();
  private postOrigins: Map<bigint, string> = new Map();

  private getCachekey(p: Post): bigint {
    return BigInt(p.getID());
  }

  private getCacheString(p: Post): string {
    return stringifyAny(p);
  }

  private isUpdated(p: Post, origin: string): boolean {
    return this.getCacheString(p) !== origin;
  }

  async getAllPostSumms(skip: number): Promise<PostSummary[]> {
    return await this._getAllPostSumms(skip);
  }

  async getUsersPostSumms(users: User[], skip: number): Promise<PostSummary[]> {
    return await this._getUsersPostSumms(users, skip);
  }

  async getPostByID(pid: bigint): Promise<Post | undefined> {
    const post = await this._getPostByID(pid);
    const cacheKey = this.getCachekey(post);
    this.posts.set(cacheKey, post);
    this.postOrigins.set(cacheKey, this.getCacheString(post));
    return post;
  }

  async getPostBySlug(slug: string): Promise<Post | undefined> {
    const post = await this._getPostBySlug(slug);
    const cacheKey = this.getCachekey(post);
    this.posts.set(cacheKey, post);
    this.postOrigins.set(cacheKey, this.getCacheString(post));
    return post;
  }

  async createPost(post: Post): Promise<Post | undefined> {
    const newPost = await this._createPost(post);
    const cacheKey = this.getCachekey(newPost);
    this.posts.set(cacheKey, newPost);
    return newPost;
  }

  async deletePost(post: Post): Promise<Post | undefined> {
    const cacheKey = this.getCachekey(post);
    this.posts.delete(cacheKey);
    this.postOrigins.delete(cacheKey);
    return await this._deletePost(post);
  }

  abstract _getAllPostSumms(skip: number): Promise<PostSummary[]>;
  abstract _getUsersPostSumms(
    users: User[],
    skip: number,
  ): Promise<PostSummary[]>;

  abstract _getPostByID(pid: bigint): Promise<Post | undefined>;
  abstract _getPostBySlug(slug: string): Promise<Post | undefined>;

  abstract _createPost(post: Post): Promise<Post | undefined>;
  abstract _updatePost(traget: Post): Promise<Post | undefined>;
  abstract _deletePost(traget: Post): Promise<Post | undefined>;

  private clear() {
    this.posts.clear();
    this.postOrigins.clear();
  }

  async save(): Promise<void> {
    const changes: Set<Post> = new Set();
    this.posts.forEach((post, key) => {
      const origin = this.postOrigins.get(key);
      if (this.isUpdated(post, origin)) changes.add(post);
    });
    this.clear();
    return this._save(changes);
  }

  async cancel(): Promise<void> {
    this.clear();
    return this._save(new Set());
  }

  abstract _save(changes: Set<Post>): Promise<void>;
}
