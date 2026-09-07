import type { Issue } from "./types";

/**
 * The whole data layer. One fetch wrapper, no client library.
 *
 * Dev runs behind the Vite proxy, production is served from the Go binary, so
 * "/api/..." is same-origin in both and needs no base URL configuration.
 */

const ADMIN_TOKEN_KEY = "marketplace.adminToken";

export class ApiError extends Error {
  readonly status: number;
  readonly body: unknown;
  readonly issues: Issue[];

  constructor(status: number, message: string, body: unknown, issues: Issue[] = []) {
    super(message);
    this.name = "ApiError";
    this.status = status;
    this.body = body;
    this.issues = issues;
  }
}

// ------------------------------------------------------------- admin token

let adminTokenCache: string | null = null;

export function getAdminToken(): string {
  if (adminTokenCache !== null) return adminTokenCache;
  try {
    adminTokenCache = sessionStorage.getItem(ADMIN_TOKEN_KEY) ?? "";
  } catch {
    adminTokenCache = "";
  }
  return adminTokenCache;
}

export function setAdminToken(token: string): void {
  adminTokenCache = token;
  try {
    if (token) sessionStorage.setItem(ADMIN_TOKEN_KEY, token);
    else sessionStorage.removeItem(ADMIN_TOKEN_KEY);
  } catch {
    /* private mode: the in-memory cache still serves this tab */
  }
}

// ------------------------------------------------------------------ fetch

function issuesOf(body: unknown): Issue[] {
  if (body && typeof body === "object" && Array.isArray((body as { issues?: unknown }).issues)) {
    return (body as { issues: Issue[] }).issues;
  }
  return [];
}

function messageOf(status: number, body: unknown): string {
  if (body && typeof body === "object") {
    const b = body as Record<string, unknown>;
    for (const key of ["message", "error", "errorMessage", "errorMsg", "erroMsg"]) {
      const v = b[key];
      if (typeof v === "string" && v) return v;
    }
    const issues = issuesOf(body);
    if (issues.length) return issues.map((i) => i.message).join("; ");
  }
  if (typeof body === "string" && body.trim()) return body.trim().slice(0, 300);
  if (status === 0) return "Cannot reach the server.";
  if (status === 401 || status === 403) return "Not authorised.";
  if (status === 404) return "Not found.";
  return `Request failed (${status}).`;
}

export type FetchOptions = {
  method?: string;
  body?: unknown;
  signal?: AbortSignal;
};

export async function apiFetch<T>(path: string, opts: FetchOptions = {}): Promise<T> {
  const headers: Record<string, string> = { Accept: "application/json" };
  if (opts.body !== undefined) headers["Content-Type"] = "application/json";
  if (path.startsWith("/api/admin")) {
    const token = getAdminToken();
    if (token) headers["Authorization"] = `Bearer ${token}`;
  }

  let res: Response;
  try {
    res = await fetch(path, {
      method: opts.method ?? (opts.body === undefined ? "GET" : "POST"),
      headers,
      body: opts.body === undefined ? undefined : JSON.stringify(opts.body),
      signal: opts.signal,
      credentials: "same-origin",
    });
  } catch (err) {
    if ((err as { name?: string })?.name === "AbortError") throw err;
    throw new ApiError(0, "Cannot reach the server.", null);
  }

  const text = await res.text();
  let parsed: unknown = null;
  if (text) {
    try {
      parsed = JSON.parse(text);
    } catch {
      parsed = text;
    }
  }

  if (!res.ok) {
    throw new ApiError(res.status, messageOf(res.status, parsed), parsed, issuesOf(parsed));
  }
  return parsed as T;
}

export const api = {
  get: <T>(path: string, signal?: AbortSignal) => apiFetch<T>(path, { signal }),
  post: <T>(path: string, body?: unknown) => apiFetch<T>(path, { method: "POST", body: body ?? {} }),
  put: <T>(path: string, body?: unknown) => apiFetch<T>(path, { method: "PUT", body: body ?? {} }),
  patch: <T>(path: string, body?: unknown) => apiFetch<T>(path, { method: "PATCH", body: body ?? {} }),
  del: <T>(path: string) => apiFetch<T>(path, { method: "DELETE" }),
};

export function errorText(err: unknown): string {
  if (err instanceof ApiError) return err.message;
  if (err instanceof Error) return err.message;
  if (typeof err === "string") return err;
  return "Something went wrong.";
}

// ------------------------------------------------------------- URL helpers

/** Build "/api/products?q=..." skipping empty values. Ids stay strings. */
export function withQuery(path: string, params: Record<string, string | number | undefined | null>): string {
  const qs = new URLSearchParams();
  for (const [k, v] of Object.entries(params)) {
    if (v === undefined || v === null) continue;
    const s = String(v);
    if (s === "") continue;
    qs.set(k, s);
  }
  const q = qs.toString();
  return q ? `${path}?${q}` : path;
}
