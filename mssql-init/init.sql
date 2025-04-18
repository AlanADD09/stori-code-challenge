-- mssql-init/init.sql
IF NOT EXISTS (SELECT * FROM sys.databases WHERE name = 'Stori')
BEGIN
    CREATE DATABASE Stori;
END
GO

USE Stori;
GO

IF NOT EXISTS (SELECT * FROM sys.tables WHERE name = 'transactions')
BEGIN
    CREATE TABLE transactions (
        id INT IDENTITY(1,1) PRIMARY KEY,
        date DATE NOT NULL,
        amount DECIMAL(18,2) NOT NULL
    );
END
GO