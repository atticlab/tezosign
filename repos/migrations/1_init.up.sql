create table msig.contracts
(
	ctr_id serial not null
		constraint contracts_pk
			primary key,
	ctr_address varchar(36) not null,
	ctr_last_block_level int
);

create unique index contracts_address_uindex
	on msig.contracts (ctr_address);

create table msig.requests
(
	req_id serial not null
		constraint requests_pk
			primary key,
	ctr_id int not null
		constraint requests_contracts_id_fk
			references msig.contracts,
    req_hash varchar(32) not null,
	req_status varchar default 'wait' not null,
	req_counter int,
	req_created_at timestamp without time zone default now() not null,
    req_info text not null,
    req_network_id varchar not null,
	req_type text not null,
 	req_operation_id varchar(51)
);

create unique index requests_req_operation_id_uindex
	on msig.requests (req_operation_id);

create table msig.signatures
(
	sig_id serial not null
		constraint signatures_pk
			primary key,
	req_id int not null
		constraint signatures_requests_id_fk
			references msig.requests,
	sig_index int not null,
	sig_data varchar not null,
	sig_type varchar not null
);

create unique index signatures_sign_uindex
	on msig.signatures (sig_data);


CREATE VIEW request_json_signatures AS
select req_id,
       json_agg(json_build_object('index', sig_index, 'type', sig_type, 'signature', sig_data)) as signatures
from (select * from signatures order by sig_index, sig_type) AS s
GROUP BY req_id;

CREATE OR REPLACE VIEW request_json_signatures_typed AS
select req_id, sig_type,
       json_agg(json_build_object('index', sig_index, 'type', sig_type, 'signature', sig_data)) as signatures
from (select * from signatures order by sig_index, sig_type) AS s
GROUP BY req_id, sig_type;