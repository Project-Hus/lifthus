import { Module } from '@nestjs/common';

import { APP_FILTER } from '@nestjs/core';
import { HttpExceptionFilter } from 'src/common/filters/http-exception.filter';

import { PostCommandModule } from 'src/modules/command/post/post.command.module';
import { CommentCommandModule } from 'src/modules/command/comment/comment.command.module';

@Module({
  imports: [PostCommandModule, CommentCommandModule],
  controllers: [],
  providers: [{ provide: APP_FILTER, useClass: HttpExceptionFilter }],
})
export class CommandModule {}
