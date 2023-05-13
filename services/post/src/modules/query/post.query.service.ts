import { Injectable } from '@nestjs/common';
import { Post } from '@prisma/client';
import { Response } from 'express';
import { PrismaService } from 'src/prisma/prisma.service';

@Injectable()
export class PostQueryService {
  constructor(private readonly prismaService: PrismaService) {}
  getHello(): string {
    return 'Hello World!';
  }

  async getUserPosts(): Promise<Post[]> {
    return this.prismaService.post.findMany();
  }
}
