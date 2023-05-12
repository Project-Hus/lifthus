import { Injectable } from '@nestjs/common';
import { Post } from '@prisma/client';
import { Response } from 'express';
import { PostDBService } from 'src/prisma/post.db.service';

@Injectable()
export class PostQueryService {
  constructor(private readonly postDBService: PostDBService) {}
  getHello(): string {
    return 'Hello World!';
  }

  async getUserPosts(): Promise<Post[]> {
    return this.postDBService.posts({});
  }
}
