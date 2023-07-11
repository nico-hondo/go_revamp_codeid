--
-- Name: curriculum; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA curriculum;

ALTER SCHEMA curriculum OWNER TO postgres;

--
-- Name: progentityidseq(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.progentityidseq() RETURNS character varying
    LANGUAGE sql
    AS $$
	select lpad(''||nextval('seq_prog_entity_id'),4,'0')
end;$$;

ALTER FUNCTION public.progentityidseq() OWNER TO postgres;

--
-- Name: secdidseq(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.secdidseq() RETURNS character varying
    LANGUAGE sql
    AS $$
	select lpad(''||nextval('seq_secd_id'),4,'0')
end;$$;

ALTER FUNCTION public.secdidseq() OWNER TO postgres;

--
-- Name: sectidseq(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.sectidseq() RETURNS character varying
    LANGUAGE sql
    AS $$
	select lpad(''||nextval('seq_sect_id'),4,'0')
end;$$;


ALTER FUNCTION public.sectidseq() OWNER TO postgres;

--
-- Name: sedmidseq(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.sedmidseq() RETURNS character varying
    LANGUAGE sql
    AS $$
	select lpad(''||nextval('seq_sedm_id'),4,'0')
end;$$;

ALTER FUNCTION public.sedmidseq() OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: program_entity; Type: TABLE; Schema: curriculum; Owner: postgres
--

CREATE TABLE curriculum.program_entity (
    prog_entity_id integer NOT NULL,
    prog_title character varying(256),
    prog_headline character varying(512),
    prog_type character varying(15),
    prog_learning_type character varying(15),
    prog_rating integer,
    prog_total_trainee integer,
    prog_image character varying(256),
    prog_best_seller character(1),
    prog_price integer,
    prog_language character varying(35),
    prog_modified_date timestamp without time zone,
    prog_duration integer,
    prog_duration_type character varying(15),
    prog_tag_skill character varying(512),
    prog_city_id integer,
    prog_cate_id integer,
    prog_created_by integer,
    prog_status character varying(15),
    CONSTRAINT program_entity_prog_best_seller_check CHECK ((prog_best_seller = ANY (ARRAY['0'::bpchar, '1'::bpchar]))),
    CONSTRAINT program_entity_prog_duration_type_check CHECK (((prog_duration_type)::text = ANY ((ARRAY['month'::character varying, 'week'::character varying, 'days'::character varying])::text[]))),
    CONSTRAINT program_entity_prog_learning_type_check CHECK (((prog_learning_type)::text = ANY ((ARRAY['online'::character varying, 'offline'::character varying, 'both'::character varying])::text[]))),
    CONSTRAINT program_entity_prog_status_check CHECK (((prog_status)::text = ANY ((ARRAY['draft'::character varying, 'publish'::character varying])::text[]))),
    CONSTRAINT program_entity_prog_type_check CHECK (((prog_type)::text = ANY ((ARRAY['bootcamp'::character varying, 'course'::character varying])::text[])))
);


ALTER TABLE curriculum.program_entity OWNER TO postgres;

--
-- Name: program_entity_description; Type: TABLE; Schema: curriculum; Owner: postgres
--

CREATE TABLE curriculum.program_entity_description (
    pred_prog_entity_id integer NOT NULL,
    pred_item_learning json,
    pred_item_include json,
    pred_requirement json,
    pred_description json,
    pred_target_level json
);


ALTER TABLE curriculum.program_entity_description OWNER TO postgres;

--
-- Name: program_entity_prog_entity_id_seq; Type: SEQUENCE; Schema: curriculum; Owner: postgres
--

CREATE SEQUENCE curriculum.program_entity_prog_entity_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE curriculum.program_entity_prog_entity_id_seq OWNER TO postgres;

--
-- Name: program_entity_prog_entity_id_seq; Type: SEQUENCE OWNED BY; Schema: curriculum; Owner: postgres
--

ALTER SEQUENCE curriculum.program_entity_prog_entity_id_seq OWNED BY curriculum.program_entity.prog_entity_id;

--
-- Name: program_reviews; Type: TABLE; Schema: curriculum; Owner: postgres
--

CREATE TABLE curriculum.program_reviews (
    prow_user_entity_id integer NOT NULL,
    prow_prog_entity_id integer NOT NULL,
    prow_review character varying(512),
    prow_rating integer,
    prow_modified_date timestamp without time zone
);


ALTER TABLE curriculum.program_reviews OWNER TO postgres;

--
-- Name: section_detail; Type: TABLE; Schema: curriculum; Owner: postgres
--

CREATE TABLE curriculum.section_detail (
    secd_id integer NOT NULL,
    secd_title character varying(256),
    secd_preview character(1),
    secd_score integer,
    secd_note character varying(256),
    secd_minute integer,
    secd_modified_date timestamp without time zone,
    secd_sect_id integer,
    CONSTRAINT section_detail_secd_preview_check CHECK ((secd_preview = ANY (ARRAY['0'::bpchar, '1'::bpchar])))
);


ALTER TABLE curriculum.section_detail OWNER TO postgres;

--
-- Name: section_detail_material; Type: TABLE; Schema: curriculum; Owner: postgres
--

CREATE TABLE curriculum.section_detail_material (
    sedm_id integer NOT NULL,
    sedm_filename character varying(55),
    sedm_filesize integer,
    sedm_filetype character varying(15),
    sedm_filelink character varying(255),
    sedm_modified_date timestamp without time zone,
    sedm_secd_id integer
);

ALTER TABLE curriculum.section_detail_material OWNER TO postgres;

--
-- Name: section_detail_material_sedm_id_seq; Type: SEQUENCE; Schema: curriculum; Owner: postgres
--

CREATE SEQUENCE curriculum.section_detail_material_sedm_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE curriculum.section_detail_material_sedm_id_seq OWNER TO postgres;

--
-- Name: section_detail_material_sedm_id_seq; Type: SEQUENCE OWNED BY; Schema: curriculum; Owner: postgres
--

ALTER SEQUENCE curriculum.section_detail_material_sedm_id_seq OWNED BY curriculum.section_detail_material.sedm_id;

--
-- Name: section_detail_secd_id_seq; Type: SEQUENCE; Schema: curriculum; Owner: postgres
--

CREATE SEQUENCE curriculum.section_detail_secd_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE curriculum.section_detail_secd_id_seq OWNER TO postgres;

--
-- Name: section_detail_secd_id_seq; Type: SEQUENCE OWNED BY; Schema: curriculum; Owner: postgres
--

ALTER SEQUENCE curriculum.section_detail_secd_id_seq OWNED BY curriculum.section_detail.secd_id;
--
-- Name: sections; Type: TABLE; Schema: curriculum; Owner: postgres
--

CREATE TABLE curriculum.sections (
    sect_title character varying(100),
    sect_description character varying(256),
    sect_total_section integer,
    sect_total_lecture integer,
    sect_total_minute integer,
    sect_modified_date timestamp without time zone,
    sect_prog_entity_id integer NOT NULL,
    sect_id integer NOT NULL
);


ALTER TABLE curriculum.sections OWNER TO postgres;

--
-- Name: sections_sect_id_seq; Type: SEQUENCE; Schema: curriculum; Owner: postgres
--

CREATE SEQUENCE curriculum.sections_sect_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE curriculum.sections_sect_id_seq OWNER TO postgres;

--
-- Name: sections_sect_id_seq; Type: SEQUENCE OWNED BY; Schema: curriculum; Owner: postgres
--

ALTER SEQUENCE curriculum.sections_sect_id_seq OWNED BY curriculum.sections.sect_id;


CREATE SEQUENCE public.seq_prog_entity_id
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.seq_prog_entity_id OWNER TO postgres;

--
-- Name: seq_secd_id; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.seq_secd_id
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.seq_secd_id OWNER TO postgres;

--
-- Name: seq_sect_id; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.seq_sect_id
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.seq_sect_id OWNER TO postgres;

--
-- Name: seq_sedm_id; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.seq_sedm_id
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.seq_sedm_id OWNER TO postgres;

-- Name: program_entity prog_entity_id; Type: DEFAULT; Schema: curriculum; Owner: postgres
--

ALTER TABLE ONLY curriculum.program_entity ALTER COLUMN prog_entity_id SET DEFAULT nextval('curriculum.program_entity_prog_entity_id_seq'::regclass);

--
-- Name: section_detail secd_id; Type: DEFAULT; Schema: curriculum; Owner: postgres
--

ALTER TABLE ONLY curriculum.section_detail ALTER COLUMN secd_id SET DEFAULT nextval('curriculum.section_detail_secd_id_seq'::regclass);

--
-- Name: section_detail_material sedm_id; Type: DEFAULT; Schema: curriculum; Owner: postgres
--

ALTER TABLE ONLY curriculum.section_detail_material ALTER COLUMN sedm_id SET DEFAULT nextval('curriculum.section_detail_material_sedm_id_seq'::regclass);

--
-- Name: sections sect_id; Type: DEFAULT; Schema: curriculum; Owner: postgres
--

ALTER TABLE ONLY curriculum.sections ALTER COLUMN sect_id SET DEFAULT nextval('curriculum.sections_sect_id_seq'::regclass);
