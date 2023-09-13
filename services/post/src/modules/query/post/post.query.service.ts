import { Inject, Injectable, Logger } from '@nestjs/common';
import { User } from 'src/domain/aggregates/user/user.model';
import { CommentRepository } from 'src/modules/repositories/abstract/comment.repository';
import { PostLikeRepository } from 'src/modules/repositories/abstract/like.repository';
import { PostRepository } from 'src/modules/repositories/abstract/post.repository';
import { UserRepository } from 'src/modules/repositories/abstract/user.repository';
import { PostDto } from 'src/dto/outbound/post.dto';
import { Post } from 'src/domain/aggregates/post/post.model';

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
  }): Promise<PostDto[]> {
    try {
      const targetUsers: User[] = [];
      users.forEach((uid) =>
        targetUsers.push(this.userRepo.getUser(BigInt(uid))),
      );
      const postEnts: Post[] = await this.postRepo.getUsersPosts(
        targetUsers,
        skip,
      );
      const postDtos: PostDto[] = await Promise.all(
        postEnts.map(async (p: Post) => {
          const ln = await this.likeRepo.getLikesNum(p.getID());
          const cn = await this.commentRepo.getCommentsNum(p.getID());
          const clientLiked = !!client
            ? (await this.likeRepo.getLike(client, p.getID())).isLiked()
            : false;
          return new PostDto(p, ln, clientLiked, cn);
        }),
      );
      return postDtos;
    } catch (err) {
      throw err;
    }
  }

  async getAllPosts(
    skip: number,
    client: BigInt | undefined,
  ): Promise<PostDto[]> {
    try {
      const postEnts: Post[] = await this.postRepo.getAllPosts(skip);
      const postDtos: PostDto[] = await Promise.all(
        postEnts.map(async (p: Post) => {
          const ln = await this.likeRepo.getLikesNum(p.getID());
          const cn = await this.commentRepo.getCommentsNum(p.getID());
          const clientLiked = !!client
            ? (await this.likeRepo.getLike(client, p.getID())).isLiked()
            : false;
          return new PostDto(p, ln, clientLiked, cn);
        }),
      );
      return postDtos;
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
