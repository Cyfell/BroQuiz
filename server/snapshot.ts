import { mkdir, readFile, writeFile, rename } from 'node:fs/promises';
import { dirname, resolve } from 'node:path';
import type { GameState } from '../shared/types.js';

interface PersistedState {
  teams: GameState['teams'];
  scores: GameState['scores'];
  locked: GameState['locked'];
  config: GameState['config'];
}

export async function loadSnapshot(path: string): Promise<PersistedState | null> {
  try {
    const raw = await readFile(resolve(path), 'utf8');
    return JSON.parse(raw) as PersistedState;
  } catch (err: unknown) {
    if ((err as NodeJS.ErrnoException).code === 'ENOENT') return null;
    throw err;
  }
}

export async function writeSnapshot(path: string, state: GameState): Promise<void> {
  const abs = resolve(path);
  await mkdir(dirname(abs), { recursive: true });
  const persisted: PersistedState = {
    teams: state.teams,
    scores: state.scores,
    locked: state.locked,
    config: state.config,
  };
  const tmp = `${abs}.tmp`;
  await writeFile(tmp, JSON.stringify(persisted, null, 2), 'utf8');
  await rename(tmp, abs);
}
