import { EventEmitter } from 'node:events';
import type { GameState, Team, Buzz } from '../shared/types.js';
import { DEFAULT_COOLDOWN_MS, SNAPSHOT_PATH, TEAM_COLORS } from '../shared/config.js';
import { loadSnapshot, writeSnapshot } from './snapshot.js';

export interface StateEvents {
  change: (state: GameState) => void;
  buzz: (buzz: Buzz) => void;
}

export class GameStore extends EventEmitter {
  private state: GameState;
  private snapshotQueue: Promise<void> = Promise.resolve();

  constructor(initial: GameState) {
    super();
    this.state = initial;
  }

  static async load(): Promise<GameStore> {
    const persisted = await loadSnapshot(SNAPSHOT_PATH);
    const initial: GameState = {
      teams: persisted?.teams ?? [],
      scores: persisted?.scores ?? {},
      currentBuzz: null,
      cooldowns: {},
      locked: persisted?.locked ?? false,
      config: persisted?.config ?? { cooldownMs: DEFAULT_COOLDOWN_MS },
    };
    return new GameStore(initial);
  }

  snapshot(): GameState {
    return structuredClone(this.state);
  }

  private mutate(fn: (s: GameState) => void, persist = true): void {
    fn(this.state);
    this.emit('change', this.snapshot());
    if (persist) this.persist();
  }

  private persist(): void {
    const snap = this.snapshot();
    this.snapshotQueue = this.snapshotQueue
      .then(() => writeSnapshot(SNAPSHOT_PATH, snap))
      .catch((err) => {
        console.error('[snapshot] write failed', err);
      });
  }

  addTeam(team: Team): void {
    this.mutate((s) => {
      if (s.teams.some((t) => t.id === team.id)) {
        throw new Error(`team id ${team.id} already exists`);
      }
      s.teams.push(team);
      if (s.scores[team.id] === undefined) s.scores[team.id] = 0;
    });
  }

  updateTeam(id: number, patch: Partial<Pick<Team, 'name' | 'color'>>): void {
    this.mutate((s) => {
      const t = s.teams.find((x) => x.id === id);
      if (!t) throw new Error(`team ${id} not found`);
      if (patch.name !== undefined) t.name = patch.name;
      if (patch.color !== undefined) t.color = patch.color;
    });
  }

  removeTeam(id: number): void {
    this.mutate((s) => {
      s.teams = s.teams.filter((t) => t.id !== id);
      delete s.scores[id];
      delete s.cooldowns[id];
      if (s.currentBuzz?.teamId === id) s.currentBuzz = null;
    });
  }

  registerBuzz(teamId: number): { accepted: boolean; reason?: string } {
    const s = this.state;
    if (!s.teams.some((t) => t.id === teamId)) {
      return { accepted: false, reason: 'unknown_team' };
    }
    if (s.locked) return { accepted: false, reason: 'locked' };
    if (s.currentBuzz) return { accepted: false, reason: 'already_buzzed' };
    const cd = s.cooldowns[teamId];
    if (cd && cd > Date.now()) return { accepted: false, reason: 'cooldown' };

    const buzz: Buzz = { teamId, at: Date.now() };
    this.mutate((st) => {
      st.currentBuzz = buzz;
    }, false);
    this.emit('buzz', buzz);
    return { accepted: true };
  }

  validateBuzz(): void {
    this.mutate((s) => {
      const cur = s.currentBuzz;
      if (!cur) throw new Error('no active buzz');
      s.scores[cur.teamId] = (s.scores[cur.teamId] ?? 0) + 1;
      s.currentBuzz = null;
      s.cooldowns = {};
    });
  }

  discardBuzz(): void {
    this.mutate((s) => {
      const cur = s.currentBuzz;
      if (!cur) throw new Error('no active buzz');
      s.cooldowns[cur.teamId] = Date.now() + s.config.cooldownMs;
      s.currentBuzz = null;
    });
  }

  adjustScore(teamId: number, delta: number): void {
    this.mutate((s) => {
      if (!s.teams.some((t) => t.id === teamId)) {
        throw new Error(`team ${teamId} not found`);
      }
      s.scores[teamId] = (s.scores[teamId] ?? 0) + delta;
    });
  }

  setLocked(locked: boolean): void {
    this.mutate((s) => {
      s.locked = locked;
    });
  }

  setCooldownMs(ms: number): void {
    this.mutate((s) => {
      s.config.cooldownMs = ms;
    });
  }

  reset(): void {
    this.mutate((s) => {
      s.teams = [];
      s.scores = {};
      s.currentBuzz = null;
      s.cooldowns = {};
      s.locked = false;
    });
  }

  pickColor(): string {
    const used = new Set(this.state.teams.map((t) => t.color));
    return TEAM_COLORS.find((c) => !used.has(c)) ?? TEAM_COLORS[0];
  }
}
