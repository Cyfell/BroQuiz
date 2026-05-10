import type { FastifyPluginAsync } from 'fastify';
import type { GameStore } from '../state.js';
import type { ScoreBody } from '../../shared/types.js';

export const scoreRoutes: FastifyPluginAsync<{ store: GameStore }> = async (app, { store }) => {
  app.post<{ Body: ScoreBody }>('/score', {
    schema: {
      body: {
        type: 'object',
        required: ['teamId', 'delta'],
        properties: {
          teamId: { type: 'integer' },
          delta: { type: 'integer' },
        },
      },
    },
    handler: async (req, reply) => {
      try {
        store.adjustScore(req.body.teamId, req.body.delta);
        return reply.send({ ok: true });
      } catch (err) {
        return reply.code(404).send({ error: (err as Error).message });
      }
    },
  });
};
