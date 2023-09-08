import { MiddlewareConsumer, Module } from '@nestjs/common';
import { OpenapiController } from 'src/openapi/openapi.controller';
import { JwtModule } from '@nestjs/jwt';
import { OpenapiService } from 'src/openapi/openapi.service';
import { UidMiddleware } from 'src/shared/middlewares/uid.middleware';
import { QueryModule } from 'src/modules/query/query.module';
import { CommandModule } from 'src/modules/command/command.module';

@Module({
  imports: [
    JwtModule.register({
      secret: process.env.HUS_SECRET_KEY,
    }),
    QueryModule,
    CommandModule,
  ],
  controllers: [OpenapiController],
  providers: [OpenapiService],
})
export class AppModule {
  configure(consumer: MiddlewareConsumer) {
    consumer.apply(UidMiddleware).forRoutes('*');
  }
}
