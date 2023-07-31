# SQLite Design Notes


## Required Storage

<data> : <SQLite type>

### Access Token Data

- access_token : TEXT
  - comment(s): user access token
- time aquired : INTEGER
  - comment(s): SQLite does not have datatype for date & time; however, there are built-in date and time functions for Unix Time conversion
- expires_in : INTEGER
  - comment(s): maybe not neccesary since all tokens expire in 3600 seconds, could hard code in?
- refresh_token : TEXT
  - comment(s): self-explanatory, requires update
- scope : TEXT
  - comment(s): scope is a string of space seperated values. should I atomize?
