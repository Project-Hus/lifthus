import { Injectable } from '@nestjs/common';
import { Comment } from '@prisma/client';
import { PrismaService } from 'src/prisma/prisma.service';

@Injectable()
export class CommentQueryService {
  constructor(private prismaService: PrismaService) {}

  /**
   * Get comments of specified post.
   * @param param0
   * @returns
   */
  async getComments({
    pid,
    skip,
  }: {
    pid: number;
    skip: number;
  }): Promise<Comment[]> {
    try {
      const comments = await this.prismaService.comment.findMany({
        where: {
          postId: pid,
        },
        skip: skip,
        include: {
        
          parent: true,
          replies: true,
          mentions: {
            select: {
              mentionee: true,
            },
          },
        },
      });

      return comments;
    } catch (error) {
      throw new Error('Failed to get comments');
    }
  }
}
