import { Controller, Get } from '@nestjs/common';
import { PostService } from './post.service';

@Controller()
export class PostController {
  constructor(private readonly appService: PostService) {}

  @Get('/post')
  getHello(): string {
    return this.appService.getHello();
  }
}
