import {
  Controller,
  Get,
  Param,
  Query,
  BadRequestException,
} from '@nestjs/common';
import { PostQueryService } from './post.query.service';
import { Post as PrismaPost } from '@prisma/client';
import { Post } from 'src/modules/domain/aggregates/post/post.model';
import { PostSummary } from 'src/modules/domain/aggregates/post/postSummary.model';

@Controller('/post/query/post')
export class PostQueryController {
  constructor(private readonly postQueryService: PostQueryService) {}

  @Get('/slug/:slug')
  getPostBySlug(@Param('slug') slug: string): Promise<Post> {
    return this.postQueryService.getPostBySlug(slug);
  }

  @Get('/id/:id')
  getPostById(@Param('id') idStr: string): Promise<Post> {
    const id = Number(idStr);
    if (isNaN(id)) {
      throw new BadRequestException();
    }
    return this.postQueryService.getPostById(id);
  }

  /**
   * gets all posts from the database and returns them.
   * @param skip
   * @returns Post[]
   * @example /post/query/post/all/0
   */
  @Get('/all')
  getAllPosts(@Query('skip') skipStr: string): Promise<PostSummary[]> {
    const skip = Number(skipStr) || 0;
    return this.postQueryService.getAllPosts(Number(skip));
  }

  /**
   * gets users query param and returns posts from the users.
   * @param users
   * @returns Post[]
   * @example /post/query/post?users=1,2,3
   */
  @Get('')
  getUsersPosts(
    @Query('users') usersStr: string,
    @Query('skip') skipStr: string,
  ): Promise<PostSummary[]> {
    const users: number[] = usersStr
      .split(',')
      .map((userStr) => Number(userStr));
    if (users.some((user) => isNaN(user))) {
      throw new BadRequestException();
    }
    const skip = Number(skipStr) || 0;
    return this.postQueryService.getUsersPosts({ users, skip });
  }
}
