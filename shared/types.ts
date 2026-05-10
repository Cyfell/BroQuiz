export interface Team {
  id: number;
  name: string;
  color: string;
}

export interface Buzz {
  teamId: number;
  at: number;
}

export interface GameState {
  teams: Team[];
  scores: Record<number, number>;
  currentBuzz: Buzz | null;
  cooldowns: Record<number, number>;
  locked: boolean;
  config: {
    cooldownMs: number;
  };
}

export type ServerMessage =
  | { type: 'state'; state: GameState }
  | { type: 'buzz'; buzz: Buzz };

export interface BuzzerPressBody {
  id: number;
}

export interface CreateTeamBody {
  id: number;
  name: string;
  color?: string;
}

export interface UpdateTeamBody {
  name?: string;
  color?: string;
}

export interface ScoreBody {
  teamId: number;
  delta: number;
}

export interface AwardBody {
  teamId: number;
}
