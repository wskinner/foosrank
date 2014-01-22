-- to run enter: sqlite3 games.db < games.ddl --

CREATE TABLE IF NOT EXISTS Players (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    FirstName VARCHAR(50),
    LastName VARCHAR(50),
    PlayerId VARCHAR(100));

CREATE TABLE IF NOT EXISTS Games (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    winnerId INTEGER,
    loserId INTEGER,
    WinnerScore INTEGER,
    LoserScore INTEGER,
    GameId INTEGER,
    FOREIGN KEY(WinnerId) REFERENCES Players(id)
    FOREIGN KEY(LoserId) REFERENCES Players(id));
