import { INestApplication } from '@nestjs/common';
import { SwaggerModule, DocumentBuilder, OpenAPIObject } from '@nestjs/swagger';

/**
 * OpenAPI
 *
 * @param {INestApplication} app
 */
export function setupSwagger(app: INestApplication): OpenAPIObject {
  const options = new DocumentBuilder()
    .setTitle('Lifthus post service API')
    .setVersion('0.0.0')
    .setDescription('this document describes the API of the post service')
    .setTermsOfService('http://swagger.io/terms/')
    .setContact('lifthus', 'https://github.com/lifthus', 'lifthus531@gmail.com')
    .setLicense('MIT', '-')
    .build();

  const document = SwaggerModule.createDocument(app, options);
  // We use "swagger-ui-express" to provide the openapi document.
  // becausae nest SwaggerModule doesn't work in Lambda environment.
  // SwaggerModule.setup('/post/openapi', app, document);

  return document;
}

// currently not working in Lambda environment.
// I tried adding events for static files in template,
// and several other things, but it didn't work.
// I will try to fix it later.
