-- A job that reschedules itself does so from inside its own handler, while its
-- row is still 'running'. With uniqueness spanning running rows too, that insert
-- hit ON CONFLICT DO NOTHING and the successor was silently dropped: polling ran
-- once per order and the replay drain ran only when a heartbeat recreated it.
--
-- Uniqueness over pending rows alone is what was meant: at most one *scheduled*
-- run per (kind, key), while a run in progress may line up its successor.
-- Claiming still selects only pending rows under FOR UPDATE SKIP LOCKED, so a
-- job can never be picked up twice.
DROP INDEX IF EXISTS jobs_key_idx;
CREATE UNIQUE INDEX jobs_key_idx ON jobs (kind, key) WHERE state = 'pending';
