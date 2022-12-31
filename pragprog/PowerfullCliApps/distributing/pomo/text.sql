DROP TABLE IF EXISTS "interval";

CREATE TABLE IF NOT EXISTS "interval" (
	"id" INTEGER,
	"start_time" DATETIME NOT NULL,
	"planned_duration" INTEGER DEFAULT 0,
	"actual_duration" INTEGER DEFAULT 0,
	"category" TEXT NOT NULL,
	"state" INTEGER DEFAULT 1,
	PRIMARY KEY("id")
);

INSERT INTO interval VALUES( NULL, date('now'), 25, 25, "Pomodoro", 3);
INSERT INTO interval VALUES( NULL, date('now'), 5, 5, "ShortBreak", 3);
INSERT INTO interval VALUES( NULL, date('now'), 15, 15, "LongBreak", 3);
