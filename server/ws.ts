import type { FastifyInstance } from 'fastify';
import websocket from '@fastify/websocket';
import type { GameStore } from './state.js';
import type { ServerMessage } from '../shared/types.js';

export async function registerWebSocket(app: FastifyInstance, store: GameStore): Promise<void> {
  await app.register(websocket);

  app.get('/events', { websocket: true }, (socket) => {
    const send = (msg: ServerMessage) => {
      if (socket.readyState === socket.OPEN) {
        socket.send(JSON.stringify(msg));
      }
    };

    send({ type: 'state', state: store.snapshot() });

    const onChange = (state: ReturnType<GameStore['snapshot']>) =>
      send({ type: 'state', state });
    const onBuzz = (buzz: { teamId: number; at: number }) =>
      send({ type: 'buzz', buzz });

    store.on('change', onChange);
    store.on('buzz', onBuzz);

    socket.on('close', () => {
      store.off('change', onChange);
      store.off('buzz', onBuzz);
    });
  });
}
