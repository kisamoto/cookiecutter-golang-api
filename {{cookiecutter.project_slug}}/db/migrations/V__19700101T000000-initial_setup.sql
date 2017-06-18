--
-- Audit trail setup
--
-- create a schema named "audit"
CREATE schema audit;
REVOKE CREATE ON schema audit FROM public;
 
CREATE TABLE audit.logged_actions (
    schema_name text NOT NULL,
    TABLE_NAME text NOT NULL,
    user_name text,
    action_tstamp TIMESTAMP WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    action TEXT NOT NULL CHECK (action IN ('I','D','U')),
    original_data text,
    new_data text,
    query text
) WITH (fillfactor=100);
 
REVOKE ALL ON audit.logged_actions FROM public;
 
-- You may wish to use different permissions; this lets anybody
-- see the full audit data. In Pg 9.0 and above you can use column
-- permissions for fine-grained control.
GRANT SELECT ON audit.logged_actions TO public;
 
CREATE INDEX logged_actions_schema_table_idx 
ON audit.logged_actions(((schema_name||'.'||TABLE_NAME)::TEXT));
 
CREATE INDEX logged_actions_action_tstamp_idx 
ON audit.logged_actions(action_tstamp);
 
CREATE INDEX logged_actions_action_idx 
ON audit.logged_actions(action);

--
-- {{cookiecutter.project_name}} role based access control setup
-- 
CREATE SCHEMA auth;
REVOKE CREATE ON SCHEMA auth FROM public;

-- Organisations are sets of users with independent sets of roles
CREATE TABLE auth.organisation (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(30) NOT NULL,
    human_name  VARCHAR(50) NULL,
    description TEXT NULL
);

-- Roles are applied to an account to delegate permissions.
-- Each organisation has it's own set of roles
CREATE TABLE auth.role (
    id                  SERIAL PRIMARY KEY,
    organisation_id     BIGINT REFERENCES auth.organisation(id),
    name                VARCHAR(30) NOT NULL,
    human_name          VARCHAR(50) NULL,
    description         TEXT NULL
);

-- Permissions look after the ability to perform a certain task
CREATE TABLE auth.permission (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(30) NOT NULL,
    human_name  VARCHAR(50) NULL,
    description TEXT NULL
);

-- Mapping of roles to permissions.
-- Roles can have many permissions and permissions can 
-- belong to multiple roles
CREATE TABLE auth.role_permission (
    role_id BIGINT REFERENCES auth.role(id),
    permission_id BIGINT REFERENCES auth.permission(id)
);

-- Base representation of an account managed by Authboss
CREATE TABLE auth.account (
    id              SERIAL PRIMARY KEY,
    email           VARCHAR,
    password        BYTEA, -- bcrypt hashed password
    confirm_token   VARCHAR(100),
    confirmed       TIMESTAMP WITH TIME ZONE
);

CREATE TABLE auth.account_password_recovery (
    account_id              SERIAL REFERENCES account(id),
    recovery_token          VARCHAR(100),
    token_expiry_datetime   TIMESTAMP WITH TIME ZONE
)

-- Mapping of accounts to roles.
-- Accounts can have multiple roles.
-- Roles can be applied to multiple accounts
CREATE TABLE auth.account_role (
    account_id          BIGINT REFERENCES auth.account(id),
    role_id             BIGINT REFERENCES auth.role(id),
    granted_datetime    TIMESTAMP WITH TIME ZONE
);

-- For security and audit purposes
-- log when user logs in and from what IP address
CREATE TABLE auth.login_log (
    account_id      BIGINT REFERENCES auth.account(id),
    login_datetime  TIMESTAMP WITH TIME ZONE,
    access_ip       INET
)

--
-- {{cookiecutter.project_name}} table setup
-- 

--
-- {{cookiecutter.project_name}} data insertion
-- 
