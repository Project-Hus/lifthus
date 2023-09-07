import { Module } from '@nestjs/common';

import { PostController } from 'src/modules/command/post/post.controller';

import { PrismaService } from 'src/prisma/prisma.service';

import { RepositoryModule } from 'src/modules/repositories/repository.module';
import { PostService } from 'src/modules/command/post/post.service';

@Module({
  imports: [RepositoryModule],
  controllers: [PostController],
  providers: [PostService, PrismaService],
})
export class PostCommandModule {}
