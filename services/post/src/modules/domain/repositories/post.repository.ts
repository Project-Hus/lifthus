import { Injectable } from '@nestjs/common';
import { CreatePostInput, Post } from '../aggregates/post/post.model';
import { User } from '../aggregates/user/user.model';

import { PostSummary } from '../aggregates/post/postSummary.model';

export type UserPostLike = {
  user: User;
  post: Post;
  liked: boolean;
};

@Injectable()
export abstract class PostRepository {
  /* ========== POST CACHE ========== */
  private postsCacheID: Map<bigint, Post> = new Map();
  private postsCacheSlug: Map<string, Post> = new Map();
  private expCacheTable: Map<bigint, number> = new Map();

  private allPostsCache: PostSummary[] = [];
  private expAllPostsCache: number = 0;

  private usersPostsCache: Map<string, PostSummary[]> = new Map();
  private expUsersPostsCache: Map<string, number> = new Map();

  static readonly CACHE_EXPIRE_MILISEC = 5000;
  private setExp(pid: bigint) {
    this.expCacheTable.set(
      pid,
      Date.now() + PostRepository.CACHE_EXPIRE_MILISEC,
    );
  }

  private cachePost(post: Post) {
    this.postsCacheID.set(post.getID(), post);
    this.postsCacheSlug.set(post.getSlug(), post);
    this.setExp(post.getID());
    setTimeout(() => {
      this.flushPost(post);
    }, PostRepository.CACHE_EXPIRE_MILISEC);
  }

  private flushPost(post: Post) {
    this.postsCacheID.delete(post.getID());
    this.postsCacheSlug.delete(post.getSlug());
    this.expCacheTable.delete(post.getID());
  }

  private getCacheByID(pid: bigint): Post | undefined {
    const post = this.postsCacheID.get(pid);
    if (!post) return undefined;
    const slug = post.getSlug();
    const pExp = this.expCacheTable.get(pid);
    if (pExp < Date.now()) {
      this.flushPost(post);
      return undefined;
    }
    return post;
  }

  private getCachedBySlug(slug: string): Post | undefined {
    const post = this.postsCacheSlug.get(slug);
    if (!post) return undefined;
    const pid = post.getID();
    const pExp = this.expCacheTable.get(pid);
    if (pExp < Date.now()) {
      this.flushPost(post);
      return undefined;
    }
    return post;
  }

  private getUsersCacheKey(users: User[], skip: number): string {
    let cacheKey = '';
    for (const u of users) {
      cacheKey += u.getID().toString() + ',';
    }
    cacheKey += skip.toString();
    return cacheKey;
  }
  /* ==================== */

  async getAllPostSumms(skip: number): Promise<PostSummary[]> {
    if (Date.now() <= this.expAllPostsCache) return this.allPostsCache;
    this.allPostsCache = [];
    const allPosts = await this._getAllPostSumms(skip);
    this.allPostsCache = allPosts;
    this.expAllPostsCache = Date.now() + PostRepository.CACHE_EXPIRE_MILISEC;
    setTimeout(() => {
      this.allPostsCache = [];
      this.expAllPostsCache = 0;
    }, PostRepository.CACHE_EXPIRE_MILISEC);
    return allPosts;
  }

  async getUsersPostSumms(users: User[], skip: number): Promise<PostSummary[]> {
    const cacheKey = this.getUsersCacheKey(users, skip);
    if (Date.now() <= this.expUsersPostsCache.get(cacheKey))
      return this.usersPostsCache.get(cacheKey)!;
    this.usersPostsCache.delete(cacheKey);
    const usersPosts = await this._getUsersPostSumms(users, skip);
    this.usersPostsCache.set(cacheKey, usersPosts);
    setTimeout(() => {
      this.usersPostsCache.delete(cacheKey);
      this.expUsersPostsCache.delete(cacheKey);
    }, PostRepository.CACHE_EXPIRE_MILISEC);
    return usersPosts;
  }

  async getPostByID(pid: bigint): Promise<Post | undefined> {
    const cachedPost = this.getCacheByID(pid);
    if (cachedPost) return cachedPost;
    const post = await this._getPostByID(pid);
    if (!post) return undefined;
    this.cachePost(post);
    return post;
  }

  async getPostBySlug(slug: string): Promise<Post | undefined> {
    const cachedPost = this.getCachedBySlug(slug);
    if (cachedPost) return cachedPost;
    const post = await this._getPostBySlug(slug);
    if (!post) return undefined;
    this.cachePost(post);
    return post;
  }

  async createPost(post: Post): Promise<Post | undefined> {
    const newPost = await this._createPost(post);
    return newPost;
  }

  async updatePost(target: Post): Promise<Post | undefined> {
    const updatedPost = await this._updatePost(target);
    if (updatedPost) this.cachePost(updatedPost);
    return updatedPost;
  }

  async deletePost(post: Post): Promise<Post | undefined> {
    this.flushPost(post);
    return await this._deletePost(post);
  }

  async likePost(post: Post, user: User) {
    return this._likePost(post, user);
  }
  async unlikePost(post: Post, user: User) {
    return this._unlikePost(post, user);
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

  abstract _likePost(post: Post, user: User): Promise<Post | undefined>;
  abstract _unlikePost(post: Post, user: User): Promise<Post | undefined>;
}
