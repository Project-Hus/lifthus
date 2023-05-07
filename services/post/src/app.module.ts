import { Module } from '@nestjs/common';
import { PostService, OpenapiService } from './api/post/post.service';
import { PostController } from './api/post/post.controller';

@Module({
  imports: [],
  controllers: [PostController],
  providers: [PostService, OpenapiService],
})
export class AppModule {}
