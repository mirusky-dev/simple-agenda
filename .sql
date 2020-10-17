CREATE SCHEMA IF NOT EXISTS calendar;

CREATE TABLE calendar.appointments (
	id bigserial NOT NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	deleted_at timestamp NULL,
	"name" text NOT NULL,
	"date" date NOT NULL,
	"hour" time NULL,
	"local" text NULL,
	CONSTRAINT appointments_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_appointments_deleted_at ON calendar.appointments USING btree (deleted_at);
