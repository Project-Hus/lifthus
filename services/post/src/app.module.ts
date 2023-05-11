import { MiddlewareConsumer, Module } from '@nestjs/common';
import { OpenapiController } from './openapi/openapi.controller';
import { JwtModule } from '@nestjs/jwt';
import { OpenapiService } from './openapi/openapi.service';
import { UidMiddleware } from './common/middlewares/uid.middleware';
import { QueryController } from './modules/query/query.controller';
import { PostQueryService } from './modules/query/post.query.service';
import { PostController } from './modules/post/post.controller';
import { PostService } from './modules/post/post.service';
import { CommentController } from './modules/comment/comment.controller';
import { QueryModule } from './modules/query.module';
import { MutationModule } from './modules/mutation.module';

@Module({
  imports: [
    JwtModule.register({
      secret: process.env.HUS_SECRET_KEY,
    }),
    QueryModule,
    MutationModule,
  ],
  controllers: [OpenapiController],
  providers: [OpenapiService],
})
export class AppModule {
  configure(consumer: MiddlewareConsumer) {
    consumer.apply(UidMiddleware).forRoutes('*');
  }
}
