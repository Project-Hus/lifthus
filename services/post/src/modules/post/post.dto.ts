/**
 * @param author number
 * @param content string
 */
export interface PostDto {
  userGroup?: PrimaryKey;
  author: PrimaryKey;
  content: string;
}
