import { MiddlewareConsumer, Module } from '@nestjs/common';
import { OpenapiController } from './openapi/openapi.controller';
import { JwtModule } from '@nestjs/jwt';
import { OpenapiService } from './openapi/openapi.service';
import { UidMiddleware } from './common/middlewares/uid.middleware';
import { QueryController } from './modules/query/query.controller';
import { PostQueryService } from './modules/query/post.query.service';
import { PostController } from './modules/post/post.controller';
import { PostService } from './modules/post/post.service';

@Module({
  imports: [
    JwtModule.register({
      secret: process.env.HUS_SECRET_KEY,
    }),
  ],
  controllers: [QueryController, PostController, OpenapiController],
  providers: [PostQueryService, PostService, OpenapiService],
})
export class AppModule {
  configure(consumer: MiddlewareConsumer) {
    consumer.apply(UidMiddleware).forRoutes('*');
  }
}
