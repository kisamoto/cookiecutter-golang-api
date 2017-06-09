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
-- {{cookiecutter.project_name}} table setup
-- 
CREATE TABLE public.account_credentials (
    
);
