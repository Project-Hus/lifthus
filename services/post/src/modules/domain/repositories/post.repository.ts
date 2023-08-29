import { Injectable } from '@nestjs/common';
import {
  CreatePostInput,
  InsertPostInput,
  Post,
  UpdatePostInput,
} from '../aggregates/post/post.model';
import { User } from '../aggregates/user/user.model';
import { InsertCommentInput } from '../aggregates/post/comment.model';

@Injectable()
export abstract class PostRepository {
  /* ========== POST CACHE ========== */
  private postsCacheID: Map<bigint, Post> = new Map();
  private postsCacheSlug: Map<string, Post> = new Map();
  private expCacheTable: Map<bigint, number> = new Map();

  private allPostsCache: Post[] = [];
  private expAllPostsCache: number = 0;

  private usersPostsCache: Map<string, Post[]> = new Map();
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

  async getAllPosts(skip: number): Promise<Post[]> {
    if (Date.now() <= this.expAllPostsCache) return this.allPostsCache;
    this.allPostsCache = [];
    const allPosts = await this._getAllPosts(skip);
    this.allPostsCache = allPosts;
    this.expAllPostsCache = Date.now() + PostRepository.CACHE_EXPIRE_MILISEC;
    setTimeout(() => {
      this.allPostsCache = [];
      this.expAllPostsCache = 0;
    }, PostRepository.CACHE_EXPIRE_MILISEC);
    return allPosts;
  }

  async getUsersPosts(users: User[], skip: number): Promise<Post[]> {
    const cacheKey = this.getUsersCacheKey(users, skip);
    if (Date.now() <= this.expUsersPostsCache.get(cacheKey))
      return this.usersPostsCache.get(cacheKey)!;
    this.usersPostsCache.delete(cacheKey);
    const usersPosts = await this._getUsersPosts(users, skip);
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

  async createPost(postInput: CreatePostInput) {
    const newPost = await this._createPost(Post.getInsertInput(postInput));
    if (newPost) this.cachePost(newPost);
    return newPost;
  }

  async updatePost(post: Post) {
    const updatedPost = await this._updatePost(post);
    if (updatedPost) this.cachePost(updatedPost);
    return updatedPost;
  }

  async deletePost(post: Post) {
    this.flushPost(post);
    return await this._deletePost(post);
  }

  async createComment(newComment: InsertCommentInput) {
    return await this._createComment(newComment);
  }
  async updateComment(comment: Comment) {
    return await this._updateComment(comment);
  }
  async deleteComment(comment: Comment) {
    return await this._deleteComment(comment);
  }

  abstract _getAllPosts(skip: number): Promise<Post[]>;
  abstract _getUsersPosts(users: User[], skip: number): Promise<Post[]>;

  abstract _getPostByID(pid: bigint): Promise<Post | undefined>;
  abstract _getPostBySlug(slug: string): Promise<Post | undefined>;

  abstract _createPost(newPost: InsertPostInput): Promise<Post | undefined>;
  abstract _updatePost(post: Post): Promise<Post | undefined>;
  abstract _deletePost(post: Post): Promise<Post | undefined>;

  abstract _createComment(
    comment: InsertCommentInput,
  ): Promise<Comment | undefined>;
  abstract _updateComment(comment: Comment): Promise<Comment | undefined>;
  abstract _deleteComment(comment: Comment): Promise<Comment | undefined>;
}
