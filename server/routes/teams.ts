import type { FastifyPluginAsync } from 'fastify';
import type { GameStore } from '../state.js';
import type { CreateTeamBody, UpdateTeamBody } from '../../shared/types.js';

export const teamsRoutes: FastifyPluginAsync<{ store: GameStore }> = async (app, { store }) => {
  app.post<{ Body: CreateTeamBody }>('/teams', {
    schema: {
      body: {
        type: 'object',
        required: ['id', 'name'],
        properties: {
          id: { type: 'integer' },
          name: { type: 'string', minLength: 1 },
          color: { type: 'string' },
        },
      },
    },
    handler: async (req, reply) => {
      const { id, name, color } = req.body;
      try {
        store.addTeam({ id, name, color: color ?? store.pickColor() });
        return reply.code(201).send({ ok: true });
      } catch (err) {
        return reply.code(409).send({ error: (err as Error).message });
      }
    },
  });

  app.patch<{ Params: { id: string }; Body: UpdateTeamBody }>('/teams/:id', {
    handler: async (req, reply) => {
      const id = Number(req.params.id);
      try {
        store.updateTeam(id, req.body);
        return reply.send({ ok: true });
      } catch (err) {
        return reply.code(404).send({ error: (err as Error).message });
      }
    },
  });

  app.delete<{ Params: { id: string } }>('/teams/:id', {
    handler: async (req, reply) => {
      const id = Number(req.params.id);
      store.removeTeam(id);
      return reply.send({ ok: true });
    },
  });
};
