import { Injectable } from '@nestjs/common';
import { Post } from '../../../domain/aggregates/post/post.model';
import { User } from '../../../domain/aggregates/user/user.model';
import { PostUpdates } from 'src/domain/aggregates/post/post.vo';

@Injectable()
export abstract class PostRepository {
  async getAllPosts(skip: number): Promise<Post[]> {
    return await this._getAllPosts(skip);
  }

  async getUsersPosts(users: User[], skip: number): Promise<Post[]> {
    return await this._getUsersPosts(users, skip);
  }

  async getPostByID(pid: bigint): Promise<Post | null> {
    return await this._getPostByID(pid);
  }

  async getPostBySlug(slug: string): Promise<Post | null> {
    return await this._getPostBySlug(slug);
  }

  async createPost(post: Post): Promise<Post> {
    return await this._createPost(post);
  }

  async deletePost(post: Post): Promise<Post> {
    return await this._deletePost(post);
  }

  async save(post: Post): Promise<Post> {
    return this._save(post.getID(), post.getUpdates());
  }

  // Abstract methods to be implemented by the actual repository

  abstract _getAllPosts(skip: number): Promise<Post[]>;
  abstract _getUsersPosts(users: User[], skip: number): Promise<Post[]>;

  abstract _getPostByID(pid: bigint): Promise<Post | null>;
  abstract _getPostBySlug(slug: string): Promise<Post | null>;

  abstract _createPost(post: Post): Promise<Post>;
  abstract _deletePost(traget: Post): Promise<Post>;

  abstract _save(pid: bigint, updates: PostUpdates): Promise<Post>;
}
