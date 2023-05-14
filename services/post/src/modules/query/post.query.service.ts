import { Injectable } from '@nestjs/common';
import { Post, Prisma } from '@prisma/client';
import { Response } from 'express';
import { PrismaService } from 'src/prisma/prisma.service';
import { PostQueryDto } from './post.query.dto';

@Injectable()
export class PostQueryService {
  constructor(private readonly prismaService: PrismaService) {}
  getHello(): string {
    return 'Hello World!';
  }

  a: Prisma.PostSelect;

  async getUserPosts(uid: number, skip: number): Promise<PostQueryDto[]> {
    return this.prismaService.post.findMany({
      select: {
        id: true,
        author: true,
        createdAt: true,
        updatedAt: true,
        slug: true,
        content: true,
        likenum: true,
      },
      where: {},
      orderBy: {
        createdAt: 'desc',
      },
      take: 10,
      skip: skip,
    });
  }
}
