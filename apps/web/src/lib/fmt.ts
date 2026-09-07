import type { Minor } from "./types";

/**
 * Formatting helpers.
 *
 * Minor-unit values are formatted by *string surgery*, never by arithmetic:
 * dividing by 100 in JavaScript would put a money value through a float, and
 * an int64 fen value can exceed Number.MAX_SAFE_INTEGER. We split the digit
 * string instead, so the rendering is exact for any magnitude.
 */

function splitSigned(v: Minor | null | undefined): { neg: boolean; digits: string } {
  let s = String(v ?? "").trim();
  if (s === "") return { neg: false, digits: "0" };
  const neg = s.startsWith("-");
  if (neg || s.startsWith("+")) s = s.slice(1);
  // Tolerate a value that already carries a decimal point by dropping it and
  // keeping only the digits the server sent; minor units have no fraction.
  const digits = s.replace(/[^0-9]/g, "");
  return { neg, digits: digits === "" ? "0" : digits };
}

function group(whole: string): string {
  return whole.replace(/\B(?=(\d{3})+(?!\d))/g, ",");
}

/** "1850" -> "18.50". Pass a symbol for "¥18.50". Exact, no arithmetic. */
export function minorText(v: Minor | null | undefined, symbol = ""): string {
  const { neg, digits } = splitSigned(v);
  const padded = digits.padStart(3, "0");
  const whole = padded.slice(0, -2).replace(/^0+(?=\d)/, "");
  const frac = padded.slice(-2);
  return `${neg ? "-" : ""}${symbol}${group(whole)}.${frac}`;
}

/** CNY fen -> "¥18.50". */
export function fenText(v: Minor | null | undefined): string {
  return minorText(v, "¥");
}

/** Parts per million -> "5.128400", a plain 6-decimal rate. */
export function ppmText(v: Minor | null | undefined): string {
  const { neg, digits } = splitSigned(v);
  const padded = digits.padStart(7, "0");
  const whole = padded.slice(0, -6).replace(/^0+(?=\d)/, "");
  return `${neg ? "-" : ""}${group(whole)}.${padded.slice(-6)}`;
}

/** Basis points -> "2.50%". */
export function bpsText(bps: number | null | undefined): string {
  const n = Number.isFinite(bps as number) ? (bps as number) : 0;
  return `${minorText(String(Math.trunc(n)))}%`;
}

/** A count that is safe to do maths on. Never use this on money or an id. */
export function toInt(v: unknown, fallback = 0): number {
  if (typeof v === "number") return Number.isFinite(v) ? Math.trunc(v) : fallback;
  if (typeof v === "string" && v.trim() !== "") {
    const n = Number(v);
    if (Number.isFinite(n)) return Math.trunc(n);
  }
  return fallback;
}

export function num(v: number | null | undefined): string {
  return group(String(toInt(v)));
}

function parseDate(s: string | null | undefined): Date | null {
  if (!s) return null;
  // Accept ISO 8601 and the naive "2006-01-02 15:04:05" the 1688 protocol uses.
  const iso = /^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$/.test(s) ? s.replace(" ", "T") : s;
  const d = new Date(iso);
  return Number.isNaN(d.getTime()) ? null : d;
}

export function dateTimeText(s: string | null | undefined): string {
  const d = parseDate(s);
  if (!d) return "—";
  return d.toLocaleString(undefined, {
    year: "numeric",
    month: "short",
    day: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
}

export function dateText(s: string | null | undefined): string {
  const d = parseDate(s);
  if (!d) return "—";
  return d.toLocaleDateString(undefined, { year: "numeric", month: "short", day: "numeric" });
}

export function timeText(s: string | null | undefined): string {
  const d = parseDate(s);
  if (!d) return "—";
  return d.toLocaleTimeString(undefined, { hour: "2-digit", minute: "2-digit", second: "2-digit" });
}

/** Grams -> "1.25 kg" / "820 g". Integer maths on a weight is fine. */
export function weightText(g: number | null | undefined): string {
  const n = toInt(g);
  if (n <= 0) return "—";
  if (n < 1000) return `${num(n)} g`;
  const kg = Math.trunc(n / 1000);
  const rem = Math.trunc((n % 1000) / 10);
  return `${kg}.${String(rem).padStart(2, "0")} kg`;
}

/** Millimetres -> centimetres with one decimal, for a "L x W x H" line. */
export function mmText(mm: number | null | undefined): string {
  const n = toInt(mm);
  if (n <= 0) return "—";
  return `${Math.trunc(n / 10)}.${n % 10} cm`;
}

export function dimensionsText(l?: number, w?: number, h?: number): string {
  if (!toInt(l) && !toInt(w) && !toInt(h)) return "—";
  return `${mmText(l)} × ${mmText(w)} × ${mmText(h)}`;
}

export function pluralise(n: number, one: string, many = `${one}s`): string {
  return `${num(n)} ${n === 1 ? one : many}`;
}

export function titleCase(s: string): string {
  if (!s) return "";
  return s
    .replace(/[_-]+/g, " ")
    .toLowerCase()
    .replace(/\b[a-z]/g, (c) => c.toUpperCase());
}

/** Split a camelCase or snake_case settings key into a readable label. */
export function labelise(key: string): string {
  const spaced = key
    .replace(/([a-z0-9])([A-Z])/g, "$1 $2")
    .replace(/[_-]+/g, " ")
    .trim();
  return spaced.charAt(0).toUpperCase() + spaced.slice(1);
}
