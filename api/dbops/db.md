# 数据库表设计

## users
```
TABLE: users
id UNSIGNED INT, PRIMARY KEY, AUTO_INCREAMENT
login_name VARCHAR(64), UNIQUE KEY
pwd TEXT
```

## video
```
TABLE: video_info
id VARCHAR(64), PRIMARY KEY, NOT NULL
author_id UNSIGNED INT
name TEXT
display_ctime TEXT
create_time DATETIME
```

## comments
```
TABLE: comments
id VARCHAR(64), PRIMARY KEY, NOT NULL
video_id VARCHAR(64)
author_id UNSIGNED INT
content TEXT
time DATETIME
```

## sessions
```
TABLE:sessions
session_id TINYTEXT, PRIMARY KEY, NOT NULL
TLL TINYTEXT
login_name VARCHAR(64)
```