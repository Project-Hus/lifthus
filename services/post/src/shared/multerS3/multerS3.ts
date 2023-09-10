import aws from 'aws-sdk';
import multerS3 from 'multer-s3';
import crypto from 'crypto';
import { POST_IMAGE_MAX_SIZE } from 'src/shared/constraints';

if (process.env.NODE_ENV === 'native') {
  aws.config.loadFromPath('../../../config/s3.json');
} else {
  aws.config.update({
    accessKeyId: process.env.LIFTHUS_ACCESS_KEY_ID,
    secretAccessKey: process.env.LIFTHUS_SECRET_ACCESS_KEY,
    region: 'us-west-2',
  });
}

const s3 = new aws.S3();

export const getMulterS3Option = () => {
  return {
    storage: multerS3({
      s3: s3,
      bucket: 'lifthus-post-bucket',
      acl: 'public-read',
      contentType: multerS3.AUTO_CONTENT_TYPE,
      key: function (req, file, cb) {
        cb(
          null,
          `post/images/${Date.now()}_${crypto.randomBytes(4).toString('hex')}_${
            file.originalname
          }`,
        );
      },
    }),
    limits: {
      fileSize: POST_IMAGE_MAX_SIZE,
    },
    fileFilter: fileFilter,
  };
};

const fileFilter = (req, file, cb) => {
  const mtype = file.mimetype.split('/')[0];
  if (mtype === 'image') {
    cb(null, true);
  } else {
    cb(null, false);
    cb(new Error('only image'));
  }
};
