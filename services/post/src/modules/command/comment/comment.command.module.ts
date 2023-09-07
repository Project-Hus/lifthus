import { Module } from '@nestjs/common';

import { CommentController } from 'src/modules/command/comment/comment.controller';

import { PrismaService } from 'src/prisma/prisma.service';
import { CommentService } from 'src/modules/command/comment/comment.service';

import { RepositoryModule } from 'src/modules/repositories/repository.module';

@Module({
  imports: [RepositoryModule],
  controllers: [CommentController],
  providers: [CommentService, PrismaService],
})
export class CommentCommandModule {}
