import { Injectable } from '@nestjs/common';
import aws from 'aws-sdk';

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

@Injectable()
export class S3Service {
  async uploadImages(files: any) {
    return 'SUCCESS';
  }
}
