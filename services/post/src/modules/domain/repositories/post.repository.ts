import { Injectable } from '@nestjs/common';
import {
  CreatePostInput,
  InsertPostInput,
  Post,
} from '../aggregates/post/post.model';
import {
  DeletePostInput,
  LikePostInput,
  UnlikePostInput,
  User,
} from '../aggregates/user/user.model';
import { InsertCommentInput } from '../aggregates/comment/comment.model';

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

  async isLiked(post: Post, user: User): Promise<UserPostLike> {
    return await this._isLiked(post, user);
  }

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

  async createPost(postInput: CreatePostInput): Promise<Post | undefined> {
    const newPost = await this._createPost(Post.createPre(postInput));
    if (newPost) this.cachePost(newPost);
    return newPost;
  }

  async updatePost(target: Post): Promise<Post | undefined> {
    const updatedPost = await this._updatePost(target);
    if (updatedPost) this.cachePost(updatedPost);
    return updatedPost;
  }

  async deletePost(dpi: DeletePostInput): Promise<Post | undefined> {
    this.flushPost(dpi.post);
    return await this._deletePost(dpi.post);
  }

  async getLikeNum(target: Post): Promise<number> {
    return this._getLikeNum(target);
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

  async likePost(lpi: LikePostInput) {
    if (!lpi) return;
    return this._likePost(lpi);
  }
  async unlikePost(upi: UnlikePostInput) {
    if (!upi) return;
    return this._unlikePost(upi);
  }

  async likeComment(comment: Comment, user: User) {}
  async unlikeComment(comment: Comment, user: User) {}

  abstract _isLiked(post: Post, user: User): Promise<UserPostLike>;

  abstract _getAllPosts(skip: number): Promise<Post[]>;
  abstract _getUsersPosts(users: User[], skip: number): Promise<Post[]>;

  abstract _getPostByID(pid: bigint): Promise<Post | undefined>;
  abstract _getPostBySlug(slug: string): Promise<Post | undefined>;

  abstract _createPost(newPost: InsertPostInput): Promise<Post | undefined>;
  abstract _updatePost(traget: Post): Promise<Post | undefined>;
  abstract _deletePost(traget: Post): Promise<Post | undefined>;

  abstract _getLikeNum(traget: Post): Promise<number>;

  abstract _createComment(
    comment: InsertCommentInput,
  ): Promise<Comment | undefined>;
  abstract _updateComment(comment: Comment): Promise<Comment | undefined>;
  abstract _deleteComment(comment: Comment): Promise<Comment | undefined>;

  abstract _likePost(lpi: LikePostInput): Promise<Post | undefined>;
  abstract _unlikePost(upi: UnlikePostInput): Promise<Post | undefined>;

  abstract _likeComment(
    comment: Comment,
    user: User,
  ): Promise<Comment | undefined>;
  abstract _unlikeComment(
    comment: Comment,
    user: User,
  ): Promise<Comment | undefined>;
}
