import { Injectable, Logger } from '@nestjs/common';
import { Post } from '@prisma/client';
import { PrismaService } from 'src/prisma/prisma.service';

@Injectable()
export class PostQueryService {
  constructor(private readonly prismaService: PrismaService) {}
  getHello(): string {
    return 'Hello World!';
  }

  getAllPosts(): Promise<Post[]> {
    return Promise.reject('not implemented');
  }

  async getUserPosts(uid: number, skip: number): Promise<Post[]> {
    return this.prismaService.post.findMany({
      include: {
        images: {
          select: {
            id: true,
            url: true,
          },
          orderBy: {
            order: 'asc',
          },
        },
        comments: {
          select: {
            id: true,
            postId: true,
            author: true,
            createdAt: true,
            updatedAt: true,
            content: true,
            likenum: true,
            mentions: {
              select: {
                mentionee: true,
              },
            },
            replies: {
              select: {
                id: true,
                parentId: true,
                author: true,
                createdAt: true,
                updatedAt: true,
                content: true,
                likenum: true,
                mentions: {
                  select: {
                    mentionee: true,
                  },
                },
              },
            },
          },
          orderBy: {
            createdAt: 'desc',
          },
        },
        mentions: {
          select: {
            mentionee: true,
          },
        },
      },
      where: { author: uid },
      orderBy: {
        createdAt: 'desc',
      },
      take: 10,
      skip: skip,
    });
  }
}
