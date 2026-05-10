import type { FastifyPluginAsync } from 'fastify';
import type { GameStore } from '../state.js';
import type { BuzzerPressBody } from '../../shared/types.js';

export const buzzerRoutes: FastifyPluginAsync<{ store: GameStore }> = async (app, { store }) => {
  app.post<{ Body: BuzzerPressBody }>('/buzzer', {
    schema: {
      body: {
        type: 'object',
        required: ['id'],
        properties: { id: { type: 'integer' } },
      },
    },
    handler: async (req, reply) => {
      const { id } = req.body;
      const result = store.registerBuzz(id);
      if (!result.accepted) {
        return reply.code(409).send({ accepted: false, reason: result.reason });
      }
      return reply.send({ accepted: true });
    },
  });
};
