import { titleCase } from "../lib/fmt";

type Tone = "" | "Ok" | "Warn" | "Danger" | "Info";

/**
 * Order, parcel, job and supplier-order statuses all flow through here so one
 * vocabulary maps to one colour everywhere in the app.
 */
const TONES: Array<[RegExp, Tone]> = [
  [/^(delivered|completed|complete|success|succeeded|paid|received|confirmed|done|ok|active|published|running_ok)$/, "Ok"],
  [/(cancel|fail|error|refus|reject|expire|delist|dead|blocked)/, "Danger"],
  [/(pending|waiting|await|queued|new|created|unpaid|todo|hold|retry|partial)/, "Warn"],
  [/(shipped|shipping|transit|relay|relayed|processing|running|sync|import|paying|packed)/, "Info"],
];

export function toneFor(status: string | null | undefined): Tone {
  const s = (status ?? "").trim().toLowerCase();
  if (!s) return "";
  for (const [re, tone] of TONES) if (re.test(s)) return tone;
  return "";
}

export function StatusBadge({
  status,
  label,
  title,
}: {
  status: string | null | undefined;
  label?: string;
  title?: string;
}) {
  const tone = toneFor(status);
  const cls = tone ? `badge badge${tone}` : "badge";
  const text = label ?? (status ? titleCase(status) : "Unknown");
  return (
    <span className={cls} title={title ?? status ?? undefined}>
      {text}
    </span>
  );
}
