CREATE TYPE addressType AS ( 
	city  	        text,
	state  	        text,
	country  	    text,
	line_1  	    text,
	line_1_check    text,
	line_2  	    text,
	zip  			text,
	zip_check  	    text
); 

CREATE TABLE IF NOT EXISTS cards (
    id                      text COLLATE pg_catalog."default" NOT NULL,
    object                  text COLLATE pg_catalog."default" NOT NULL,
    brand                   text COLLATE pg_catalog."default" NOT NULL,
    address                 addressType,  
    country                 text COLLATE pg_catalog."default" NOT NULL,
    customer                text COLLATE pg_catalog."default",
    cvc_check               text COLLATE pg_catalog."default" NOT NULL,
    dynamic_last4           text COLLATE pg_catalog."default",
    exp_month               bigint NOT NULL,
    exp_year                bigint NOT NULL,
    fingerprint             text COLLATE pg_catalog."default" NOT NULL,
    funding                 text COLLATE pg_catalog."default" NOT NULL,
    last4                   text COLLATE pg_catalog."default" NOT NULL,
    metadata                text COLLATE pg_catalog."default" NOT NULL,
    name                    text COLLATE pg_catalog."default",
    tokenization_method     text COLLATE pg_catalog."default" NOT NULL,
    created_at              timestamp without time zone NOT NULL,
    updated_at              timestamp without time zone NOT NULL,
    deleted_at              timestamp without time zone,
);