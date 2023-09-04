import { Inject, Injectable } from '@nestjs/common';
import { Comment as PrismaComment } from '@prisma/client';
import { Comment } from 'src/domain/aggregates/comment/comment.model';
import { CommentRepository } from 'src/domain/repositories/comment.repository';
import { PostRepository } from 'src/domain/repositories/post.repository';
import { PrismaService } from 'src/prisma/prisma.service';

@Injectable()
export class CommentQueryService {
  constructor(
    @Inject(PrismaService) private readonly prismaService: PrismaService,
    @Inject(PostRepository) private readonly postRepo: PostRepository,
    @Inject(CommentRepository) private readonly commentRepo: CommentRepository,
  ) {}

  async getComments({ pid, skip }): Promise<Comment[]> {
    try {
      const post = await this.postRepo.getPostByID(pid);
      const comments = await this.commentRepo.getComments(post);
      return comments;
    } catch (error) {
      throw new Error('Failed to get comments');
    }
  }
}
