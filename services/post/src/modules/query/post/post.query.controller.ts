import { Controller, Get, Param, Query, Inject } from '@nestjs/common';
import { PostQueryService } from 'src/modules/query/post/post.query.service';

import { PostDto } from 'src/dto/outbound/post.dto';
import { Uid } from 'src/shared/decorators/authParam.decorator';

@Controller('/post/query/post')
export class PostQueryController {
  constructor(
    @Inject(PostQueryService)
    private readonly postQueryService: PostQueryService,
  ) {}

  @Get('/slug/:slug')
  getPostBySlug(
    @Param('slug') slug: string,
    @Uid() client: BigInt | undefined,
  ): Promise<PostDto> {
    return this.postQueryService.getPostBySlug(slug, client);
  }

  @Get('/id/:id')
  getPostById(
    @Param('id') idStr: string,
    @Uid() client: BigInt | undefined,
  ): Promise<PostDto> {
    return this.postQueryService.getPostById(idStr, client);
  }

  /**
   * gets all posts from the database and returns them.
   * @param skip
   * @returns Post[]
   * @example /post/query/post/all/0
   */
  @Get('/all')
  getAllPosts(
    @Query('skip') skipStr: string,
    @Uid() client: BigInt | undefined,
  ): Promise<PostDto[]> {
    const skip = Number(skipStr) || 0;
    return this.postQueryService.getAllPosts(skip, client);
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
    @Uid() client: BigInt | undefined,
  ): Promise<PostDto[]> {
    const users: string[] = usersStr.split(',').map((userStr) => userStr);
    const skip = Number(skipStr) || 0;
    return this.postQueryService.getUsersPosts({ users, skip, client });
  }
}
