--

-- Name: payment; Type: SCHEMA; Schema: -; Owner: postgres

--

CREATE SCHEMA payment;

ALTER SCHEMA payment OWNER TO postgres;

--

-- Name: bank; Type: TABLE; Schema: payment; Owner: postgres

--

CREATE TABLE
    payment.bank (
        bank_entity_id integer NOT NULL,
        bank_code character varying(10),
        bank_name character varying(55),
        bank_modified_date timestamp without time zone
    );

ALTER TABLE payment.bank OWNER TO postgres;

--

-- Name: fintech; Type: TABLE; Schema: payment; Owner: postgres

--

CREATE TABLE
    payment.fintech (
        fint_entity_id integer NOT NULL,
        fint_code character varying(10),
        fint_name character varying(55),
        fint_modified_date timestamp without time zone
    );

ALTER TABLE payment.fintech OWNER TO postgres;

--

-- Name: transaction_payment; Type: TABLE; Schema: payment; Owner: postgres

--

CREATE TABLE
    payment.transaction_payment (
        trpa_id integer NOT NULL,
        trpa_code_number character varying(55),
        trpa_order_number character varying(25),
        trpa_debit numeric,
        trpa_credit numeric,
        trpa_type character varying(15),
        trpa_note character varying(255),
        trpa_modified_date timestamp without time zone,
        trpa_source_id character varying(25) NOT NULL,
        trpa_target_id character varying(25) NOT NULL,
        trpa_user_entity_id integer,
        CONSTRAINT transaction_payment_trpa_type_check CHECK ( ( (trpa_type):: text = ANY ( (
                        ARRAY ['top-up':: character varying,
                        'transfer':: character varying,
                        'order':: character varying,
                        'refund':: character varying]
                    ):: text []
                )
            )
        )
    );

ALTER TABLE payment.transaction_payment OWNER TO postgres;

--

-- Name: transaction_payment_trpa_id_seq; Type: SEQUENCE; Schema: payment; Owner: postgres

--

CREATE SEQUENCE payment.transaction_payment_trpa_id_seq AS integer START
WITH
    1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER TABLE
    payment.transaction_payment_trpa_id_seq OWNER TO postgres;

--

-- Name: transaction_payment_trpa_id_seq; Type: SEQUENCE OWNED BY; Schema: payment; Owner: postgres

--

ALTER SEQUENCE payment.transaction_payment_trpa_id_seq OWNED BY payment.transaction_payment.trpa_id;

--

-- Name: users_account; Type: TABLE; Schema: payment; Owner: postgres

--

CREATE TABLE
    payment.users_account (
        usac_bank_entity_id integer NOT NULL,
        usac_user_entity_id integer NOT NULL,
        usac_account_number character varying(25),
        usac_saldo numeric,
        usac_type character varying(15),
        usac_start_date timestamp without time zone,
        usac_end_date timestamp without time zone,
        usac_modified_date timestamp without time zone,
        usac_status character varying(15),
        CONSTRAINT users_account_usac_status_check CHECK ( ( (usac_status):: text = ANY ( (
                        ARRAY ['active':: character varying,
                        'inactive':: character varying,
                        'blocked':: character varying]
                    ):: text []
                )
            )
        ),
        CONSTRAINT users_account_usac_type_check CHECK ( ( (usac_type):: text = ANY ( (
                        ARRAY ['debit':: character varying,
                        'credit-card':: character varying,
                        'payment':: character varying]
                    ):: text []
                )
            )
        )
    );

ALTER TABLE payment.users_account OWNER TO postgres;