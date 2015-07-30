CREATE TABLE alignments (
    id integer NOT NULL,
    name text,
    abbreviation text,
    description text
);

CREATE SEQUENCE alignments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE campaigns (
    id integer NOT NULL,
    userid integer NOT NULL,
    name text NOT NULL
);

CREATE SEQUENCE campaigns_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE creature_types (
    id integer NOT NULL,
    name text,
    description text
);

CREATE SEQUENCE creature_types_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


CREATE TABLE languages (
    id integer NOT NULL,
    name text,
    standard boolean
);

CREATE SEQUENCE languages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE monster_actions (
    id integer NOT NULL,
    monster_id integer,
    name text,
    description text
);

CREATE SEQUENCE monster_actions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE monsters (
    id integer NOT NULL,
    name text,
    hit_points integer,
    challenge_rating text,
    xp_reward integer,
    armor_class integer,
    armor_type text,
    size integer
);

CREATE SEQUENCE monsters_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE sizes (
    id integer NOT NULL,
    name text,
    space_squares text,
    space_hexes text
);

CREATE SEQUENCE sizes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE users (
    id integer NOT NULL,
    email text NOT NULL,
    username text NOT NULL,
    password text NOT NULL
);

CREATE SEQUENCE users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;