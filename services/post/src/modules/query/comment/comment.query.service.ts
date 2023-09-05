import { Inject, Injectable } from '@nestjs/common';
import { Comment } from 'src/domain/aggregates/comment/comment.model';
import { CommentRepository } from 'src/modules/repositories/abstract/comment.repository';
import { PostRepository } from 'src/modules/repositories/abstract/post.repository';
import { CommentDto } from 'src/dto/outbound/comment.dto';
import { CommentLikeRepository } from 'src/modules/repositories/abstract/like.repository';

@Injectable()
export class CommentQueryService {
  constructor(
    @Inject(PostRepository) private readonly postRepo: PostRepository,
    @Inject(CommentRepository) private readonly commentRepo: CommentRepository,
    @Inject(CommentLikeRepository)
    private readonly likeRepo: CommentLikeRepository,
  ) {}

  async getComments({
    pid,
    skip,
  }: {
    pid: string;
    skip: number;
  }): Promise<CommentDto[]> {
    try {
      const comments: Comment[] = await this.commentRepo.getComments(
        BigInt(pid),
      );
      const commentDtos: CommentDto[] = await Promise.all(
        comments.map(async (c) => {
          const ln = await this.likeRepo.getLikesNum(c.getID());
          const rps = await Promise.all(
            c.getReplies().map(async (rp) => {
              const ln = await this.likeRepo.getLikesNum(rp.getID());
              return new CommentDto(rp, ln);
            }),
          );
          return new CommentDto(c, ln, rps);
        }),
      );
      return commentDtos;
    } catch (error) {
      return Promise.reject(error);
    }
  }
}
