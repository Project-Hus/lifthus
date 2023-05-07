import { Module } from '@nestjs/common';
import { PostService } from './api/post/post.service';
import { PostController } from './api/post.controller';
import { OpenapiService } from './api/openapi/openapi.service';
import { OpenapiController } from './api/openapi/openapi.controller';

@Module({
  imports: [],
  controllers: [PostController, OpenapiController],
  providers: [PostService, OpenapiService],
})
export class AppModule {}
