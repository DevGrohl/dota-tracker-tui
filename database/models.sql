CREATE TABLE IF NOT EXISTS "matches" (
  "match_id" integer NOT NULL,
  "start_time" timestamp with time zone NOT NULL,
  "lobby_type" integer NOT NULL,
  CONSTRAINT pk_matches PRIMARY KEY ("match_id")
);

