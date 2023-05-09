import { Module } from '@nestjs/common';
import { PostService } from './post/post.service';
import { OpenapiController } from './openapi/openapi.controller';
import { JwtModule } from '@nestjs/jwt';
import { PostController } from './post/post.controller';
import { OpenapiService } from './openapi/openapi.service';

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
