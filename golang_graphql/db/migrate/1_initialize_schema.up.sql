begin;
create extension if not exists "uuid-ossp";
CREATE TABLE locations(
    longitude int4 NULL,
    latitude int4 NULL,
    city varchar NOT NULL,
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    CONSTRAINT location_pk PRIMARY KEY (id),
    CONSTRAINT locations_ukey UNIQUE(city)
    
);

CREATE TABLE companies (

    "name" varchar NOT NULL,

    id uuid NOT NULL DEFAULT uuid_generate_v4(),

    location_id uuid  NULL,

    CONSTRAINT companies_pk PRIMARY KEY (id),
    CONSTRAINT company_location_fk FOREIGN KEY (location_id) REFERENCES locations(id),
    CONSTRAINT companies_ukey UNIQUE(name)
);




CREATE TABLE employees (

    "name" varchar NOT NULL,

    age int4 NOT NULL,

    gender varchar NOT NULL,

    company_id uuid NOT NULL,

    id uuid NOT NULL DEFAULT uuid_generate_v4(),

    location_id uuid NOT NULL,

    CONSTRAINT employees_fk FOREIGN KEY (location_id) REFERENCES locations(id),

    CONSTRAINT employees_fk_1 FOREIGN KEY (company_id) REFERENCES companies(id),
        CONSTRAINT employees_ukey UNIQUE(name)
);
commit;