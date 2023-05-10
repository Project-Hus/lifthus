import { Module } from '@nestjs/common';
import { PostService } from './post/post.service';
import { OpenapiController } from './openapi/openapi.controller';
import { JwtModule } from '@nestjs/jwt';
import { PostController } from './post/post.controller';
import { OpenapiService } from './openapi/openapi.service';
import { GetService } from './get/get.service';
import { GetController } from './get/get.controller';

@Module({
  imports: [
    JwtModule.register({
      secret: process.env.HUS_SECRET_KEY,
    }),
  ],
  controllers: [GetController, PostController, OpenapiController],
  providers: [GetService, PostService, OpenapiService],
})
export class AppModule {}
