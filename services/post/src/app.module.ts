import { Module } from '@nestjs/common';
import { PostService } from './api/post/post.service';
import { PostController } from './api/post/post.controller';

@Module({
  imports: [],
  controllers: [PostController],
  providers: [PostService],
})
export class AppModule {}
