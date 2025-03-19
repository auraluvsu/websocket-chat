create database tcpDB;
use tcpDB;
create table tblUser (
    name    varchar(255)
    userID  varchar(255)
    isAdmin boolean
    dateCreated datetime
);

create table tblRoom (
    roomID  varchar(255)
    roomname    varchar
    port    int
);
--Not sure why my formatter isn't working...
--Oh well, life goes on i guess.
