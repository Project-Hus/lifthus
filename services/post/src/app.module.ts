import { MiddlewareConsumer, Module } from '@nestjs/common';
import { PostService } from './post/post.service';
import { OpenapiController } from './openapi/openapi.controller';
import { JwtModule } from '@nestjs/jwt';
import { PostController } from './post/post.controller';
import { OpenapiService } from './openapi/openapi.service';
import { GetService } from './get/get.service';
import { GetController } from './get/get.controller';
import { UidMiddleware } from './common/middlewares/uid.middleware';

@Module({
  imports: [
    JwtModule.register({
      secret: process.env.HUS_SECRET_KEY,
    }),
  ],
  controllers: [GetController, PostController, OpenapiController],
  providers: [GetService, PostService, OpenapiService],
})
export class AppModule {
  configure(consumer: MiddlewareConsumer) {
    consumer.apply(UidMiddleware).forRoutes('*');
  }
}
