/**
 * @param author number
 * @param content string
 */
export interface CommentDto {
  postId: PrimaryKey;
  parentId?: PrimaryKey;
  author: PrimaryKey;
  content: string;
}
