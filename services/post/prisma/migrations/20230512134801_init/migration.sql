/*
  Warnings:

  - You are about to drop the column `postId` on the `CommentLike` table. All the data in the column will be lost.
  - Added the required column `commentId` to the `CommentLike` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE `CommentLike` DROP FOREIGN KEY `CommentLike_postId_fkey`;

-- AlterTable
ALTER TABLE `CommentLike` DROP COLUMN `postId`,
    ADD COLUMN `commentId` BIGINT UNSIGNED NOT NULL;

-- AddForeignKey
ALTER TABLE `CommentLike` ADD CONSTRAINT `CommentLike_commentId_fkey` FOREIGN KEY (`commentId`) REFERENCES `Comment`(`id`) ON DELETE RESTRICT ON UPDATE CASCADE;
