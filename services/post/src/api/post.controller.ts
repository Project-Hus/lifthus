import { Controller, Get } from '@nestjs/common';
import { PostService } from './post/post.service';

@Controller('/post/post')
export class PostController {
  constructor(private readonly appService: PostService) {}
}
