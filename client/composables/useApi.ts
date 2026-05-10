async function call<T>(path: string, init?: RequestInit): Promise<T> {
  const headers: Record<string, string> = { ...(init?.headers as Record<string, string> ?? {}) };
  if (init?.body !== undefined && init?.body !== null) {
    headers['content-type'] ??= 'application/json';
  }
  const res = await fetch(path, { ...init, headers });
  if (!res.ok) {
    const text = await res.text();
    throw new Error(`${res.status} ${text}`);
  }
  return res.json() as Promise<T>;
}

export const api = {
  createTeam: (id: number, name: string, color?: string) =>
    call('/teams', { method: 'POST', body: JSON.stringify({ id, name, color }) }),
  updateTeam: (id: number, patch: { name?: string; color?: string }) =>
    call(`/teams/${id}`, { method: 'PATCH', body: JSON.stringify(patch) }),
  deleteTeam: (id: number) => call(`/teams/${id}`, { method: 'DELETE' }),
  validate: () => call('/round/validate', { method: 'POST' }),
  discard: () => call('/round/discard', { method: 'POST' }),
  lock: () => call('/round/lock', { method: 'POST' }),
  unlock: () => call('/round/unlock', { method: 'POST' }),
  setCooldown: (cooldownMs: number) =>
    call('/round/cooldown', { method: 'POST', body: JSON.stringify({ cooldownMs }) }),
  score: (teamId: number, delta: number) =>
    call('/score', { method: 'POST', body: JSON.stringify({ teamId, delta }) }),
  reset: () => call('/reset', { method: 'POST' }),
  testBuzz: (id: number) =>
    call('/buzzer', { method: 'POST', body: JSON.stringify({ id }) }),
};
