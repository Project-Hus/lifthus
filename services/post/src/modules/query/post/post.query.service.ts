import { Inject, Injectable, Logger } from '@nestjs/common';
import { Post } from 'src/domain/aggregates/post/post.model';
import { PostSummary } from 'src/domain/aggregates/post/postSummary.model';
import { User } from 'src/domain/aggregates/user/user.model';
import { CommentRepository } from 'src/modules/repositories/abstract/comment.repository';
import { PostLikeRepository } from 'src/modules/repositories/abstract/like.repository';
import { PostRepository } from 'src/modules/repositories/abstract/post.repository';
import { UserRepository } from 'src/modules/repositories/abstract/user.repository';
import { PostDto } from 'src/dto/outbound/post.dto';
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

  async getUsersPosts({
    users,
    skip,
    client,
  }: {
    users: string[];
    skip: number;
    client: BigInt | undefined;
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
          const ln = await this.likeRepo.getLikesNum(pse.id);
          const cn = await this.commentRepo.getCommentsNum(pse.id);
          const clientLiked = !!client
            ? (await this.likeRepo.getLike(client, pse.id)).isLiked()
            : false;
          return new PostSummaryDto(pse, ln, clientLiked, cn);
        }),
      );
      return postSummDtos;
    } catch (err) {
      throw err;
    }
  }

  async getAllPosts(
    skip: number,
    client: BigInt | undefined,
  ): Promise<PostSummaryDto[]> {
    try {
      const PostSummEnts: PostSummary[] = await this.postRepo.getAllPostSumms(
        skip,
      );
      const postSummDtos: PostSummaryDto[] = await Promise.all(
        PostSummEnts.map(async (pse: PostSummary) => {
          const ln = await this.likeRepo.getLikesNum(pse.id);
          const cn = await this.commentRepo.getCommentsNum(pse.id);
          const clientLiked = !!client
            ? (await this.likeRepo.getLike(client, pse.id)).isLiked()
            : false;
          return new PostSummaryDto(pse, ln, clientLiked, cn);
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
  async getPostBySlug(
    slug: string,
    client: BigInt | undefined,
  ): Promise<PostDto> {
    try {
      const post = await this.postRepo.getPostBySlug(slug);
      const ln = await this.likeRepo.getLikesNum(post.getID());
      const cn = await this.commentRepo.getCommentsNum(post.getID());
      const clientLiked = !!client
        ? (await this.likeRepo.getLike(client, post.getID())).isLiked()
        : false;
      return new PostDto(post, ln, clientLiked, cn);
    } catch (err) {
      return Promise.reject(err);
    }
  }

  async getPostById(id: string, client: BigInt | undefined): Promise<PostDto> {
    try {
      const post = await this.postRepo.getPostByID(BigInt(id));
      const ln = await this.likeRepo.getLikesNum(post.getID());
      const cn = await this.commentRepo.getCommentsNum(post.getID());
      const clientLiked = !!client
        ? (await this.likeRepo.getLike(client, post.getID())).isLiked()
        : false;
      return new PostDto(post, ln, clientLiked, cn);
    } catch (err) {
      return Promise.reject(err);
    }
  }
}
