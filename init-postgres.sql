create table meeting (
    id char(20) not null,
    host varchar(50) not null,
    guest varchar(50) not null,
    date varchar(20) not null,
    duration int not null,
	primary key (id)
)
;
insert into meeting (id, host, guest, date, duration) values ('213f45de1', 'deli', 'bayram', '20210423', 30)
;

