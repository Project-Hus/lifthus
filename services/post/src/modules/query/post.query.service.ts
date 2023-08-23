import { Injectable, Logger } from '@nestjs/common';
import { Post } from '@prisma/client';
import { PrismaService } from 'src/prisma/prisma.service';

@Injectable()
export class PostQueryService {
  constructor(private readonly prismaService: PrismaService) {}
  getHello(): string {
    return 'Hello World!';
  }

  /**
   * Get users' posts without comments.
   * @param users
   * @param skip
   * @returns
   */
  async getUsersPostsNoComments({
    users,
    skip,
  }: {
    users: number[];
    skip: number;
  }): Promise<Post[]> {
    try {
      const userPosts = await this.prismaService.post.findMany({
        where: {
          author: {
            in: users,
          },
        },
        include: {
          images: {
            select: {
              id: true,
              src: true,
            },
            orderBy: {
              order: 'asc',
            },
          },
          mentions: {
            select: {
              mentionee: true,
            },
          },
        },
        orderBy: {
          createdAt: 'desc',
        },
        take: 10,
        skip: skip,
      });

      return userPosts;
    } catch (error) {
      throw new Error('Failed to get user posts');
    }
  }

  /**
   * Gets post by slug.
   * @param slug
   * @returns
   */
  getPostBySlug(slug: string): Promise<Post> {
    return this.prismaService.post.findUnique({
      where: {
        slug: encodeURIComponent(slug),
      },
      include: {
        images: {
          select: {
            id: true,
            src: true,
          },
          orderBy: {
            order: 'asc',
          },
        },
        mentions: {
          select: {
            mentionee: true,
          },
        },
      },
    });
  }

  /**
   * Gets post by id.
   * @param id
   * @returns
   */
  getPostById(id: number): Promise<Post> {
    return this.prismaService.post.findUnique({
      where: {
        id: id,
      },
      include: {
        images: {
          select: {
            id: true,
            src: true,
          },
          orderBy: {
            order: 'asc',
          },
        },
        mentions: {
          select: {
            mentionee: true,
          },
        },
      },
    });
  }

  /**
   * Gets all posts.
   * @param skip
   * @returns
   */
  getAllPosts(skip: number): Promise<Post[]> {
    return this.prismaService.post.findMany({
      include: {
        images: {
          select: {
            id: true,
            src: true,
          },
          orderBy: {
            order: 'asc',
          },
        },
        comments: {
          select: {
            id: true,
            postId: true,
            author: true,
            createdAt: true,
            updatedAt: true,
            content: true,
            likenum: true,
            mentions: {
              select: {
                mentionee: true,
              },
            },
            replies: {
              select: {
                id: true,
                parentId: true,
                author: true,
                createdAt: true,
                updatedAt: true,
                content: true,
                likenum: true,
                mentions: {
                  select: {
                    mentionee: true,
                  },
                },
              },
            },
          },
        },
        mentions: {
          select: {
            mentionee: true,
          },
        },
      },
      orderBy: {
        createdAt: 'desc',
      },
      take: 10,
      skip: skip,
    });
  }

  /**
   * Gets users' ids and returns posts from the users
   * @param param0
   * @returns
   */
  async getUsersPosts({
    users,
    skip,
  }: {
    users: number[];
    skip: number;
  }): Promise<Post[]> {
    try {
      const userPosts = await this.prismaService.post.findMany({
        where: {
          author: {
            in: users,
          },
        },
        include: {
          images: {
            select: {
              id: true,
              src: true,
            },
            orderBy: {
              order: 'asc',
            },
          },
          comments: {
            select: {
              id: true,
              postId: true,
              author: true,
              createdAt: true,
              updatedAt: true,
              content: true,
              likenum: true,
              mentions: {
                select: {
                  mentionee: true,
                },
              },
              replies: {
                select: {
                  id: true,
                  parentId: true,
                  author: true,
                  createdAt: true,
                  updatedAt: true,
                  content: true,
                  likenum: true,
                  mentions: {
                    select: {
                      mentionee: true,
                    },
                  },
                },
              },
            },
          },
          mentions: {
            select: {
              mentionee: true,
            },
          },
        },
        orderBy: {
          createdAt: 'desc',
        },
        take: 10,
        skip: skip,
      });

      return userPosts;
    } catch (error) {
      throw new Error('Failed to get user posts');
    }
  }

  /**
   * Get posts by user id
   * @param uid
   * @param skip
   * @returns
   */
  getUserPosts(uid: number, skip: number): Promise<Post[]> {
    return this.prismaService.post.findMany({
      include: {
        images: {
          select: {
            id: true,
            src: true,
          },
          orderBy: {
            order: 'asc',
          },
        },
        comments: {
          select: {
            id: true,
            postId: true,
            author: true,
            createdAt: true,
            updatedAt: true,
            content: true,
            likenum: true,
            mentions: {
              select: {
                mentionee: true,
              },
            },
            replies: {
              select: {
                id: true,
                parentId: true,
                author: true,
                createdAt: true,
                updatedAt: true,
                content: true,
                likenum: true,
                mentions: {
                  select: {
                    mentionee: true,
                  },
                },
              },
            },
          },
        },
        mentions: {
          select: {
            mentionee: true,
          },
        },
      },
      where: { author: uid },
      orderBy: {
        createdAt: 'desc',
      },
      take: 10,
      skip: skip,
    });
  }
}
