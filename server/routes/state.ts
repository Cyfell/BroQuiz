import type { FastifyPluginAsync } from 'fastify';
import type { GameStore } from '../state.js';

export const stateRoutes: FastifyPluginAsync<{ store: GameStore }> = async (app, { store }) => {
  app.get('/state', async (_req, reply) => {
    return reply.send(store.snapshot());
  });
};
