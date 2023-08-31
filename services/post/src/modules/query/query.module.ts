import { Module } from '@nestjs/common';
import { QueryController } from './query.controller';
import { PostQueryService } from './post/post.query.service';
import { PrismaService } from 'src/prisma/prisma.service';
import { CommentQueryService } from './comment/comment.query.service';
import { APP_FILTER } from '@nestjs/core';
import { HttpExceptionFilter } from 'src/common/filters/http-exception.filter';
import { RepositoryModule } from '../repositories/repository.module';
import { DomainModule } from '../domain/domain.module';

@Module({
  controllers: [QueryController],
  imports: [DomainModule, RepositoryModule],
  providers: [
    PrismaService,
    PostQueryService,
    CommentQueryService,
    { provide: APP_FILTER, useClass: HttpExceptionFilter },
  ],
})
export class QueryModule {}
