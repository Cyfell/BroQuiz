import { buildApp } from './app.js';
import { DEFAULT_HOST, DEFAULT_PORT } from '../shared/config.js';
import { fileURLToPath } from 'node:url';
import { dirname, resolve } from 'node:path';

const isProd = process.env.NODE_ENV === 'production';
const port = Number(process.env.PORT ?? DEFAULT_PORT);
const host = process.env.HOST ?? DEFAULT_HOST;

async function main() {
  const { app } = await buildApp();

  if (isProd) {
    const here = dirname(fileURLToPath(import.meta.url));
    const distRoot = resolve(here, '..', 'dist');
    const fastifyStatic = (await import('@fastify/static')).default;
    await app.register(fastifyStatic, {
      root: distRoot,
      wildcard: false,
    });
    app.setNotFoundHandler((req, reply) => {
      if (req.method === 'GET') return reply.sendFile('index.html');
      return reply.code(404).send({ error: 'not_found' });
    });
  }

  await app.listen({ host, port });
  app.log.info(`broquiz listening on http://${host}:${port}`);
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
