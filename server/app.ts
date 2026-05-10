import Fastify, { type FastifyInstance } from 'fastify';
import { GameStore } from './state.js';
import { registerWebSocket } from './ws.js';
import { buzzerRoutes } from './routes/buzzer.js';
import { teamsRoutes } from './routes/teams.js';
import { roundRoutes } from './routes/round.js';
import { scoreRoutes } from './routes/score.js';
import { stateRoutes } from './routes/state.js';
import { resetRoutes } from './routes/reset.js';

export async function buildApp(): Promise<{ app: FastifyInstance; store: GameStore }> {
  const app = Fastify({ logger: { level: process.env.LOG_LEVEL ?? 'info' } });
  const store = await GameStore.load();

  await registerWebSocket(app, store);

  await app.register(buzzerRoutes, { store });
  await app.register(teamsRoutes, { store });
  await app.register(roundRoutes, { store });
  await app.register(scoreRoutes, { store });
  await app.register(stateRoutes, { store });
  await app.register(resetRoutes, { store });

  return { app, store };
}
