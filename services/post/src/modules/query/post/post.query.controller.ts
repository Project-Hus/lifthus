import {
  Controller,
  Get,
  Param,
  Query,
  BadRequestException,
  Inject,
} from '@nestjs/common';
import { PostQueryService } from 'src/modules/query/post/post.query.service';

import { PostSummaryDto } from 'src/dto/outbound/postSummary.dto';
import { PostDto } from 'src/dto/outbound/post.dto';

@Controller('/post/query/post')
export class PostQueryController {
  constructor(
    @Inject(PostQueryService)
    private readonly postQueryService: PostQueryService,
  ) {}

  @Get('/slug/:slug')
  getPostBySlug(@Param('slug') slug: string): Promise<PostDto> {
    return this.postQueryService.getPostBySlug(slug);
  }

  @Get('/id/:id')
  getPostById(@Param('id') idStr: string): Promise<PostDto> {
    return this.postQueryService.getPostById(idStr);
  }

  /**
   * gets all posts from the database and returns them.
   * @param skip
   * @returns Post[]
   * @example /post/query/post/all/0
   */
  @Get('/all')
  getAllPosts(@Query('skip') skipStr: string): Promise<PostSummaryDto[]> {
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
  ): Promise<PostSummaryDto[]> {
    const users: string[] = usersStr.split(',').map((userStr) => userStr);
    const skip = Number(skipStr) || 0;
    return this.postQueryService.getUsersPosts({ users, skip });
  }
}
