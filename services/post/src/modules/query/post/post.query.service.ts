import { Inject, Injectable } from '@nestjs/common';
import { Post } from 'src/modules/domain/aggregates/post/post.model';
import { PostSummary } from 'src/modules/domain/aggregates/post/postSummary.model';
import { User } from 'src/modules/domain/aggregates/user/user.model';
import { PostRepository } from 'src/modules/domain/repositories/post.repository';
import { UserRepository } from 'src/modules/domain/repositories/user.repository';

import { PrismaService } from 'src/prisma/prisma.service';

@Injectable()
export class PostQueryService {
  constructor(
    @Inject(UserRepository) private readonly userRepo: UserRepository,
    @Inject(PostRepository) private readonly postRepo: PostRepository,
  ) {}
  getHello(): string {
    return 'Hello World!';
  }

  async getUsersPosts({
    users,
    skip,
  }: {
    users: number[];
    skip: number;
  }): Promise<PostSummary[]> {
    try {
      const targetUsers: User[] = [];
      users.forEach((uid) =>
        targetUsers.push(this.userRepo.getUser(BigInt(uid))),
      );
      return this.postRepo.getUsersPostSumms(targetUsers, skip);
    } catch (err) {
      throw err;
    }
  }

  getAllPosts(skip: number): Promise<PostSummary[]> {
    try {
      return this.postRepo.getAllPostSumms(skip);
    } catch (err) {
      throw err;
    }
  }

  /**
   * Gets post by slug.
   * @param slug
   * @returns
   */
  getPostBySlug(slug: string): Promise<Post> {
    try {
      return this.postRepo.getPostBySlug(slug);
    } catch (err) {
      throw err;
    }
  }

  getPostById(id: number): Promise<Post> {
    return this.postRepo.getPostByID(BigInt(id));
  }
}
