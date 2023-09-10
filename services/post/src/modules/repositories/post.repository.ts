import { Inject, Injectable } from '@nestjs/common';

import { PostRepository } from 'src/modules/repositories/abstract/post.repository';

import { PrismaService } from 'src/modules/repositories/prisma/prisma.service';

import { User } from 'src/domain/aggregates/user/user.model';
import { Post } from 'src/domain/aggregates/post/post.model';
import { PostSummary } from 'src/domain/aggregates/post/postSummary.model';

import { Prisma } from '@prisma/client';
import {
  PostContents,
  PostIds,
  PostUpdates,
} from 'src/domain/aggregates/post/post.vo';
import { Timestamps } from 'src/domain/vo';

type PostSummWithImage = {
  id: bigint;
  userGroup: bigint | null;
  author: bigint;
  createdAt: Date;
  updatedAt: Date;
  slug: string;
  images: PrismaPostImage[];
};

type PostWithImages = {
  id: bigint;
  userGroup: bigint | null;
  author: bigint;
  createdAt: Date;
  updatedAt: Date;
  slug: string;
  content: string;
  images: PrismaPostImage[];
};

type PrismaPostImage = {
  id: bigint;
  src: string;
  order: number;
};

@Injectable()
export class PrismaPostRepository extends PostRepository {
  constructor(
    @Inject(PrismaService) private readonly prismaService: PrismaService,
  ) {
    super();
  }

  async _getAllPostSumms(skip: number): Promise<PostSummary[]> {
    try {
      const allPosts: PostSummWithImage[] =
        await this.prismaService.post.findMany({
          include: {
            images: {
              select: {
                order: true,
                id: true,
                src: true,
              },
              orderBy: {
                order: 'asc',
              },
              take: 1,
            },
          },
          orderBy: {
            createdAt: 'desc',
          },
          take: 10,
          skip: skip,
        });

      return allPosts.map((post: PostSummWithImage) => {
        return new PostSummary(
          post.author,
          post.id,
          post.slug,
          post.images.map((image) => image.src),
          post.createdAt,
          post.updatedAt,
        );
      });
    } catch (e) {
      return Promise.reject(e);
    }
  }

  async _getUsersPostSumms(
    users: User[],
    skip: number,
  ): Promise<PostSummary[]> {
    try {
      const usersPosts = await this.prismaService.post.findMany({
        where: {
          author: {
            in: users.map((user) => user.getID()),
          },
        },
        include: {
          images: {
            select: {
              order: true,
              id: true,
              src: true,
            },
            orderBy: {
              order: 'asc',
            },
            take: 1,
          },
        },
        orderBy: {
          createdAt: 'desc',
        },
        take: 10,
        skip: skip,
      });
      return usersPosts.map((post: PostSummWithImage) => {
        return new PostSummary(
          post.author,
          post.id,
          post.slug,
          post.images.map((image) => image.src),
          post.createdAt,
          post.updatedAt,
        );
      });
    } catch (e) {
      return Promise.reject(e);
    }
  }

  async _getPostByID(pid: bigint): Promise<Post | null> {
    try {
      const post: PostWithImages = await this.prismaService.post.findUnique({
        where: {
          id: pid,
        },
        include: {
          images: {
            select: {
              order: true,
              id: true,
              src: true,
            },
            orderBy: {
              order: 'asc',
            },
          },
        },
      });
      if (!post) return null;
      const ids = new PostIds(post.id, post.slug);
      const contents = new PostContents(
        post.images.map((image) => image.src),
        post.content,
      );
      const timestamps = new Timestamps(post.createdAt, post.updatedAt);
      return Post.from(post.author, ids, contents, timestamps);
    } catch (e) {
      return Promise.reject(e);
    }
  }

  async _getPostBySlug(slug: string): Promise<Post | null> {
    try {
      const post: PostWithImages = await this.prismaService.post.findUnique({
        where: {
          slug: slug,
        },
        include: {
          images: {
            select: {
              order: true,
              id: true,
              src: true,
            },
            orderBy: {
              order: 'asc',
            },
          },
        },
      });
      if (!post) return null;
      const ids = new PostIds(post.id, post.slug);
      const contents = new PostContents(
        post.images.map((image) => image.src),
        post.content,
      );
      const timestamps = new Timestamps(post.createdAt, post.updatedAt);
      return Post.from(post.author, ids, contents, timestamps);
    } catch (e) {
      return Promise.reject(e);
    }
  }

  async _createPost(post: Post): Promise<Post> {
    if (post.isPersisted()) return undefined;
    try {
      let createInput: Prisma.PostCreateInput = {
        author: post.getAuthor(),
        slug: post.getSlug(),
        content: post.getContent(),
        images: {
          create: post.getImageSrcs().map((src, idx) => {
            return {
              src,
              order: idx,
            };
          }),
        },
      };

      const newPost = await this.prismaService.post.create({
        data: createInput,
        include: {
          images: {
            select: {
              order: true,
              id: true,
              src: true,
            },
            orderBy: {
              order: 'asc',
            },
          },
        },
      });

      const ids = new PostIds(newPost.id, newPost.slug);
      const contents = new PostContents(
        newPost.images.map((image) => image.src),
        newPost.content,
      );
      const timestamps = new Timestamps(newPost.createdAt, newPost.updatedAt);
      return Post.from(newPost.author, ids, contents, timestamps);
    } catch (e) {
      return Promise.reject(e);
    }
  }

  async _deletePost(target: Post): Promise<Post> {
    try {
      const deletedPost = await this.prismaService.post.delete({
        where: {
          id: target.getID(),
        },
      });
      const ids = new PostIds(deletedPost.id, deletedPost.slug);
      const contents = new PostContents([], deletedPost.content);
      const timestamps = new Timestamps(
        deletedPost.createdAt,
        deletedPost.updatedAt,
      );
      return Post.from(deletedPost.author, ids, contents, timestamps);
    } catch (e) {
      return Promise.reject(e);
    }
  }

  async _save(pid: bigint, updates: PostUpdates): Promise<Post> {
    try {
      const updatedPost = await this.prismaService.post.update({
        where: {
          id: pid,
        },
        include: {
          images: {
            select: {
              order: true,
              id: true,
              src: true,
            },
            orderBy: {
              order: 'asc',
            },
          },
        },
        data: updates,
      });

      const ids = new PostIds(updatedPost.id, updatedPost.slug);
      const contents = new PostContents(
        updatedPost.images.map((image) => image.src),
        updatedPost.content,
      );
      const timestamps = new Timestamps(
        updatedPost.createdAt,
        updatedPost.updatedAt,
      );
      return Post.from(updatedPost.author, ids, contents, timestamps);
    } catch (e) {
      return Promise.reject(e);
    }
  }
}
