import { Inject, Injectable } from '@nestjs/common';
import { Post } from 'src/domain/aggregates/post/post.model';
import { PostSummary } from 'src/domain/aggregates/post/postSummary.model';
import { User } from 'src/domain/aggregates/user/user.model';
import { CommentRepository } from 'src/domain/repositories/comment.repository';
import { PostLikeRepository } from 'src/domain/repositories/like.repository';
import { PostRepository } from 'src/domain/repositories/post.repository';
import { UserRepository } from 'src/domain/repositories/user.repository';
import { PostSummaryDto } from 'src/dto/outbound/postSummary.dto';

@Injectable()
export class PostQueryService {
  constructor(
    @Inject(UserRepository) private readonly userRepo: UserRepository,
    @Inject(PostRepository) private readonly postRepo: PostRepository,
    @Inject(CommentRepository) private readonly commentRepo: CommentRepository,
    @Inject(PostLikeRepository)
    private readonly likeRepo: PostLikeRepository,
  ) {}
  getHello(): string {
    return 'Hello World!';
  }

  async getUsersPosts({
    users,
    skip,
  }: {
    users: string[];
    skip: number;
  }): Promise<PostSummaryDto[]> {
    try {
      const targetUsers: User[] = [];
      users.forEach((uid) =>
        targetUsers.push(this.userRepo.getUser(BigInt(uid))),
      );
      const PostSummEnts: PostSummary[] = await this.postRepo.getUsersPostSumms(
        targetUsers,
        skip,
      );
      const postSummDtos: PostSummaryDto[] = await Promise.all(
        PostSummEnts.map(async (pse: PostSummary) => {
          const ps = pse.getSumm();
          const ln = await this.likeRepo.getLikesNum(ps.id);
          const cn = await this.commentRepo.getCommentsNum(ps.id);
          return new PostSummaryDto(pse, ln, cn);
        }),
      );
      return postSummDtos;
    } catch (err) {
      throw err;
    }
  }

  async getAllPosts(skip: number): Promise<PostSummaryDto[]> {
    try {
      const PostSummEnts: PostSummary[] = await this.postRepo.getAllPostSumms(
        skip,
      );
      const postSummDtos: PostSummaryDto[] = await Promise.all(
        PostSummEnts.map(async (pse: PostSummary) => {
          const ps = pse.getSumm();
          const ln = await this.likeRepo.getLikesNum(ps.id);
          const cn = await this.commentRepo.getCommentsNum(ps.id);
          return new PostSummaryDto(pse, ln, cn);
        }),
      );
      return postSummDtos;
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
