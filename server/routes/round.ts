import type { FastifyPluginAsync } from 'fastify';
import type { GameStore } from '../state.js';

export const roundRoutes: FastifyPluginAsync<{ store: GameStore }> = async (app, { store }) => {
  app.post('/round/validate', async (_req, reply) => {
    try {
      store.validateBuzz();
      return reply.send({ ok: true });
    } catch (err) {
      return reply.code(400).send({ error: (err as Error).message });
    }
  });

  app.post('/round/discard', async (_req, reply) => {
    try {
      store.discardBuzz();
      return reply.send({ ok: true });
    } catch (err) {
      return reply.code(400).send({ error: (err as Error).message });
    }
  });

  app.post('/round/lock', async (_req, reply) => {
    store.setLocked(true);
    return reply.send({ ok: true });
  });

  app.post('/round/unlock', async (_req, reply) => {
    store.setLocked(false);
    return reply.send({ ok: true });
  });

  app.post<{ Body: { cooldownMs: number } }>('/round/cooldown', {
    schema: {
      body: {
        type: 'object',
        required: ['cooldownMs'],
        properties: { cooldownMs: { type: 'integer', minimum: 0 } },
      },
    },
    handler: async (req, reply) => {
      store.setCooldownMs(req.body.cooldownMs);
      return reply.send({ ok: true });
    },
  });
};
