CREATE TABLE IF NOT EXISTS sleeps (
    "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    "user_id" uuid NOT NULL,
    "date" DATE NOT NULL,
    "sleep_time" TIME NOT NULL,
    "wakeup_time" TIME NOT NULL,
    "sleep_duration" INTERVAL NOT NULL,
    UNIQUE("user_id", "date")
);