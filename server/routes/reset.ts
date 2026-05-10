import type { FastifyPluginAsync } from 'fastify';
import type { GameStore } from '../state.js';

export const resetRoutes: FastifyPluginAsync<{ store: GameStore }> = async (app, { store }) => {
  app.post('/reset', async (_req, reply) => {
    store.reset();
    return reply.send({ ok: true });
  });
};
