create table msig.assets
(
	ast_id serial not null
		constraint asset_pk
			primary key,
    ast_name varchar not null,
    ast_contract_type varchar not null,
    ast_address varchar(36) not null,
    ast_dexter_address varchar(36),
    ast_scale int not null,
    ast_ticker varchar not null,
	ctr_id int
		constraint assets_contracts_ctr_id_fk
			references msig.contracts
);

create unique index assets_ctr_id_ast_address_uindex
	on assets (ctr_id, ast_address);
