-- Create a keyspace
CREATE KEYSPACE IF NOT EXISTS store WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : '1' };

-- Create a table
CREATE TABLE IF NOT EXISTS access_tokens (
access_token varchar PRIMARY KEY,
user_id bigint,
client_id bigint,
expires bigint
);