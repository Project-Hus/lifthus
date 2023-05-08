import { Module } from '@nestjs/common';
import { PostService } from './api/post.service';
import { PostController } from './api/post.controller';
import { OpenapiService } from './api/openapi/openapi.service';
import { OpenapiController } from './api/openapi/openapi.controller';
import { JwtModule } from '@nestjs/jwt';

@Module({
  imports: [
    JwtModule.register({
      secret: process.env.HUS_SECRET_KEY,
    }),
  ],
  controllers: [PostController, OpenapiController],
  providers: [PostService, OpenapiService],
})
export class AppModule {}
